package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

//const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetName defines a SetName message
// 定义MsgSetName的结构
type MsgSetName struct {
	Name  string         `json:"name"`  // 目标域名
	Value string         `json:"value"` // 对应的值/解析值
	Owner sdk.AccAddress `json:"owner"` // 拥有者
}

// NewMsgSetName is a constructor function for MsgSetName
// MsgSetName构造函数
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

// Route should return the name of the module
// 返回路由消息的键， 这里就是模块名
func (msg MsgSetName) Route() string { return RouterKey }

// Type should return the action
// 操作名
func (msg MsgSetName) Type() string { return "set_name" }

// ValidateBasic runs stateless checks on the message
// 基本参数检测
func (msg MsgSetName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
// GetSignBytes对整个MsgSetName消息本身进行编码、排序以进行后续的签名
func (msg MsgSetName) GetSignBytes() []byte {
	// MustMarshalJSON序列化msg为字节切片
	// MustSortJSON返回根据key排序的json
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
// 定义该需要谁的签名
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
