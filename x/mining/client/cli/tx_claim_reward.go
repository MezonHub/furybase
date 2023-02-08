package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/furyunderverse/furybase/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdClaimReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-reward [stake-pool-index] [index]",
		Short: "Claim reward",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStakePoolIndex, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}
			argIndex, err := sdk.ParseUint(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaimReward(
				clientCtx.GetFromAddress().String(),
				uint32(argStakePoolIndex.Uint64()),
				uint32(argIndex.Uint64()),
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
