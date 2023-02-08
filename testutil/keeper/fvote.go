package keeper

import (
	"sync"
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ledgermodule "github.com/furyunderverse/furybase/x/ledger"
	ledgertypes "github.com/furyunderverse/furybase/x/ledger/types"
	"github.com/furyunderverse/furybase/x/fvote/keeper"
	"github.com/furyunderverse/furybase/x/fvote/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	fvoteStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	fvoteMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	fvoteOnce        sync.Once
)

func FvoteKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	fvoteOnce.Do(func() {
		stateStore.MountStoreWithDB(fvoteStoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(fvoteMemStoreKey, sdk.StoreTypeMemory, nil)
	})

	sudoKeeper, _ := SudoKeeper(t)
	relayersKeeper, _ := RelayersKeeper(t)
	ledgerKeeper, _ := LedgerKeeper(t)
	require.NoError(t, stateStore.LoadLatestVersion())

	fvoteRouter := types.NewRouter()
	fvoteRouter.AddRoute(ledgertypes.RouterKey, ledgermodule.NewProposalHandler(ledgerKeeper))
	fvoteKeeper := keeper.NewKeeper(
		cdc,
		fvoteStoreKey,
		fvoteMemStoreKey,

		sudoKeeper,
		relayersKeeper,
		fvoteRouter,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return fvoteKeeper, ctx
}
