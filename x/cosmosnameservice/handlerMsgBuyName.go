package cosmosnameservice

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/keeper"
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/types"
)

// Handle a message to buy name
func handleMsgBuyName(ctx sdk.Context, k keeper.Keeper, msg types.MsgBuyName) (*sdk.Result, error) {
	// Checks if the the bid price is greater than the price paid by the current owner
	// 1.检查当前出价是否高于目前的价格, 注意Msg本身的检查只是简单的检查，这里需要额外数据的检查就只能在Handler中做
	// GetPrice返回coin类型对象，其IsAllGT函数是比较大小（逐个字母比较，全部大于返回true）
	// 当需要的价格 > bid那么就返回错误
	if k.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough") // If not, throw an error
	}
	// 2.检查当前域名是否已经有拥有者了
	// 不论是已拥有或者没有人拥有， 如果购买者支付出价出现错误，那么都会造成资金的回滚
	if k.HasCreator(ctx, msg.Name) {
		// 如果已经是别人拥有的，那么购买者支付对应的出价给域名原来的拥有者
		// coin转移方向： msg.Buyer => Creator
		// 金额： Bid
		err := k.CoinKeeper.SendCoins(ctx, msg.Buyer, k.GetCreator(ctx, msg.Name), msg.Bid)
		if err != nil {
			return nil, err
		}
	} else {
		// 如果没有，那么从购买者处减去出价金额, 发送给一个不可回收的地址（burns）
		_, err := k.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return nil, err
		}
	}
	// 分别为域名设置新的所有者与金额
	k.SetCreator(ctx, msg.Name, msg.Buyer)
	k.SetPrice(ctx, msg.Name, msg.Bid)
	return &sdk.Result{}, nil
}
