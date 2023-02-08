package fbank _test

import (
	"testing"

	keepertest "github.com/furybase/furybase/testutil/keeper"
	"github.com/furybase/furybase/testutil/nullify"
	"github.com/furybase/furybase/x/fbank "
	"github.com/furybase/furybase/x/fbank /types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Fbank Keeper(t)
	fbank .InitGenesis(ctx, *k, genesisState)
	got := fbank .ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
