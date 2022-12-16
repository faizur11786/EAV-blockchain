package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"belShare/x/eav/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group eav queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListUser())
	cmd.AddCommand(CmdShowUser())
	cmd.AddCommand(CmdListMerchant())
	cmd.AddCommand(CmdShowMerchant())
	cmd.AddCommand(CmdListEntityType())
	cmd.AddCommand(CmdShowEntityType())
	cmd.AddCommand(CmdListAttribute())
	cmd.AddCommand(CmdShowAttribute())
	cmd.AddCommand(CmdListValue())
	cmd.AddCommand(CmdShowValue())
	// this line is used by starport scaffolding # 1

	return cmd
}
