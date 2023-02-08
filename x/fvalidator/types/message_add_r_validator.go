package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddFValidator = "add_r_validator"

var _ sdk.Msg = &MsgAddFValidator{}

func NewMsgAddFValidator(creator string, denom string, poolAddress string, valAddress string) *MsgAddFValidator {
	return &MsgAddFValidator{
		Creator:     creator,
		Denom:       denom,
		PoolAddress: poolAddress,
		ValAddress:  valAddress,
	}
}

func (msg *MsgAddFValidator) Route() string {
	return RouterKey
}

func (msg *MsgAddFValidator) Type() string {
	return TypeMsgAddFValidator
}

func (msg *MsgAddFValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddFValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddFValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
