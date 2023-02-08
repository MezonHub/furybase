package keeper

import (
	"github.com/furybase/furybase/x/fdex/types"
)

var _ types.QueryServer = Keeper{}
