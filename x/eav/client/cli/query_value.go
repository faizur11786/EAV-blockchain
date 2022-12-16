package cli

import (
	"context"

	"belShare/x/eav/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-value",
		Short: "list all value",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllValueRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ValueAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-value [entity-type-id] [entity-id] [attribute-id]",
		Short: "shows a value",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argEntityId := args[1]
			argAttributeId := args[2]

			params := &types.QueryGetValueRequest{
				EntityId:    argEntityId,
				AttributeId: argAttributeId,
			}

			res, err := queryClient.Value(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
