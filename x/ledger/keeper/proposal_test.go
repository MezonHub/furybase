package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/furyunderverse/furybase/testutil/keeper"
	"github.com/furyunderverse/furybase/x/ledger/keeper"
	//"github.com/furyunderverse/furybase/x/ledger/types"
)

func setupSettings(t testing.TB) {
	k, ctx := keepertest.LedgerKeeper(t)
	s, _ := keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
	t.Log(s)
}
