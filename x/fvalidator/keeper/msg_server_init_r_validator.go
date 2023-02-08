package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyunderverse/furybase/x/fvalidator/types"
	sudoTypes "github.com/furyunderverse/furybase/x/sudo/types"
)

// init fvalidator and can only init once
func (k msgServer) InitFValidator(goCtx context.Context, msg *types.MsgInitFValidator) (*types.MsgInitFValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	fvalidatorList := k.Keeper.GetSelectedFValidatorListByDenomPoolAddress(ctx, msg.Denom, msg.PoolAddress)
	if len(fvalidatorList) > 0 {
		return nil, types.ErrFValidatorAlreadyInit
	}

	if err := k.FBank Keeper.CheckAccAddress(ctx, msg.Denom, msg.PoolAddress); err != nil {
		return nil, err
	}

	addresses := ""
	for _, address := range msg.ValAddressList {
		if err := k.FBank Keeper.CheckValAddress(ctx, msg.Denom, address); err != nil {
			return nil, err
		}
		fValidator := types.FValidator{
			Denom:       msg.Denom,
			PoolAddress: msg.PoolAddress,
			ValAddress:  address,
		}

		if k.Keeper.HasSelectedFValidator(ctx, &fValidator) {
			return nil, types.ErrFValidatorAlreadyExist
		}

		k.Keeper.AddSelectedFValidator(ctx, &fValidator)

		addresses = addresses + ":" + address
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeInitFValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, msg.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyAddresses, addresses[1:]),
		),
	)
	return &types.MsgInitFValidatorResponse{}, nil
}
