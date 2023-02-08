package fmintreward_test

import (
	"testing"

	keepertest "github.com/furybase/furybase/testutil/keeper"
	"github.com/furybase/furybase/testutil/nullify"
	"github.com/furybase/furybase/x/fmintreward"
	"github.com/furybase/furybase/x/fmintreward/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FmintrewardKeeper(t)
	fmintreward.InitGenesis(ctx, *k, genesisState)
	got := fmintreward.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
