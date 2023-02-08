package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/furybase/furybase/x/fvalidator/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		sudoKeeper   types.SudoKeeper
		FBank Keeper  types.FBank Keeper
		ledgerKeeper types.LedgerKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	sudoKeeper types.SudoKeeper,
	fBank Keeper types.FBank Keeper,
	ledgerKeeper types.LedgerKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:          cdc,
		storeKey:     storeKey,
		memKey:       memKey,
		paramstore:   ps,
		sudoKeeper:   sudoKeeper,
		FBank Keeper:  fBank Keeper,
		ledgerKeeper: ledgerKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) AddSelectedFValidator(ctx sdk.Context, fValidator *types.FValidator) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SelectedRValdidatorStoreKey(fValidator.Denom, fValidator.PoolAddress, fValidator.ValAddress), []byte{})
}

func (k Keeper) RemoveSelectedFValidator(ctx sdk.Context, fValidator *types.FValidator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SelectedRValdidatorStoreKey(fValidator.Denom, fValidator.PoolAddress, fValidator.ValAddress))
}

func (k Keeper) HasSelectedFValidator(ctx sdk.Context, fValidator *types.FValidator) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SelectedRValdidatorStoreKey(fValidator.Denom, fValidator.PoolAddress, fValidator.ValAddress))
}

func (k Keeper) GetSelectedFValidatorListByDenomPoolAddress(ctx sdk.Context, denom, poolAddress string) []*types.FValidator {
	store := ctx.KVStore(k.storeKey)
	denomLen := len([]byte(denom))
	poolAddressLen := len([]byte(poolAddress))

	key := make([]byte, 1+1+denomLen+1+poolAddressLen)
	copy(key[0:], types.SelectedFValidatorStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[2+denomLen] = byte(poolAddressLen)
	copy(key[2+denomLen+1:], []byte(poolAddress))

	iterator := sdk.KVStorePrefixIterator(store, key)
	defer iterator.Close()

	list := make([]*types.FValidator, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		address := string(key[2+denomLen+1+poolAddressLen+1:])

		fValidator := types.FValidator{
			Denom:       denom,
			PoolAddress: poolAddress,
			ValAddress:  address,
		}

		list = append(list, &fValidator)
	}
	return list
}

func (k Keeper) GetSelectedFValidatorList(ctx sdk.Context) []*types.FValidator {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SelectedFValidatorStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.FValidator, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		denomLen := int(key[1])
		denom := string(key[2 : 2+denomLen])
		poolAddressLen := int(key[2+denomLen])
		poolAddress := string(key[2+denomLen+1 : 2+denomLen+1+poolAddressLen])
		valAddress := string(key[2+denomLen+1+poolAddressLen+1:])

		fValidator := types.FValidator{
			Denom:       denom,
			PoolAddress: poolAddress,
			ValAddress:  valAddress,
		}

		list = append(list, &fValidator)
	}
	return list
}

func (k Keeper) SetLatestVotedCycle(ctx sdk.Context, cycle *types.Cycle) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LatestVotedCycleStoreKey(cycle.Denom, cycle.PoolAddress), k.cdc.MustMarshal(cycle))
}

func (k Keeper) GetLatestVotedCycle(ctx sdk.Context, denom, poolAddress string) *types.Cycle {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.LatestVotedCycleStoreKey(denom, poolAddress))
	if bts == nil {
		return &types.Cycle{
			Denom:       denom,
			PoolAddress: poolAddress,
			Version:     0,
			Number:      0,
		}
	}
	cycle := types.Cycle{}
	k.cdc.MustUnmarshal(bts, &cycle)

	return &cycle
}

func (k Keeper) GetAllLatestVotedCycle(ctx sdk.Context) []*types.Cycle {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.LatestVotedCycleStoreKeyPrefix)
	defer iterator.Close()

	cycleList := make([]*types.Cycle, 0)
	for ; iterator.Valid(); iterator.Next() {
		cycle := types.Cycle{}
		k.cdc.MustUnmarshal(iterator.Value(), &cycle)
		cycleList = append(cycleList, &cycle)
	}

	return cycleList
}

func (k Keeper) SetLatestDealedCycle(ctx sdk.Context, cycle *types.Cycle) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LatestDealedCycleStoreKey(cycle.Denom, cycle.PoolAddress), k.cdc.MustMarshal(cycle))
}

func (k Keeper) GetLatestDealedCycle(ctx sdk.Context, denom, poolAddress string) *types.Cycle {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.LatestDealedCycleStoreKey(denom, poolAddress))
	if bts == nil {
		return &types.Cycle{
			Denom:       denom,
			PoolAddress: poolAddress,
			Version:     0,
			Number:      0,
		}
	}
	cycle := types.Cycle{}
	k.cdc.MustUnmarshal(bts, &cycle)

	return &cycle
}

