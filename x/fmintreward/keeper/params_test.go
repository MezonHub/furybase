package keeper_test

import (
	"testing"

	testkeeper "github.com/furybase/furybase/testutil/keeper"
	"github.com/furybase/furybase/x/fmintreward/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FmintrewardKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
