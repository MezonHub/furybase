package keeper_test

import (
	"testing"

	testkeeper "github.com/furyunderverse/furybase/testutil/keeper"
	"github.com/furyunderverse/furybase/x/fbank /types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Fbank Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}