func (k Keeper) GetAllLatestDealedCycle(ctx sdk.Context) []*types.Cycle {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.LatestDealedCycleStoreKeyPrefix)
	defer iterator.Close()

	cycleList := make([]*types.Cycle, 0)
	for ; iterator.Valid(); iterator.Next() {
		cycle := types.Cycle{}
		k.cdc.MustUnmarshal(iterator.Value(), &cycle)
		cycleList = append(cycleList, &cycle)
	}

	return cycleList
}

func (k Keeper) SetCycleSeconds(ctx sdk.Context, cycleSeconds *types.CycleSeconds) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.CycleSecondsStoreKey(cycleSeconds.Denom), k.cdc.MustMarshal(cycleSeconds))
}

func (k Keeper) GetCycleSeconds(ctx sdk.Context, denom string) *types.CycleSeconds {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.CycleSecondsStoreKey(denom))
	if bts == nil {
		return &types.CycleSeconds{
			Denom:   denom,
			Version: 0,
			Seconds: types.DefaultCycleSeconds,
		}
	}

	cycleSeconds := types.CycleSeconds{}
	k.cdc.MustUnmarshal(bts, &cycleSeconds)
	return &cycleSeconds
}

func (k Keeper) GetAllCycleSeconds(ctx sdk.Context) []*types.CycleSeconds {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.CycleSecondsStoreKeyPrefix)
	defer iterator.Close()

	cycleSecondsList := make([]*types.CycleSeconds, 0)
	for ; iterator.Valid(); iterator.Next() {
		cycleSeconds := types.CycleSeconds{}
		k.cdc.MustUnmarshal(iterator.Value(), &cycleSeconds)
		cycleSecondsList = append(cycleSecondsList, &cycleSeconds)
	}
	return cycleSecondsList
}

func (k Keeper) SetShuffleSeconds(ctx sdk.Context, shuffleSeconds *types.ShuffleSeconds) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ShuffleSecondsStoreKey(shuffleSeconds.Denom), k.cdc.MustMarshal(shuffleSeconds))
}

func (k Keeper) GetShuffleSeconds(ctx sdk.Context, denom string) *types.ShuffleSeconds {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ShuffleSecondsStoreKey(denom))
	if bts == nil {
		return &types.ShuffleSeconds{
			Denom:   denom,
			Version: 0,
			Seconds: types.DefaultShuffleSeconds,
		}
	}
	shuffleSeconds := types.ShuffleSeconds{}

	k.cdc.MustUnmarshal(bts, &shuffleSeconds)

	return &shuffleSeconds
}

func (k Keeper) GetAllShuffleSeconds(ctx sdk.Context) []*types.ShuffleSeconds {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ShuffleSecondsStoreKeyPrefix)
	defer iterator.Close()

	shuffleSecondsList := make([]*types.ShuffleSeconds, 0)
	for ; iterator.Valid(); iterator.Next() {
		shuffleSeconds := types.ShuffleSeconds{}
		k.cdc.MustUnmarshal(iterator.Value(), &shuffleSeconds)
		shuffleSecondsList = append(shuffleSecondsList, &shuffleSeconds)
	}

	return shuffleSecondsList
}

func (k Keeper) SetDealingFValidator(ctx sdk.Context, fValidator *types.DealingFValidator) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.DealingFValidatorStoreKey(fValidator.Denom, fValidator.PoolAddress), k.cdc.MustMarshal(fValidator))
}

func (k Keeper) RemoveDealingFValidator(ctx sdk.Context, denom, poolAddress string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DealingFValidatorStoreKey(denom, poolAddress))
}

func (k Keeper) GetDealingFValidator(ctx sdk.Context, denom, poolAddress string) (*types.DealingFValidator, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.DealingFValidatorStoreKey(denom, poolAddress))
	if bts == nil {
		return nil, false
	}
	fvalidator := types.DealingFValidator{}
	k.cdc.MustUnmarshal(bts, &fvalidator)

	return &fvalidator, true
}

func (k Keeper) GetAllDealingFvalidators(ctx sdk.Context) []*types.DealingFValidator {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DealingFValidatorStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.DealingFValidator, 0)
	for ; iterator.Valid(); iterator.Next() {
		dealingFvalidator := types.DealingFValidator{}
		k.cdc.MustUnmarshal(iterator.Value(), &dealingFvalidator)
		list = append(list, &dealingFvalidator)
	}
	return list
}
