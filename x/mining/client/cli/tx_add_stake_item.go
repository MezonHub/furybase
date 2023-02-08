package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/furybase/furybase/utils"
	"github.com/furybase/furybase/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdAddStakeItem() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-stake-item [lock-second] [power-reward-rate] [enable(true/false)]",
		Short: "Add stake item",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argLockSecond, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}
			argPowerRewardRate, err := utils.NewDecFromStr(args[1])
			if err != nil {
				return err
			}

			enable := true
			if args[2] == "false" {
				enable = false
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddStakeItem(
				clientCtx.GetFromAddress().String(),
				argLockSecond.Uint64(),
				argPowerRewardRate,
				enable,
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
