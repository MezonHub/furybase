package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/furyunderverse/furybase/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAddValToWhitelist_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddValToWhitelist
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddValToWhitelist{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddValToWhitelist{
				Creator: sample.AccAddress(),
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
