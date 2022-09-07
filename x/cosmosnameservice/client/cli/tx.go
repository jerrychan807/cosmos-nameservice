package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/types"
)

// GetTxCmd returns the transaction commands for this module
// 返回模块的交易命令
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	nameserviceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// 把txWhois.go中的命令都添加
	nameserviceTxCmd.AddCommand(flags.PostCommands(
		// this line is used by starport scaffolding
		GetCmdBuyName(cdc),
		GetCmdSetWhois(cdc),
		GetCmdDeleteWhois(cdc),
	)...)

	return nameserviceTxCmd
}
