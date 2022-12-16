package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateShop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-shop [address] [attributs]",
		Short: "Broadcast message create-shop",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// argAddress := args[0]
			// argAttributs  := args[1]
			// err = json.Unmarshal([]byte(args[1]), argAttributs)
			// if err != nil {
			// 	return err
			// }

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// msg := types.NewMsgCreateShop(
			// 	clientCtx.GetFromAddress().String(),
			// 	argAddress,
			// 	argAttributs,
			// )
			// if err := msg.ValidateBasic(); err != nil {
			// 	return err
			// }
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), nil)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
