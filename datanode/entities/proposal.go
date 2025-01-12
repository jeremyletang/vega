// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.DATANODE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package entities

import (
	"encoding/json"
	"fmt"
	"time"

	"code.vegaprotocol.io/vega/libs/num"
	v2 "code.vegaprotocol.io/vega/protos/data-node/api/v2"
	"code.vegaprotocol.io/vega/protos/vega"
	"google.golang.org/protobuf/encoding/protojson"
)

type ProposalType string

var (
	ProposalTypeNewMarket              = ProposalType("newMarket")
	ProposalTypeNewAsset               = ProposalType("newAsset")
	ProposalTypeUpdateAsset            = ProposalType("updateAsset")
	ProposalTypeUpdateMarket           = ProposalType("updateMarket")
	ProposalTypeUpdateNetworkParameter = ProposalType("updateNetworkParameter")
	ProposalTypeNewFreeform            = ProposalType("newFreeform")
)

type _Proposal struct{}

type ProposalID = ID[_Proposal]

type Proposal struct {
	ID                      ProposalID
	Reference               string
	PartyID                 PartyID
	State                   ProposalState
	Rationale               ProposalRationale
	Terms                   ProposalTerms
	Reason                  ProposalError
	ErrorDetails            string
	ProposalTime            time.Time
	VegaTime                time.Time
	RequiredMajority        num.Decimal
	RequiredParticipation   num.Decimal
	RequiredLPMajority      *num.Decimal
	RequiredLPParticipation *num.Decimal
	TxHash                  TxHash
}

func (p *Proposal) ToProto() *vega.Proposal {
	var lpMajority *string
	if !p.RequiredLPMajority.IsZero() {
		lpMajority = toPtr(p.RequiredLPMajority.String())
	}
	var lpParticipation *string
	if !p.RequiredLPParticipation.IsZero() {
		lpParticipation = toPtr(p.RequiredLPParticipation.String())
	}

	pp := vega.Proposal{
		Id:                                     p.ID.String(),
		Reference:                              p.Reference,
		PartyId:                                p.PartyID.String(),
		State:                                  vega.Proposal_State(p.State),
		Rationale:                              p.Rationale.ProposalRationale,
		Timestamp:                              p.ProposalTime.UnixNano(),
		Terms:                                  p.Terms.ProposalTerms,
		Reason:                                 vega.ProposalError(p.Reason),
		ErrorDetails:                           p.ErrorDetails,
		RequiredMajority:                       p.RequiredMajority.String(),
		RequiredParticipation:                  p.RequiredParticipation.String(),
		RequiredLiquidityProviderMajority:      lpMajority,
		RequiredLiquidityProviderParticipation: lpParticipation,
	}
	return &pp
}

func (p Proposal) Cursor() *Cursor {
	pc := ProposalCursor{
		State:    p.State,
		VegaTime: p.VegaTime,
		ID:       p.ID,
	}
	return NewCursor(pc.String())
}

func (p Proposal) ToProtoEdge(_ ...any) (*v2.GovernanceDataEdge, error) {
	return &v2.GovernanceDataEdge{
		Node: &vega.GovernanceData{
			Proposal: p.ToProto(),
		},
		Cursor: p.Cursor().Encode(),
	}, nil
}

func ProposalFromProto(pp *vega.Proposal, txHash TxHash) (Proposal, error) {
	var err error
	var majority num.Decimal
	if len(pp.RequiredMajority) <= 0 {
		majority = num.DecimalZero()
	} else if majority, err = num.DecimalFromString(pp.RequiredMajority); err != nil {
		return Proposal{}, err
	}

	var participation num.Decimal
	if len(pp.RequiredParticipation) <= 0 {
		participation = num.DecimalZero()
	} else if participation, err = num.DecimalFromString(pp.RequiredParticipation); err != nil {
		return Proposal{}, err
	}

	lpMajority := num.DecimalZero()
	if pp.RequiredLiquidityProviderMajority != nil && len(*pp.RequiredLiquidityProviderMajority) > 0 {
		if lpMajority, err = num.DecimalFromString(*pp.RequiredLiquidityProviderMajority); err != nil {
			return Proposal{}, err
		}
	}
	lpParticipation := num.DecimalZero()
	if pp.RequiredLiquidityProviderParticipation != nil && len(*pp.RequiredLiquidityProviderParticipation) > 0 {
		if lpParticipation, err = num.DecimalFromString(*pp.RequiredLiquidityProviderParticipation); err != nil {
			return Proposal{}, err
		}
	}

	p := Proposal{
		ID:                      ProposalID(pp.Id),
		Reference:               pp.Reference,
		PartyID:                 PartyID(pp.PartyId),
		State:                   ProposalState(pp.State),
		Rationale:               ProposalRationale{pp.Rationale},
		Terms:                   ProposalTerms{pp.Terms},
		Reason:                  ProposalError(pp.Reason),
		ErrorDetails:            pp.ErrorDetails,
		ProposalTime:            time.Unix(0, pp.Timestamp),
		RequiredMajority:        majority,
		RequiredParticipation:   participation,
		RequiredLPMajority:      &lpMajority,
		RequiredLPParticipation: &lpParticipation,
		TxHash:                  txHash,
	}
	return p, nil
}

type ProposalRationale struct {
	*vega.ProposalRationale
}

func (pt ProposalRationale) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(pt)
}

func (pt *ProposalRationale) UnmarshalJSON(b []byte) error {
	pt.ProposalRationale = &vega.ProposalRationale{}
	return protojson.Unmarshal(b, pt)
}

type ProposalTerms struct {
	*vega.ProposalTerms
}

func (pt ProposalTerms) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(pt)
}

func (pt *ProposalTerms) UnmarshalJSON(b []byte) error {
	pt.ProposalTerms = &vega.ProposalTerms{}
	return protojson.Unmarshal(b, pt)
}

type ProposalCursor struct {
	State    ProposalState `json:"state"`
	VegaTime time.Time     `json:"vega_time"`
	ID       ProposalID    `json:"id"`
}

func (pc ProposalCursor) String() string {
	bs, err := json.Marshal(pc)
	if err != nil {
		panic(fmt.Errorf("failed to marshal proposal cursor: %w", err))
	}
	return string(bs)
}

func (pc *ProposalCursor) Parse(cursorString string) error {
	if cursorString == "" {
		return nil
	}
	return json.Unmarshal([]byte(cursorString), pc)
}

func toPtr[T any](t T) *T {
	return &t
}
