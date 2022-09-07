package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Originally, this file was named MsgCreateWhois, and has been modified using search-and-replace to our Msg needs.
// 根据MsgCreateWhois文件改写
// MsgBuyName defines the BuyName message
type MsgBuyName struct {
	Name  string         `json:"name"`  // 想购买的域名
	Bid   sdk.Coins      `json:"bid"`   // 出价
	Buyer sdk.AccAddress `json:"buyer"` // 购买者
}

// NewMsgBuyName is the constructor function for MsgBuyName
// MsgBuyName构造函数
func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

// Route should return the name of the module
// 路由返回模块名nameService
func (msg MsgBuyName) Route() string { return RouterKey }

// Type should return the action
// 操作类型
func (msg MsgBuyName) Type() string { return "buy_name" }

// ValidateBasic runs stateless checks on the message
// 基本的检查
func (msg MsgBuyName) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
// 返回消息的编码后格式[]byte
func (msg MsgBuyName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
// 要求签名的对象
func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}
