package fvalidator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyunderverse/furybase/x/fvalidator/keeper"
	"github.com/furyunderverse/furybase/x/fvalidator/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	for _, cycleSeconds := range genState.CycleSecondsList {
		k.SetCycleSeconds(ctx, cycleSeconds)
	}
	for _, latestDealedCycle := range genState.LatestDealedCycleList {
		k.SetLatestDealedCycle(ctx, latestDealedCycle)
	}

	for _, latestVotedCycle := range genState.LatestVotedCycleList {
		k.SetLatestVotedCycle(ctx, latestVotedCycle)
	}
	for _, selectedFValidator := range genState.SelectedFValidatorList {
		if err := k.FBank Keeper.CheckValAddress(ctx, selectedFValidator.Denom, selectedFValidator.ValAddress); err != nil {
			panic(err)
		}
		if err := k.FBank Keeper.CheckAccAddress(ctx, selectedFValidator.Denom, selectedFValidator.PoolAddress); err != nil {
			panic(err)
		}
		k.AddSelectedFValidator(ctx, selectedFValidator)
	}

	for _, shuffleSeconds := range genState.ShuffleSecondsList {
		k.SetShuffleSeconds(ctx, shuffleSeconds)
	}

	for _, dealingFValidator := range genState.DealingFValidatorList {
		if err := k.FBank Keeper.CheckValAddress(ctx, dealingFValidator.Denom, dealingFValidator.NewValAddress); err != nil {
			panic(err)
		}
		if err := k.FBank Keeper.CheckValAddress(ctx, dealingFValidator.Denom, dealingFValidator.OldValAddress); err != nil {
			panic(err)
		}
		if err := k.FBank Keeper.CheckAccAddress(ctx, dealingFValidator.Denom, dealingFValidator.PoolAddress); err != nil {
			panic(err)
		}
		k.SetDealingFValidator(ctx, dealingFValidator)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.CycleSecondsList = k.GetAllCycleSeconds(ctx)
	genesis.LatestDealedCycleList = k.GetAllLatestDealedCycle(ctx)
	genesis.LatestVotedCycleList = k.GetAllLatestVotedCycle(ctx)
	genesis.SelectedFValidatorList = k.GetSelectedFValidatorList(ctx)
	genesis.ShuffleSecondsList = k.GetAllShuffleSeconds(ctx)
	genesis.DealingFValidatorList = k.GetAllDealingFvalidators(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
