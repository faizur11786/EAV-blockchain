package cli

import (
	"context"

	"belShare/x/eav/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListAttribute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-attribute",
		Short: "list all attribute",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAttributeRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AttributeAll(context.Background(), params)
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

func CmdShowAttribute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-attribute [guid]",
		Short: "shows a attribute",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argGuid := args[0]

			params := &types.QueryGetAttributeRequest{
				Guid: argGuid,
			}

			res, err := queryClient.Attribute(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
