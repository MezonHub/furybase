package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furybase/furybase/x/mining/types"
	sudotypes "github.com/furybase/furybase/x/sudo/types"
)

func (k msgServer) SetMaxRewardPoolNumber(goCtx context.Context, msg *types.MsgSetMaxRewardPoolNumber) (*types.MsgSetMaxRewardPoolNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	k.Keeper.SetMaxRewardPoolNumber(ctx, msg.Number)

	return &types.MsgSetMaxRewardPoolNumberResponse{}, nil
}
