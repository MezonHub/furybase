package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/furybase/furybase/testutil/sample"
	"github.com/furybase/furybase/x/fvalidator/types"
	"github.com/stretchr/testify/require"
)

func TestMsgAddFValidator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgInitFValidator
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgInitFValidator{
				Creator:        "invalid_address",
				Denom:          sample.AccAddress(),
				ValAddressList: []string{"cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0"},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgInitFValidator{
				Creator:        sample.AccAddress(),
				Denom:          sample.AccAddress(),
				ValAddressList: []string{"cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
