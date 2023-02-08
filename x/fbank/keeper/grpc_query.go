package keeper

import (
	"github.com/furyunderverse/furybase/x/fbank /types"
)

var _ types.QueryServer = Keeper{}
