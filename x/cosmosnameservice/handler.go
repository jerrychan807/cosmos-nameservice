package cosmosnameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/keeper"
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/types"
)

// NewHandler returns a handler for "nameservice" type messages.
// 返回一个操作nameservice各类消息的Handler对象
func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		//.(type)获取接口实例实际的类型指针, 以此调用实例所有可调用的方法，包括接口方法及自有方法。
		//需要注意的是该写法必须与switch case联合使用，case中列出实现该接口的类型。
		switch msg := msg.(type) {
		// 添加操作类型
		case types.MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case types.MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case types.MsgDeleteName:
			return handleMsgDeleteName(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type()))
		}
	}
}
