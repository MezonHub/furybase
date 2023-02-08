package keeper

import (
	"github.com/furybase/furybase/x/mining/types"
)

var _ types.QueryServer = Keeper{}
