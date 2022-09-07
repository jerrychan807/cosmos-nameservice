package keeper

import (
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/types"
	// this line is used by starport scaffolding # 1
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for nameservice clients.
// 创建了一个keeper层的查询对象给客户端
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] { // 根据客户端输入的路径的第一个变量，确定查询的类型
		// this line is used by starport scaffolding # 2
		// 客户端输入的内容在types/querier.go中定义了常量作为路由
		case types.QueryResolveName:
			return resolveName(ctx, path[1:], k)
		case types.QueryListWhois:
			return listWhois(ctx, k)
		case types.QueryGetWhois:
			return getWhois(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown nameservice query endpoint")
		}
	}
}
