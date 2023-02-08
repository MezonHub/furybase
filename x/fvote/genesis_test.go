package fvote_test

import (
	"testing"

	keepertest "github.com/furybase/furybase/testutil/keeper"
	"github.com/furybase/furybase/x/fvote"
	"github.com/furybase/furybase/x/fvote/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FvoteKeeper(t)
	fvote.InitGenesis(ctx, *k, genesisState)
	got := fvote.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
