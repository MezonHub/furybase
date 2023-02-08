package keeper

import (
	"github.com/furybase/furybase/x/fstaking/types"
)

var _ types.QueryServer = Keeper{}
