package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/furybase/furybase/x/claim/types"
)

var _ = strconv.Itoa(0)

func CmdSetMerkleRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-merkle-root [round] [merkle-root]",
		Short: "Set merkle root ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMerkleRound, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}
			argMerkleRoot := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetMerkleRoot(
				clientCtx.GetFromAddress().String(),
				argMerkleRound.Uint64(),
				argMerkleRoot,
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
