package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//定义最小域名转卖金额
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

type Whois struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Value   string         `json:"value" yaml:"value"`
	Price   sdk.Coins      `json:"price"`
}

//返回一个新的Whois，因为新的域名还没有人购买，所以设置金额为最小
func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}
