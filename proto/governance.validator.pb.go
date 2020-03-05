// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/governance.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *NetworkConfiguration) Validate() error {
	return nil
}
func (this *UpdateMarket) Validate() error {
	return nil
}
func (this *NewMarket) Validate() error {
	return nil
}
func (this *UpdateNetwork) Validate() error {
	if nil == this.Changes {
		return github_com_mwitkow_go_proto_validators.FieldError("Changes", fmt.Errorf("message must exist"))
	}
	if this.Changes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Changes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Changes", err)
		}
	}
	return nil
}
func (this *ProposalTerms) Validate() error {
	if !(this.ClosingTimestamp > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ClosingTimestamp", fmt.Errorf(`value '%v' must be greater than '0'`, this.ClosingTimestamp))
	}
	if !(this.EnactmentTimestamp > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("EnactmentTimestamp", fmt.Errorf(`value '%v' must be greater than '0'`, this.EnactmentTimestamp))
	}
	if !(this.MinParticipationStake > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MinParticipationStake", fmt.Errorf(`value '%v' must be greater than '0'`, this.MinParticipationStake))
	}
	if oneOfNester, ok := this.GetChange().(*ProposalTerms_UpdateMarket); ok {
		if oneOfNester.UpdateMarket != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.UpdateMarket); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UpdateMarket", err)
			}
		}
	}
	if oneOfNester, ok := this.GetChange().(*ProposalTerms_NewMarket); ok {
		if oneOfNester.NewMarket != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.NewMarket); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NewMarket", err)
			}
		}
	}
	if oneOfNester, ok := this.GetChange().(*ProposalTerms_UpdateNetwork); ok {
		if oneOfNester.UpdateNetwork != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.UpdateNetwork); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UpdateNetwork", err)
			}
		}
	}
	return nil
}
func (this *Proposal) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	if _, ok := Proposal_State_name[int32(this.State)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("State", fmt.Errorf(`value '%v' must be a valid Proposal_State field`, this.State))
	}
	if nil == this.Terms {
		return github_com_mwitkow_go_proto_validators.FieldError("Terms", fmt.Errorf("message must exist"))
	}
	if this.Terms != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Terms); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Terms", err)
		}
	}
	return nil
}
func (this *Vote) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	if _, ok := Vote_Value_name[int32(this.Value)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("Value", fmt.Errorf(`value '%v' must be a valid Vote_Value field`, this.Value))
	}
	if this.ProposalID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProposalID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProposalID))
	}
	return nil
}
