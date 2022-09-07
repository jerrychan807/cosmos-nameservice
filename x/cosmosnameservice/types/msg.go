package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// import (
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
// )

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`
/*
// verify interface at compile time
var _ sdk.Msg = &Msg<Action>{}

// Msg<Action> - struct for unjailing jailed validator
type Msg<Action> struct {
	ValidatorAddr sdk.ValAddress `json:"address" yaml:"address"` // address of the validator operator
}

// NewMsg<Action> creates a new Msg<Action> instance
func NewMsg<Action>(validatorAddr sdk.ValAddress) Msg<Action> {
	return Msg<Action>{
		ValidatorAddr: validatorAddr,
	}
}

const <action>Const = "<action>"

// nolint
func (msg Msg<Action>) Route() string { return RouterKey }
func (msg Msg<Action>) Type() string  { return <action>Const }
func (msg Msg<Action>) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg Msg<Action>) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg Msg<Action>) ValidateBasic() error {
	if msg.ValidatorAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address")
	}
	return nil
}
*/

// Transactions messages must fulfill the Msg
// 实现接口必须实现所有的函数
type Msg interface {
	// Return the message type.
	// Must be alphanumeric or empty.
	// 返回消息类型， 必须是字母或者空
	Type() string

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	// 返回消息的可读string,用于在标签中使用
	Route() string

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	// 做一些基本的验证, 不需要访问其他信息
	ValidateBasic() error

	// Get the canonical byte representation of the Msg.
	// 获取Msg的规范字节表示即字节数组
	GetSignBytes() []byte

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	// 返回必须签名的签名者地址集合， 所有的签名必须在当下还有效， 返回的签名集合会以某种确定的顺序
	GetSigners() []sdk.AccAddress
}
