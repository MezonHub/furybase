package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyunderverse/furybase/x/fvalidator/types"
	sudoTypes "github.com/furyunderverse/furybase/x/sudo/types"
)

func (k msgServer) AddFValidator(goCtx context.Context, msg *types.MsgAddFValidator) (*types.MsgAddFValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	if err := k.FBank Keeper.CheckValAddress(ctx, msg.Denom, msg.ValAddress); err != nil {
		return nil, err
	}

	if err := k.FBank Keeper.CheckAccAddress(ctx, msg.Denom, msg.PoolAddress); err != nil {
		return nil, err
	}

	fValidator := types.FValidator{
		Denom:       msg.Denom,
		PoolAddress: msg.PoolAddress,
		ValAddress:  msg.ValAddress,
	}

	if k.Keeper.HasSelectedFValidator(ctx, &fValidator) {
		return nil, types.ErrFValidatorAlreadyExist
	}

	snapShots := k.ledgerKeeper.CurrentEraSnapshots(ctx, msg.Denom)
	if len(snapShots.ShotIds) > 0 {
		return nil, types.ErrLedgerIsBusyWithEra
	}

	chainEra, found := k.ledgerKeeper.GetChainEra(ctx, msg.Denom)
	if !found {
		return nil, types.ErrLedgerChainEraNotExist
	}

	k.Keeper.AddSelectedFValidator(ctx, &fValidator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddFValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, msg.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyChainEra, fmt.Sprintf("%d", chainEra.Era)),
			sdk.NewAttribute(types.AttributeKeyAddedAddress, msg.ValAddress),
		),
	)
	return &types.MsgAddFValidatorResponse{}, nil
}
