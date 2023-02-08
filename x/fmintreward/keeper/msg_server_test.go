package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/furyunderverse/furybase/testutil/keeper"
	"github.com/furyunderverse/furybase/x/fmintreward/keeper"
	"github.com/furyunderverse/furybase/x/fmintreward/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FmintrewardKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
