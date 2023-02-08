package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/furybase/furybase/x/fvalidator/types"
	fvotetypes "github.com/furybase/furybase/x/fvote/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateFValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-r-validator [denom] [pool-address] [old-address] [new-address] [cycleVersion] [cycleNumber]",
		Short: "Update fvalidator",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPoolAddress := args[1]
			argOldAddress := args[2]
			argNewAddress := args[3]
			argCycleVersion, err := sdk.ParseUint(args[4])
			if err != nil {
				return err
			}
			argCycleNumber, err := sdk.ParseUint(args[5])
			if err != nil {
				return err
			}

			cycle := types.Cycle{
				Denom:       argDenom,
				PoolAddress: argPoolAddress,
				Version:     argCycleVersion.Uint64(),
				Number:      argCycleNumber.Uint64(),
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			content := types.NewUpdateFValidatorProposal(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argPoolAddress,
				argOldAddress,
				argNewAddress,
				&cycle)
			msg, err := fvotetypes.NewMsgSubmitProposal(clientCtx.GetFromAddress(), content)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
