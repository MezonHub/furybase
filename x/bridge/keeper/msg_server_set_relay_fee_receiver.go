package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyunderverse/furybase/x/bridge/types"
	sudoTypes "github.com/furyunderverse/furybase/x/sudo/types"
)

func (k msgServer) SetRelayFeeReceiver(goCtx context.Context, msg *types.MsgSetRelayFeeReceiver) (*types.MsgSetRelayFeeReceiverResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	accAddress, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	k.Keeper.SetRelayFeeReceiver(ctx, accAddress)

	return &types.MsgSetRelayFeeReceiverResponse{}, nil
}
