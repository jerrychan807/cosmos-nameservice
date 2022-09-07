package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
// codec上注册具体的类型
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgBuyName{}, "cosmosnameservice/BuyName", nil)
	cdc.RegisterConcrete(MsgSetName{}, "cosmosnameservice/SetName", nil)
	cdc.RegisterConcrete(MsgDeleteName{}, "cosmosnameservice/DeleteName", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	// 创建实例codec
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	// Register the go-crypto to the codec
	// 注册加密函数信息
	codec.RegisterCrypto(ModuleCdc)
	// 封装
	ModuleCdc.Seal()
}
