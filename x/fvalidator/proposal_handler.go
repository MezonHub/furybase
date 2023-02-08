package fvalidator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/furybase/furybase/x/fvalidator/keeper"
	"github.com/furybase/furybase/x/fvalidator/types"
	fvotetypes "github.com/furybase/furybase/x/fvote/types"
)

func NewProposalHandler(k keeper.Keeper) fvotetypes.Handler {
	return func(ctx sdk.Context, content fvotetypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateFValidatorProposal:
			return k.ProcessUpdateFValidatorProposal(ctx, c)
		case *types.UpdateFValidatorReportProposal:
			return k.ProcessUpdateFValidatorReportProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}
