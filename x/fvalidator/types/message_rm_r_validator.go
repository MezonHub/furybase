package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmFValidator = "rm_r_validator"

var _ sdk.Msg = &MsgRmFValidator{}

func NewMsgRmFValidator(creator string, denom string, poolAddress string, oldAddress string, newAddress string) *MsgRmFValidator {
	return &MsgRmFValidator{
		Creator:     creator,
		Denom:       denom,
		PoolAddress: poolAddress,
		OldAddress:  oldAddress,
		NewAddress:  newAddress,
	}
}

func (msg *MsgRmFValidator) Route() string {
	return RouterKey
}

func (msg *MsgRmFValidator) Type() string {
	return TypeMsgRmFValidator
}

func (msg *MsgRmFValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmFValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmFValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
