package keeper

import (
	"github.com/furybase/furybase/x/fvalidator/types"
)

var _ types.QueryServer = Keeper{}
