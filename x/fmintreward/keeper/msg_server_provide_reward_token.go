package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyunderverse/furybase/x/fmintreward/types"
)

func (k msgServer) ProvideRewardToken(goCtx context.Context, msg *types.MsgProvideRewardToken) (*types.MsgProvideRewardTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if err := k.bankKeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, sdk.NewCoins(msg.Amount)); err != nil {
		return nil, err
	}

	return &types.MsgProvideRewardTokenResponse{}, nil
}
