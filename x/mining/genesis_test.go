package mining_test

import (
	"testing"

	keepertest "github.com/furyunderverse/furybase/testutil/keeper"
	"github.com/furyunderverse/furybase/testutil/nullify"
	"github.com/furyunderverse/furybase/x/mining"
	"github.com/furyunderverse/furybase/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.DefaultGenesis()

	k, ctx := keepertest.MiningKeeper(t)
	mining.InitGenesis(ctx, *k, *genesisState)
	got := mining.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
