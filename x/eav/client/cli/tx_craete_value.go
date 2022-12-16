package cli

import (
	"strconv"

	"belShare/x/eav/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCraeteValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "craete-value [attribute-id] [entity-id] [value]",
		Short: "Broadcast message craete-value",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAttributeId := args[0]
			argEntityId := args[1]
			argValue := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCraeteValue(
				clientCtx.GetFromAddress().String(),
				argAttributeId,
				argEntityId,
				argValue,
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
