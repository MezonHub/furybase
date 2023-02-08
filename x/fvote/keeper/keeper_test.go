package keeper_test

import (
	"testing"

	testkeeper "github.com/furyunderverse/furybase/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestKeeper_ProposalLife(t *testing.T) {
	k, ctx := testkeeper.FvoteKeeper(t)

	pl := k.ProposalLife(ctx)
	require.Equal(t, int64(0), pl)

	testProposalLife := int64(100)
	k.SetProposalLife(ctx, testProposalLife)
	pl = k.ProposalLife(ctx)
	require.Equal(t, testProposalLife, pl)
}
