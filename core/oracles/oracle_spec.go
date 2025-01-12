// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package oracles

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/libs/num"
	oraclespb "code.vegaprotocol.io/vega/protos/vega/oracles/v1"
)

var (
	// ErrMissingPubKeys is returned when the oraclespb.OracleSpec is missing
	// its public keys.
	ErrMissingPubKeys = errors.New("public keys are required")
	// ErrAtLeastOneFilterIsRequired is returned when the oraclespb.OracleSpec
	// has no expected properties nor filters. At least one of these should be
	// defined.
	ErrAtLeastOneFilterIsRequired = errors.New("at least one filter is required")
	// ErrInvalidTimestamp is returned when the timestamp has a negative value
	// which may happen in case of unsigned integer overflow.
	ErrInvalidTimestamp = errors.New("invalid timestamp")
	// ErrMissingPropertyKey is returned when a property key is undefined.
	ErrMissingPropertyKey = errors.New("a property key is required")
	// ErrMissingPropertyName is returned when a property as no name.
	ErrMissingPropertyName = errors.New("a property name is required")
	// ErrInvalidPropertyKey is returned if validation finds a reserved Vega property key.
	ErrInvalidPropertyKey = errors.New("property key is reserved")
)

type OracleSpecID string

type OracleSpec struct {
	// id is a unique identifier for the OracleSpec
	id OracleSpecID
	// pubKeys list all the authorized public keys from where an OracleData can
	// come from.
	pubKeys map[string]struct{}
	// filters holds all the expected property keys with the conditions they
	// should match.
	filters map[string]*filter
	// OriginalSpec is the protobuf description of OracleSpec
	OriginalSpec *types.OracleSpec
}

type filter struct {
	propertyName string
	propertyType oraclespb.PropertyKey_Type
	conditions   []condition
}

type condition func(string) (bool, error)

// NewOracleSpec builds an OracleSpec from a types.OracleSpec in a form that
// suits the processing of the filters.
func NewOracleSpec(originalSpec types.OracleSpec) (*OracleSpec, error) {
	if len(originalSpec.PubKeys) == 0 {
		return nil, ErrMissingPubKeys
	}

	pubKeys := map[string]struct{}{}
	for _, pk := range originalSpec.PubKeys {
		pubKeys[pk] = struct{}{}
	}

	if len(originalSpec.Filters) == 0 {
		return nil, ErrAtLeastOneFilterIsRequired
	}

	typedFilters := map[string]*filter{}
	for _, f := range originalSpec.Filters {
		if f.Key == nil {
			return nil, ErrMissingPropertyKey
		}
		if len(f.Key.Name) == 0 {
			return nil, ErrMissingPropertyName
		}

		conditions, err := toConditions(f.Key.Type, f.Conditions)
		if err != nil {
			return nil, err
		}

		typedFilter, ok := typedFilters[f.Key.Name]

		if !ok {
			typedFilters[f.Key.Name] = &filter{
				propertyName: f.Key.Name,
				propertyType: f.Key.Type,
				conditions:   conditions,
			}
			continue
		}

		if typedFilter.propertyType != f.Key.Type {
			return nil, errMismatchPropertyType(typedFilter.propertyName, typedFilter.propertyType, f.Key.Type)
		}

		typedFilter.conditions = append(typedFilter.conditions, conditions...)
	}

	return &OracleSpec{
		id:           OracleSpecID(originalSpec.ID),
		pubKeys:      pubKeys,
		filters:      typedFilters,
		OriginalSpec: &originalSpec,
	}, nil
}

func (s OracleSpec) EnsureBoundableProperty(property string, propType oraclespb.PropertyKey_Type) error {
	filter, ok := s.filters[property]
	if !ok {
		return fmt.Errorf("bound property \"%s\" not filtered by oracle spec", property)
	}

	if filter.propertyType != propType {
		return fmt.Errorf("bound type \"%v\" doesn't match filtered property type \"%s\"", propType, filter.propertyType)
	}

	return nil
}

func isInternalOracleData(data OracleData) bool {
	for k := range data.Data {
		if !strings.HasPrefix(k, BuiltinOraclePrefix) {
			return false
		}
	}

	return true
}

// MatchPubKeys tries to match the public keys from the provided OracleData object with the ones
// present in the Spec.
func (s *OracleSpec) MatchPubKeys(data OracleData) bool {
	return containsRequiredPubKeys(data.PubKeys, s.pubKeys)
}

