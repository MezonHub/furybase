package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furybase/furybase/x/fvalidator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FValidatorList(goCtx context.Context, req *types.QueryFValidatorListRequest) (*types.QueryFValidatorListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	list := make([]string, 0)
	for _, val := range k.GetSelectedFValidatorListByDenomPoolAddress(ctx, req.Denom, req.PoolAddress) {
		list = append(list, val.ValAddress)
	}

	return &types.QueryFValidatorListResponse{
		FValidatorList: list,
	}, nil
}
