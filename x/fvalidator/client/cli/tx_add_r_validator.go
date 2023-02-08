package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/furybase/furybase/x/fvalidator/types"
)

var _ = strconv.Itoa(0)

func CmdAddFValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-r-validator [denom] [pool-address] [val-address]",
		Short: "Add fvalidator",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPoolAddress := args[1]
			argValAddress := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddFValidator(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argPoolAddress,
				argValAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
