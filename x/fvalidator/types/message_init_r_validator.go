package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitFValidator = "init_r_validator"

var _ sdk.Msg = &MsgInitFValidator{}

func NewMsgInitFValidator(creator string, denom, poolAddress string, addressList []string) *MsgInitFValidator {
	return &MsgInitFValidator{
		Creator:        creator,
		Denom:          denom,
		PoolAddress:    poolAddress,
		ValAddressList: addressList,
	}
}

func (msg *MsgInitFValidator) Route() string {
	return RouterKey
}

func (msg *MsgInitFValidator) Type() string {
	return TypeMsgInitFValidator
}

func (msg *MsgInitFValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitFValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitFValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return err
	}

	if len(msg.ValAddressList) == 0 {
		return fmt.Errorf("address list is empty")
	}
	return nil
}
