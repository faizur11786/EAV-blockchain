package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdNewMerchant() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-merchant [address] [attributes]",
		Short: "Broadcast message new-merchant",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// argAddress := args[0]
			// argAttributes := new(types.Attribute)
			// err = json.Unmarshal([]byte(args[1]), argAttributes)
			// if err != nil {
			// 	return err
			// }

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// msg := types.NewMsgNewMerchant(
			// 	clientCtx.GetFromAddress().String(),
			// 	argAddress,
			// 	argAttributes,
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
