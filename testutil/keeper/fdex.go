package keeper

import (
	"sync"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/furyunderverse/furybase/x/fdex/keeper"
	"github.com/furyunderverse/furybase/x/fdex/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	fdexStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	fdexMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	fdexOnce        sync.Once
)

func FdexKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	fdexOnce.Do(func() {
		stateStore.MountStoreWithDB(fdexStoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(fdexMemStoreKey, sdk.StoreTypeMemory, nil)
	})
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		fdexStoreKey,
		fdexMemStoreKey,
		"FdexParams",
	)
	sudoKeeper, _ := SudoKeeper(t)
	k := keeper.NewKeeper(
		cdc,
		fdexStoreKey,
		fdexMemStoreKey,
		paramsSubspace,
		BankKeeper,
		sudoKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
