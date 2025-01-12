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

package statevar_test

import (
	"testing"

	"code.vegaprotocol.io/vega/core/types/statevar"
	"code.vegaprotocol.io/vega/libs/num"
	"code.vegaprotocol.io/vega/protos/vega"
	"github.com/stretchr/testify/require"
)

func TestDecimalVector(t *testing.T) {
	t.Run("test equality of two decimal vectors", testDecimalVectorEquality)
	t.Run("test two vectors are within tolerance of each other", testDecimalVectorWithinTol)
	t.Run("test converion of decimal vector to a decimal vector", testDecimalVectorToDecimal)
	t.Run("test conversion to proto", testVectorToProto)
}

// testDecimalVectorEquality tests that given the same key and equal/not equal value, equals function returns the correct value.
func testDecimalVectorEquality(t *testing.T) {
	kvb1 := &statevar.KeyValueBundle{}
	kvb1.KVT = append(kvb1.KVT, statevar.KeyValueTol{
		Key: "vector value",
		Val: &statevar.DecimalVector{Val: []num.Decimal{num.DecimalFromFloat(1.1), num.DecimalFromFloat(2.2), num.DecimalFromFloat(3.3), num.DecimalFromFloat(4.4)}},
	})

	kvb2 := &statevar.KeyValueBundle{}
	kvb2.KVT = append(kvb2.KVT, statevar.KeyValueTol{
		Key: "vector value",
		Val: &statevar.DecimalVector{Val: []num.Decimal{num.DecimalFromFloat(1.1), num.DecimalFromFloat(2.2), num.DecimalFromFloat(3.3)}},
	})

	// the vectors have different shape
	require.False(t, kvb1.Equals(kvb2))

	kvb3 := &statevar.KeyValueBundle{}
	kvb3.KVT = append(kvb3.KVT, statevar.KeyValueTol{
		Key: "vector value",
		Val: &statevar.DecimalVector{Val: []num.Decimal{num.DecimalFromFloat(2.2), num.DecimalFromFloat(3.3), num.DecimalFromFloat(4.4), num.DecimalFromFloat(1.1)}},
	})

	// the vectors have different values
	require.False(t, kvb1.Equals(kvb3))

	kvb4 := &statevar.KeyValueBundle{}
	kvb4.KVT = append(kvb4.KVT, statevar.KeyValueTol{
		Key: "vector value",
		Val: &statevar.DecimalVector{Val: []num.Decimal{num.DecimalFromFloat(1.1), num.DecimalFromFloat(2.2), num.DecimalFromFloat(3.3), num.DecimalFromFloat(4.4)}},
	})

	// they are equal
	require.True(t, kvb1.Equals(kvb4))
}

// testDecimalVectorWithinTol check that the values of the vectors are within tolerance.
func testDecimalVectorWithinTol(t *testing.T) {
	kvb1 := &statevar.KeyValueBundle{}
	kvb1.KVT = append(kvb1.KVT, statevar.KeyValueTol{
		Key:       "vector value",
		Val:       &statevar.DecimalVector{Val: []num.Decimal{num.DecimalFromFloat(1.1), num.DecimalFromFloat(2.2), num.DecimalFromFloat(3.3), num.DecimalFromFloat(4.4)}},
		Tolerance: tol,
	})

	kvb2 := &statevar.KeyValueBundle{}
	kvb2.KVT = append(kvb2.KVT, statevar.KeyValueTol{
		Key:       "vector value",
		Val:       &statevar.DecimalVector{Val: []num.Decimal{num.DecimalFromFloat(1.2), num.DecimalFromFloat(2.3), num.DecimalFromFloat(3.4), num.DecimalFromFloat(4.6)}},
		Tolerance: tol,
	})

	// almost within tolerance but not
	require.False(t, kvb1.WithinTolerance(kvb2))

	kvb3 := &statevar.KeyValueBundle{}
	kvb3.KVT = append(kvb3.KVT, statevar.KeyValueTol{
		Key:       "vector value",
		Val:       &statevar.DecimalVector{Val: []num.Decimal{num.DecimalFromFloat(1.099999999), num.DecimalFromFloat(2.100000005), num.DecimalFromFloat(3.3999999), num.DecimalFromFloat(4.3000001)}},
		Tolerance: tol,
	})

	// within tolerance on either side
	require.True(t, kvb1.WithinTolerance(kvb3))
}

// testDecimalVectorToDecimal tests conversion to decimal.
func testDecimalVectorToDecimal(t *testing.T) {
	kvb1 := &statevar.KeyValueBundle{}
	kvb1.KVT = append(kvb1.KVT, statevar.KeyValueTol{
		Key:       "vector value",
		Val:       &statevar.DecimalVector{Val: []num.Decimal{num.DecimalFromFloat(1.1), num.DecimalFromFloat(2.2), num.DecimalFromFloat(3.3), num.DecimalFromFloat(4.4)}},
		Tolerance: tol,
	})

	res := kvb1.KVT[0].Val
	switch v := res.(type) {
	case *statevar.DecimalVector:
		require.Equal(t, []num.Decimal{num.DecimalFromFloat(1.1), num.DecimalFromFloat(2.2), num.DecimalFromFloat(3.3), num.DecimalFromFloat(4.4)}, v.Val)
	default:
		t.Fail()
	}
}

// testVectorToProto tests conversion of vector variable to proto.
func testVectorToProto(t *testing.T) {
	kvb1 := &statevar.KeyValueBundle{}
	kvb1.KVT = append(kvb1.KVT, statevar.KeyValueTol{
		Key:       "vector value",
		Val:       &statevar.DecimalVector{Val: []num.Decimal{num.DecimalFromFloat(1.1), num.DecimalFromFloat(2.2), num.DecimalFromFloat(3.3), num.DecimalFromFloat(4.4)}},
		Tolerance: tol,
	})
	res := kvb1.ToProto()
	require.Equal(t, 1, len(res))
	require.Equal(t, "vector value", res[0].Key)
	require.Equal(t, "0.1", res[0].Tolerance)
	switch v := res[0].Value.Value.(type) {
	case *vega.StateVarValue_VectorVal:
		require.Equal(t, []string{num.DecimalFromFloat(1.1).String(), num.DecimalFromFloat(2.2).String(), num.DecimalFromFloat(3.3).String(), num.DecimalFromFloat(4.4).String()}, v.VectorVal.Value)
	default:
		t.Fail()
	}

	kvb2, err := statevar.KeyValueBundleFromProto(res)
	require.NoError(t, err)
	require.Equal(t, kvb1, kvb2)
}
