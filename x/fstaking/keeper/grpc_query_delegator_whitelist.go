package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furybase/furybase/x/fstaking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DelegatorWhitelist(goCtx context.Context, req *types.QueryDelegatorWhitelistRequest) (*types.QueryDelegatorWhitelistResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryDelegatorWhitelistResponse{
		DelegatorAddress: k.GetDelegatorAddressWhitelist(ctx),
	}, nil
}
