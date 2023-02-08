package sudo_test

import (
	"testing"

	keepertest "github.com/furybase/furybase/testutil/keeper"
	"github.com/furybase/furybase/testutil/sample"
	"github.com/furybase/furybase/x/sudo"
	"github.com/furybase/furybase/x/sudo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.DefaultGenesis()
	genesisState.Admin = sample.TestAdmin

	k, ctx := keepertest.SudoKeeper(t)
	sudo.InitGenesis(ctx, k, *genesisState)
	got := sudo.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
