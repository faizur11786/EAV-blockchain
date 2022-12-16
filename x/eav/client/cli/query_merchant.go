package cli

import (
	"context"

	"belShare/x/eav/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListMerchant() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-merchant",
		Short: "list all merchant",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMerchantRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.MerchantAll(context.Background(), params)
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

func CmdShowMerchant() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-merchant [guid] [creator]",
		Short: "shows a merchant",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argGuid := args[0]
			argCreator := args[1]

			params := &types.QueryGetMerchantRequest{
				Guid:    argGuid,
				Creator: argCreator,
			}

			res, err := queryClient.Merchant(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
