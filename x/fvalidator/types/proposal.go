package types

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	fvotetypes "github.com/furybase/furybase/x/fvote/types"
	"github.com/tendermint/tendermint/crypto"
)

const TypeUpdateFValidatorProposal = "update_r_validator"
const TypeUpdateFValidatorReportProposal = "update_r_validator_report"

var _ sdk.Msg = &UpdateFValidatorProposal{}
var _ sdk.Msg = &UpdateFValidatorReportProposal{}
var _ fvotetypes.Content = &UpdateFValidatorProposal{}
var _ fvotetypes.Content = &UpdateFValidatorReportProposal{}

func init() {
	fvotetypes.RegisterProposalType(TypeUpdateFValidatorProposal)
	fvotetypes.RegisterProposalTypeCodec(&UpdateFValidatorProposal{}, "fvalidator/UpdateFValidator")
	fvotetypes.RegisterProposalType(TypeUpdateFValidatorReportProposal)
	fvotetypes.RegisterProposalTypeCodec(&UpdateFValidatorReportProposal{}, "fvalidator/UpdateFValidatorReport")
}

func NewUpdateFValidatorProposal(creator string, denom string, poolAddress, oldAddress string, newAddress string, cycle *Cycle) *UpdateFValidatorProposal {
	msg := UpdateFValidatorProposal{
		Denom:       denom,
		PoolAddress: poolAddress,
		OldAddress:  oldAddress,
		NewAddress:  newAddress,
		Cycle:       cycle,
	}
	msg.setPropId()

	msg.Creator = creator

	return &msg
}

func (p *UpdateFValidatorProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
}

func (p *UpdateFValidatorProposal) ProposalRoute() string {
	return ModuleName
}

func (p *UpdateFValidatorProposal) ProposalType() string {
	return TypeUpdateFValidatorProposal
}

func (p *UpdateFValidatorProposal) InFavour() bool {
	return true
}

func (msg *UpdateFValidatorProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *UpdateFValidatorProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *UpdateFValidatorProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !(msg.Denom == msg.Cycle.Denom && msg.PoolAddress == msg.Cycle.PoolAddress) {
		return fmt.Errorf("denom or pool address not equal")
	}
	return nil
}

func NewUpdateFValidatorReportProposal(creator string, denom string, poolAddress string, cycle *Cycle, status UpdateFValidatorStatus) *UpdateFValidatorReportProposal {
	msg := UpdateFValidatorReportProposal{
		Denom:       denom,
		PoolAddress: poolAddress,
		Cycle:       cycle,
		Status:      status,
	}
	msg.setPropId()

	msg.Creator = creator

	return &msg
}

func (p *UpdateFValidatorReportProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
}

func (p *UpdateFValidatorReportProposal) ProposalRoute() string {
	return ModuleName
}

func (p *UpdateFValidatorReportProposal) ProposalType() string {
	return TypeUpdateFValidatorReportProposal
}

func (p *UpdateFValidatorReportProposal) InFavour() bool {
	return true
}

func (msg *UpdateFValidatorReportProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *UpdateFValidatorReportProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *UpdateFValidatorReportProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !(msg.Denom == msg.Cycle.Denom && msg.PoolAddress == msg.Cycle.PoolAddress) {
		return fmt.Errorf("denom or pool address not equal")
	}
	return nil
}
