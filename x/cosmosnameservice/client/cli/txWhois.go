package cli

import (
	"bufio"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/types"
)

// 购买新域名
func GetCmdBuyName(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "buy-name [name] [price]",
		Short: "Buys a new name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0]) //获取购买的名字name

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			// 获取标准读入
			inBuf := bufio.NewReader(cmd.InOrStdin())
			// 创建交易创建器
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			coins, err := sdk.ParseCoins(args[1]) //解析出价
			if err != nil {
				return err
			}
			// 构建NewMsgBuyName实例
			msg := types.NewMsgBuyName(argsName, coins, cliCtx.GetFromAddress())
			// 做基本的验证
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			// 生成或广播交易
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// 设置域名解析
func GetCmdSetWhois(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-name [value] [name]",
		Short: "Set a new name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsValue := args[0]
			argsName := args[1]

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetName(argsName, argsValue, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// 删除一个域名
func GetCmdDeleteWhois(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-name [id]",
		Short: "Delete a new name by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteName(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
