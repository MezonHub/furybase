package keeper

import (
	"github.com/furybase/furybase/x/claim/types"
)

var _ types.QueryServer = Keeper{}
