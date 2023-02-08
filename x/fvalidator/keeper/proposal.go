package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furybase/furybase/x/fvalidator/types"
)

// old val must exist && new val may exist or not exist in selectedFValidator
// pool will redelegate all delegation from old to new val
func (k Keeper) ProcessUpdateFValidatorProposal(ctx sdk.Context, p *types.UpdateFValidatorProposal) error {

	oldVal := types.FValidator{
		Denom:       p.Denom,
		PoolAddress: p.PoolAddress,
		ValAddress:  p.OldAddress,
	}
	newVal := types.FValidator{
		Denom:       p.Denom,
		PoolAddress: p.PoolAddress,
		ValAddress:  p.NewAddress,
	}
	if !k.HasSelectedFValidator(ctx, &oldVal) {
		return types.ErrFValidatorNotExist
	}
	if oldVal.ValAddress == newVal.ValAddress {
		return types.ErrOldEqualNewFValidator
	}

	if err := k.FBank Keeper.CheckValAddress(ctx, p.Denom, p.OldAddress); err != nil {
		return err
	}
	if err := k.FBank Keeper.CheckValAddress(ctx, p.Denom, p.NewAddress); err != nil {
		return err
	}
	if err := k.FBank Keeper.CheckAccAddress(ctx, p.Denom, p.PoolAddress); err != nil {
		return err
	}
	cycleSeconds := k.GetCycleSeconds(ctx, p.Denom)
	if cycleSeconds.Version != p.Cycle.Version {
		return types.ErrCycleVersionNotMatch
	}

	latestVotedCycle := k.GetLatestVotedCycle(ctx, p.Denom, p.PoolAddress)
	if !(p.Cycle.Version > latestVotedCycle.Version || (p.Cycle.Version == latestVotedCycle.Version && p.Cycle.Number > latestVotedCycle.Number)) {
		return types.ErrCycleBehindLatestCycle
	}
	latestDealedCycle := k.GetLatestDealedCycle(ctx, p.Denom, p.PoolAddress)
	if latestDealedCycle.Number != latestVotedCycle.Number || latestDealedCycle.Version != latestVotedCycle.Version {
		return types.ErrLatestVotedCycleNotDealed
	}
	snapShots := k.ledgerKeeper.CurrentEraSnapshots(ctx, p.Denom)
	if len(snapShots.ShotIds) > 0 {
		return types.ErrLedgerIsBusyWithEra
	}
	chainEra, found := k.ledgerKeeper.GetChainEra(ctx, p.Denom)
	if !found {
		return types.ErrLedgerChainEraNotExist
	}

	k.SetDealingFValidator(ctx, &types.DealingFValidator{
		Denom:         p.Denom,
		PoolAddress:   p.PoolAddress,
		OldValAddress: p.OldAddress,
		NewValAddress: p.NewAddress,
	})
	k.SetLatestVotedCycle(ctx, p.Cycle)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateFValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, p.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyChainEra, fmt.Sprintf("%d", chainEra.Era)),
			sdk.NewAttribute(types.AttributeKeyOldAddress, p.OldAddress),
			sdk.NewAttribute(types.AttributeKeyNewAddress, p.NewAddress),
			sdk.NewAttribute(types.AttributeKeyCycleVersion, fmt.Sprintf("%d", p.Cycle.Version)),
			sdk.NewAttribute(types.AttributeKeyCycleNumber, fmt.Sprintf("%d", p.Cycle.Number)),
			sdk.NewAttribute(types.AttributeKeyCycleSeconds, fmt.Sprintf("%d", cycleSeconds.Seconds)),
		),
	)

	return nil
}

func (k Keeper) ProcessUpdateFValidatorReportProposal(ctx sdk.Context, p *types.UpdateFValidatorReportProposal) error {
	latestVotedCycle := k.GetLatestVotedCycle(ctx, p.Denom, p.PoolAddress)
	if !(p.Cycle.Version == latestVotedCycle.Version && p.Cycle.Number == latestVotedCycle.Number) {
		return types.ErrReportCycleNotMatchLatestVotedCycle
	}
	dealingFValidator, found := k.GetDealingFValidator(ctx, p.Denom, p.PoolAddress)
	if !found {
		return types.ErrDealingFvalidatorNotFound
	}

	// should update fvalidator when redelegate success
	if p.Status == types.UpdateFValidatorStatusSuccess {
		k.RemoveSelectedFValidator(ctx, &types.FValidator{
			Denom:       dealingFValidator.Denom,
			PoolAddress: dealingFValidator.PoolAddress,
			ValAddress:  dealingFValidator.OldValAddress,
		})
		k.AddSelectedFValidator(ctx, &types.FValidator{
			Denom:       dealingFValidator.Denom,
			PoolAddress: dealingFValidator.PoolAddress,
			ValAddress:  dealingFValidator.NewValAddress,
		})
	}

	k.RemoveDealingFValidator(ctx, p.Denom, p.PoolAddress)
	k.SetLatestDealedCycle(ctx, p.Cycle)
	return nil
}
