package cosmosnameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/keeper"
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/types"
)

// Handle a message to set name
// 接受到msg， 进一步的处理，相当于controller层， 操作keeper层
func handleMsgSetName(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgSetName) (*sdk.Result, error) {
	// 检查MsgSetName的提供者是否为想要设置域名的域名拥有者， 即验证身份
	if !msg.Owner.Equals(keeper.GetWhoisOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}
	// 如果是本人的话设置所规定的值
	// 调用keeper进行域名的解析值设定
	keeper.SetName(ctx, msg.Name, msg.Value) // If so, set the name to the value specified in the msg.
	return &sdk.Result{}, nil                // return
}
