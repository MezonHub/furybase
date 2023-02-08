package keeper

import (
	"sync"
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	"github.com/furyunderverse/furybase/x/ledger/keeper"
	"github.com/furyunderverse/furybase/x/ledger/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	ledgertoreKey     = sdk.NewKVStoreKey(types.StoreKey)
	ledgerMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	ledgerOnce        sync.Once
)

func LedgerKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	ledgerOnce.Do(func() {
		stateStore.MountStoreWithDB(ledgertoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(ledgerMemStoreKey, sdk.StoreTypeMemory, nil)
	})
	require.NoError(t, stateStore.LoadLatestVersion())

	sudoKeeper, _ := SudoKeeper(t)
	relayersKeeper, _ := RelayersKeeper(t)
	fmintRewardKeeper, _ := FmintrewardKeeper(t)
	fBank Keeper, _ := Fbank Keeper(t)

	ledgerKeeper := keeper.NewKeeper(
		cdc,
		ledgertoreKey,
		ledgerMemStoreKey,
		sudoKeeper,
		BankKeeper,
		relayersKeeper,
		fmintRewardKeeper,
		fBank Keeper,
		//todo impl keepers below
		icacontrollerkeeper.Keeper{},
		capabilitykeeper.ScopedKeeper{},
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return *ledgerKeeper, ctx
}
