package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyunderverse/furybase/x/fvalidator/types"
	sudoTypes "github.com/furyunderverse/furybase/x/sudo/types"
)

func (k msgServer) RmFValidator(goCtx context.Context, msg *types.MsgRmFValidator) (*types.MsgRmFValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	newVal := types.FValidator{
		Denom:       msg.Denom,
		PoolAddress: msg.PoolAddress,
		ValAddress:  msg.NewAddress,
	}
	if !k.Keeper.HasSelectedFValidator(ctx, &newVal) {
		return nil, types.ErrFValidatorNotExist
	}

	latestVotedCycle := k.GetLatestVotedCycle(ctx, msg.Denom, msg.PoolAddress)
	willUseCycle := types.Cycle{
		Denom:       msg.Denom,
		PoolAddress: msg.PoolAddress,
		Version:     latestVotedCycle.Version,
		Number:      latestVotedCycle.Number + 1,
	}

	proposal := types.NewUpdateFValidatorProposal(msg.Creator, msg.Denom, msg.PoolAddress, msg.OldAddress, msg.NewAddress, &willUseCycle)

	err := k.ProcessUpdateFValidatorProposal(ctx, proposal)
	if err != nil {
		return nil, err
	}
	return &types.MsgRmFValidatorResponse{}, nil
}
