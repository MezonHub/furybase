package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestMsgRmRewardPool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgRmRewardPool
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRmRewardPool{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgRmRewardPool{
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
