package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/furybase/furybase/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAddChainId_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddChainId
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddChainId{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAddChainId{
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