// MatchData indicates if a given OracleData matches the spec or not.
func (s *OracleSpec) MatchData(data OracleData) (bool, error) {
	// if the data contains the internal oracle timestamp key, and only that key,
	// then we do not need to verify the public keys as there will not be one

	if !isInternalOracleData(data) && !containsRequiredPubKeys(data.PubKeys, s.pubKeys) {
		return false, nil
	}

	for propertyName, filter := range s.filters {
		dataValue, ok := data.Data[propertyName]
		if !ok {
			return false, nil
		}

		for _, condition := range filter.conditions {
			if matched, err := condition(dataValue); !matched || err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

// containsRequiredPubKeys verifies if all the public keys in the OracleData
// are within the list of currently authorized by the OracleSpec.
func containsRequiredPubKeys(dataPKs []string, authPks map[string]struct{}) bool {
	for _, pk := range dataPKs {
		if _, ok := authPks[pk]; !ok {
			return false
		}
	}
	return true
}

var conditionConverters = map[oraclespb.PropertyKey_Type]func(*types.OracleSpecCondition) (condition, error){
	oraclespb.PropertyKey_TYPE_INTEGER:   toIntegerCondition,
	oraclespb.PropertyKey_TYPE_DECIMAL:   toDecimalCondition,
	oraclespb.PropertyKey_TYPE_BOOLEAN:   toBooleanCondition,
	oraclespb.PropertyKey_TYPE_TIMESTAMP: toTimestampCondition,
	oraclespb.PropertyKey_TYPE_STRING:    toStringCondition,
}

func toConditions(typ oraclespb.PropertyKey_Type, cs []*types.OracleSpecCondition) ([]condition, error) {
	converter, ok := conditionConverters[typ]
	if !ok {
		return nil, errUnsupportedPropertyType(typ)
	}

	conditions := make([]condition, 0, len(cs))
	for _, c := range cs {
		cond, err := converter(c)
		if err != nil {
			return nil, err
		}

		conditions = append(conditions, cond)
	}
	return conditions, nil
}

func toIntegerCondition(c *types.OracleSpecCondition) (condition, error) {
	condValue, err := toInteger(c.Value)
	if err != nil {
		return nil, err
	}

	matcher, ok := integerMatchers[c.Operator]
	if !ok {
		return nil, err
	}

	return func(dataValue string) (bool, error) {
		parsedDataValue, err := toInteger(dataValue)
		if err != nil {
			return false, err
		}
		return matcher(parsedDataValue, condValue), nil
	}, nil
}

func toInteger(value string) (*num.Int, error) {
	convertedValue, hasError := num.IntFromString(value, 10)
	if hasError {
		return nil, fmt.Errorf("value \"%s\" is not a valid integer", value)
	}
	return convertedValue, nil
}

var integerMatchers = map[oraclespb.Condition_Operator]func(*num.Int, *num.Int) bool{
	oraclespb.Condition_OPERATOR_EQUALS:                equalsInteger,
	oraclespb.Condition_OPERATOR_GREATER_THAN:          greaterThanInteger,
	oraclespb.Condition_OPERATOR_GREATER_THAN_OR_EQUAL: greaterThanOrEqualInteger,
	oraclespb.Condition_OPERATOR_LESS_THAN:             lessThanInteger,
	oraclespb.Condition_OPERATOR_LESS_THAN_OR_EQUAL:    lessThanOrEqualInteger,
}

func equalsInteger(dataValue, condValue *num.Int) bool {
	return dataValue.EQ(condValue)
}

func greaterThanInteger(dataValue, condValue *num.Int) bool {
	return dataValue.GT(condValue)
}

func greaterThanOrEqualInteger(dataValue, condValue *num.Int) bool {
	return dataValue.GTE(condValue)
}

func lessThanInteger(dataValue, condValue *num.Int) bool {
	return dataValue.LT(condValue)
}

func lessThanOrEqualInteger(dataValue, condValue *num.Int) bool {
	return dataValue.LTE(condValue)
}

func toDecimalCondition(c *types.OracleSpecCondition) (condition, error) {
	condValue, err := toDecimal(c.Value)
	if err != nil {
		return nil, err
	}

	matcher, ok := decimalMatchers[c.Operator]
	if !ok {
		return nil, errUnsupportedOperatorForType(c.Operator, oraclespb.PropertyKey_TYPE_DECIMAL)
	}

	return func(dataValue string) (bool, error) {
		parsedDataValue, err := toDecimal(dataValue)
		if err != nil {
			return false, err
		}
		return matcher(parsedDataValue, condValue), nil
	}, nil
}

func toDecimal(value string) (num.Decimal, error) {
	return num.DecimalFromString(value)
}

var decimalMatchers = map[oraclespb.Condition_Operator]func(num.Decimal, num.Decimal) bool{
	oraclespb.Condition_OPERATOR_EQUALS:                equalsDecimal,
	oraclespb.Condition_OPERATOR_GREATER_THAN:          greaterThanDecimal,
	oraclespb.Condition_OPERATOR_GREATER_THAN_OR_EQUAL: greaterThanOrEqualDecimal,
	oraclespb.Condition_OPERATOR_LESS_THAN:             lessThanDecimal,
	oraclespb.Condition_OPERATOR_LESS_THAN_OR_EQUAL:    lessThanOrEqualDecimal,
}

func equalsDecimal(dataValue, condValue num.Decimal) bool {
	return dataValue.Equal(condValue)
}

func greaterThanDecimal(dataValue, condValue num.Decimal) bool {
	return dataValue.GreaterThan(condValue)
}

func greaterThanOrEqualDecimal(dataValue, condValue num.Decimal) bool {
	return dataValue.GreaterThanOrEqual(condValue)
}

func lessThanDecimal(dataValue, condValue num.Decimal) bool {
	return dataValue.LessThan(condValue)
}

func lessThanOrEqualDecimal(dataValue, condValue num.Decimal) bool {
	return dataValue.LessThanOrEqual(condValue)
}

func toTimestampCondition(c *types.OracleSpecCondition) (condition, error) {
	condValue, err := toTimestamp(c.Value)
	if err != nil {
		return nil, err
	}

	matcher, ok := timestampMatchers[c.Operator]
	if !ok {
		return nil, errUnsupportedOperatorForType(c.Operator, oraclespb.PropertyKey_TYPE_TIMESTAMP)
	}

	return func(dataValue string) (bool, error) {
		parsedDataValue, err := toTimestamp(dataValue)
		if err != nil {
			return false, err
		}
		return matcher(parsedDataValue, condValue), nil
	}, nil
}

func toTimestamp(value string) (int64, error) {
	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return parsedValue, err
	}

	if parsedValue < 0 {
		return parsedValue, ErrInvalidTimestamp
	}
	return parsedValue, nil
}

var timestampMatchers = map[oraclespb.Condition_Operator]func(int64, int64) bool{
	oraclespb.Condition_OPERATOR_EQUALS:                equalsTimestamp,
	oraclespb.Condition_OPERATOR_GREATER_THAN:          greaterThanTimestamp,
	oraclespb.Condition_OPERATOR_GREATER_THAN_OR_EQUAL: greaterThanOrEqualTimestamp,
	oraclespb.Condition_OPERATOR_LESS_THAN:             lessThanTimestamp,
	oraclespb.Condition_OPERATOR_LESS_THAN_OR_EQUAL:    lessThanOrEqualTimestamp,
}

func equalsTimestamp(dataValue, condValue int64) bool {
	return dataValue == condValue
}

func greaterThanTimestamp(dataValue, condValue int64) bool {
	return dataValue > condValue
}

func greaterThanOrEqualTimestamp(dataValue, condValue int64) bool {
	return dataValue >= condValue
}

func lessThanTimestamp(dataValue, condValue int64) bool {
	return dataValue < condValue
}

func lessThanOrEqualTimestamp(dataValue, condValue int64) bool {
	return dataValue <= condValue
}

func toBooleanCondition(c *types.OracleSpecCondition) (condition, error) {
	condValue, err := toBoolean(c.Value)
	if err != nil {
		return nil, err
	}

	matcher, ok := booleanMatchers[c.Operator]
	if !ok {
		return nil, errUnsupportedOperatorForType(c.Operator, oraclespb.PropertyKey_TYPE_BOOLEAN)
	}

	return func(dataValue string) (bool, error) {
		parsedDataValue, err := toBoolean(dataValue)
		if err != nil {
			return false, err
		}
		return matcher(parsedDataValue, condValue), nil
	}, nil
}

func toBoolean(value string) (bool, error) {
	return strconv.ParseBool(value)
}

var booleanMatchers = map[oraclespb.Condition_Operator]func(bool, bool) bool{
	oraclespb.Condition_OPERATOR_EQUALS: equalsBoolean,
}

func equalsBoolean(dataValue, condValue bool) bool {
	return dataValue == condValue
}

func toStringCondition(c *types.OracleSpecCondition) (condition, error) {
	matcher, ok := stringMatchers[c.Operator]
	if !ok {
		return nil, errUnsupportedOperatorForType(c.Operator, oraclespb.PropertyKey_TYPE_STRING)
	}

	return func(dataValue string) (bool, error) {
		return matcher(dataValue, c.Value), nil
	}, nil
}

var stringMatchers = map[oraclespb.Condition_Operator]func(string, string) bool{
	oraclespb.Condition_OPERATOR_EQUALS: equalsString,
}

func equalsString(dataValue, condValue string) bool {
	return dataValue == condValue
}

// errMismatchPropertyType is returned when a property is redeclared in
// conditions but with a different type.
func errMismatchPropertyType(prop string, first, newP oraclespb.PropertyKey_Type) error {
	return fmt.Errorf(
		"cannot redeclared property %s with different type, first %s then %s",
		prop, first, newP,
	)
}

// errUnsupportedOperatorForType is returned when the property type does not
// support the specified operator.
func errUnsupportedOperatorForType(o oraclespb.Condition_Operator, t oraclespb.PropertyKey_Type) error {
	return fmt.Errorf("unsupported operator %s for type %s", o, t)
}

// errUnsupportedPropertyType is returned when the filter specifies an
// unsupported property key type.
func errUnsupportedPropertyType(typ oraclespb.PropertyKey_Type) error {
	return fmt.Errorf("property type %s", typ)
}
