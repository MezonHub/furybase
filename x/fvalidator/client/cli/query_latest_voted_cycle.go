package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/furybase/furybase/x/fvalidator/types"
)

var _ = strconv.Itoa(0)

func CmdLatestVotedCycle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "latest-voted-cycle [denom] [pool-address]",
		Short: "Query latest voted cycle",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]
			reqPoolAddress := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryLatestVotedCycleRequest{
				Denom:       reqDenom,
				PoolAddress: reqPoolAddress,
			}

			res, err := queryClient.LatestVotedCycle(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
