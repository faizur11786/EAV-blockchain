package cli

import (
	"strconv"

	"belShare/x/eav/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEntityAttributes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "entity-attributes [entity-id]",
		Short: "Query entity-attributes",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqEntityId := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEntityAttributesRequest{

				EntityId: reqEntityId,
			}

			res, err := queryClient.EntityAttributes(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
