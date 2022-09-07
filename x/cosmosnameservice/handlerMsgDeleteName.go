package cosmosnameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/keeper"
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/types"
)

// Handle a message to delete name
// 删除
func handleMsgDeleteName(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteName) (*sdk.Result, error) {
	// 使用id检查是否存在该域名
	if !k.WhoisExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	// 检查是否是本人
	if !msg.Creator.Equals(k.GetWhoisOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	k.DeleteWhois(ctx, msg.ID)
	return &sdk.Result{}, nil
}
