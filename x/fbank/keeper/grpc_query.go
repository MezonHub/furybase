package keeper

import (
	"github.com/furybase/furybase/x/fbank /types"
)

var _ types.QueryServer = Keeper{}
