package cli

import (
	"fmt"
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

func CmdUpdateFValidatorReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-r-validator-report [denom] [pool-address] [cycle-version] [cycle-number] [status]",
		Short: "Update fvalidator report",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPoolAddress := args[1]
			argCycleVersion, err := sdk.ParseUint(args[2])
			if err != nil {
				return err
			}
			argCycleNumber, err := sdk.ParseUint(args[3])
			if err != nil {
				return err
			}
			status, exist := types.UpdateFValidatorStatus_value[args[4]]
			if !exist {
				return fmt.Errorf("status not exist")
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cycle := types.Cycle{
				Denom:       argDenom,
				PoolAddress: argPoolAddress,
				Version:     argCycleVersion.Uint64(),
				Number:      argCycleNumber.Uint64(),
			}
			content := types.NewUpdateFValidatorReportProposal(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argPoolAddress,
				&cycle, types.UpdateFValidatorStatus(status))
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
