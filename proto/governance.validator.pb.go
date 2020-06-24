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
	if this.MarginConfiguration != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarginConfiguration); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarginConfiguration", err)
		}
	}
	return nil
}
func (this *UpdateMarket) Validate() error {
	return nil
}
func (this *FutureProduct) Validate() error {
	if this.Maturity == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Maturity", fmt.Errorf(`value '%v' must not be an empty string`, this.Maturity))
	}
	if this.Asset == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Asset", fmt.Errorf(`value '%v' must not be an empty string`, this.Asset))
	}
	return nil
}
func (this *InstrumentConfiguration) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Code == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Code", fmt.Errorf(`value '%v' must not be an empty string`, this.Code))
	}
	if this.BaseName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("BaseName", fmt.Errorf(`value '%v' must not be an empty string`, this.BaseName))
	}
	if this.QuoteName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("QuoteName", fmt.Errorf(`value '%v' must not be an empty string`, this.QuoteName))
	}
	if oneOfNester, ok := this.GetProduct().(*InstrumentConfiguration_Future); ok {
		if oneOfNester.Future != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Future); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Future", err)
			}
		}
	}
	return nil
}
func (this *NewMarketConfiguration) Validate() error {
	if nil == this.Instrument {
		return github_com_mwitkow_go_proto_validators.FieldError("Instrument", fmt.Errorf("message must exist"))
	}
	if this.Instrument != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Instrument); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Instrument", err)
		}
	}
	if oneOfNester, ok := this.GetRiskParameters().(*NewMarketConfiguration_Simple); ok {
		if oneOfNester.Simple != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Simple); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Simple", err)
			}
		}
	}
	if oneOfNester, ok := this.GetRiskParameters().(*NewMarketConfiguration_LogNormal); ok {
		if oneOfNester.LogNormal != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.LogNormal); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LogNormal", err)
			}
		}
	}
	if !(this.DecimalPlaces < 150) {
		return github_com_mwitkow_go_proto_validators.FieldError("DecimalPlaces", fmt.Errorf(`value '%v' must be less than '150'`, this.DecimalPlaces))
	}
	if oneOfNester, ok := this.GetTradingMode().(*NewMarketConfiguration_Continuous); ok {
		if oneOfNester.Continuous != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Continuous); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Continuous", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTradingMode().(*NewMarketConfiguration_Discrete); ok {
		if oneOfNester.Discrete != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Discrete); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Discrete", err)
			}
		}
	}
	return nil
}
func (this *NewMarket) Validate() error {
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
func (this *NewAsset) Validate() error {
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
	if oneOfNester, ok := this.GetChange().(*ProposalTerms_NewAsset); ok {
		if oneOfNester.NewAsset != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.NewAsset); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NewAsset", err)
			}
		}
	}
	return nil
}
func (this *GovernanceData) Validate() error {
	if this.Proposal != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proposal); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proposal", err)
		}
	}
	for _, item := range this.Yes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Yes", err)
			}
		}
	}
	for _, item := range this.No {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("No", err)
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
