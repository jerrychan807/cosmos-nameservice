package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/jerrychan807/cosmos-nameservice/x/cosmosnameservice/types"
)

// GetWhoisCount get the total number of whois
// 获取whois的所有数量
func (k Keeper) GetWhoisCount(ctx sdk.Context) int64 {
	// 获取keeper对象的key,也就是在key.go文件中定义的存储StoreKey常量
	// 提取上下文中对应key的store
	store := ctx.KVStore(k.storeKey)
	// 将key.go中的数量统计前缀转换为byte数组
	byteKey := []byte(types.WhoisCountPrefix)
	// 根据WhoisCountPrefix查询存储的数量值
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	// 不存在则返回默认的0
	if bz == nil {
		return 0
	}

	// Parse bytes
	// 解析字查询到的数量string => int
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetWhoisCount set the total number of whois
// 存储whois的总数
func (k Keeper) SetWhoisCount(ctx sdk.Context, count int64) {
	// 获取数据库 => 格式化key => 存储
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.WhoisCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	// 设置
	store.Set(byteKey, bz)
}

// CreateWhois creates a whois. This function is included in starport type scaffolding.
// We won't use this function in our application, so it can be commented out.
// 脚手架自动创建的创建whois结构体的函数, 不想使用可以注释掉, 一般写在Tpye{your type}.go中,这里就是TypeWhois.go
// func (k Keeper) CreateWhois(ctx sdk.Context, whois types.Whois) {
// 	store := ctx.KVStore(k.storeKey)
// 	key := []byte(types.WhoisPrefix + whois.Value)
// 	value := k.cdc.MustMarshalBinaryLengthPrefixed(whois)
// 	store.Set(key, value)
// }

// GetWhois returns the whois information
// 获取whois结构体数据, key就是域名的name
func (k Keeper) GetWhois(ctx sdk.Context, key string) (types.Whois, error) {
	// 1. 获取namservice数据库
	store := ctx.KVStore(k.storeKey)
	var whois types.Whois
	// 2. 拼接key, whois前缀 + 获取参数key
	byteKey := []byte(types.WhoisPrefix + key)
	// 3. 使用keeper的codec类型cdc解码, 赋值给whois结构体
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &whois)
	if err != nil {
		return whois, err
	}
	// 4. 返回
	return whois, nil
}

// SetWhois sets a whois. We modified this function to use the `name` value as the key instead of msg.ID
// 存储Whois, 我们修改了这个函数，使用' name '值作为键，而不是msg.ID
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois types.Whois) {
	store := ctx.KVStore(k.storeKey)
	// 使用cdc编码参数whois结构体返回的bz是byte切片
	// MustMarshalBinaryLengthPrefixed 有Must代表不返回其错误直接Panic处理， 即使有err的话，没有的话返回可能的错误
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(whois)
	// 使用' name '值作为键，而不是msg.ID
	key := []byte(types.WhoisPrefix + name)
	// 存储
	store.Set(key, bz)
}

// DeleteWhois deletes a whois
// 删除一个whois结构体
func (k Keeper) DeleteWhois(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.WhoisPrefix + key))
}

//
// Functions used by querier  为用户查询的函数
//

// 查询whois的集合
func listWhois(ctx sdk.Context, k Keeper) ([]byte, error) {
	var whoisList []types.Whois
	store := ctx.KVStore(k.storeKey)
	// 根据前缀创建循环迭代器, 遍历所有包含此前缀字段的key对应的value
	// KVStorePrefixIterator 按升序迭代所有带有特定前缀的键
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.WhoisPrefix))
	// 遍历
	for ; iterator.Valid(); iterator.Next() {
		var whois types.Whois
		// 1. 迭代器获取包含特定前缀的完整key
		// 2. 获取whois（[]byte）进行解码
		// 3. 赋值给whois
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &whois)
		// 添加到集合中
		whoisList = append(whoisList, whois)
	}
	// 再将整个list解码/序列化成字节数组，
	// MustMarshalJSONIndent有Must代表不返回其错误直接Panic处理， 即使有err的话，没有的话返回可能的错误
	res := codec.MustMarshalJSONIndent(k.cdc, whoisList)
	return res, nil
}

// 查询单个Whois
func getWhois(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	// 获取key, path的第一个参数(用户命令行输入)
	key := path[0]
	// 调用keeper的基本方法GetWhois(见上方)
	whois, err := k.GetWhois(ctx, key)
	if err != nil {
		return nil, err
	}

	// 编码/序列化为字节数组,
	res, err = codec.MarshalJSONIndent(k.cdc, whois)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Resolves a name, returns the value
// 解析域名对应的值,也就是whois中的字段value
func resolveName(ctx sdk.Context, path []string, keeper Keeper) ([]byte, error) {
	// 直接调用keeper的解析函数(见下方), key是path[0]
	value := keeper.ResolveName(ctx, path[0])

	if value == "" {
		return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not resolve name")
	}

	// 编码/序列化为字节数组
	// QueryResResolve是types/querier文件下的函数, QueryResResolve就是解析值的一个结构体
	// 因为MarshalJSONIndent解析需要一个结构体， 所以创建了这样的QueryResResolve结构体以赋值
	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: value})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
// 获取域名的创建者
func (k Keeper) GetCreator(ctx sdk.Context, key string) sdk.AccAddress {
	whois, _ := k.GetWhois(ctx, key)
	return whois.Creator
}

// Check if the key exists in the store
// 检查当前域名key/name是否存在
func (k Keeper) Exists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.WhoisPrefix + key))
}

// ResolveName - returns the string that the name resolves to
// 获取域名的解析值
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	whois, _ := k.GetWhois(ctx, name)
	return whois.Value
}

// SetName - sets the value string that a name resolves to
// 设置域名对应的解析值
func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
	whois, _ := k.GetWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx, name, whois)
}

// HasOwner - returns whether or not the name already has an owner
// 检查当前域名是否有创建人
func (k Keeper) HasCreator(ctx sdk.Context, name string) bool {
	whois, _ := k.GetWhois(ctx, name)
	return !whois.Creator.Empty()
}

// SetOwner - sets the current owner of a name
// 设置域名的当前拥有者
func (k Keeper) SetCreator(ctx sdk.Context, name string, creator sdk.AccAddress) {
	whois, _ := k.GetWhois(ctx, name)
	whois.Creator = creator
	k.SetWhois(ctx, name, whois)
}

// GetPrice - gets the current price of a name
// 获取域名的价格
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	whois, _ := k.GetWhois(ctx, name)
	return whois.Price
}

// SetPrice - sets the current price of a name
// 设置域名的价格
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whois, _ := k.GetWhois(ctx, name)
	whois.Price = price
	k.SetWhois(ctx, name, whois)
}

// Check if the name is present in the store or not
// 检查当前的name参数是否存在于store中，注意于Exists函数的区别：没有加前缀
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// Get an iterator over all names in which the keys are the names and the values are the whois
// 获取固定前缀的迭代器
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.WhoisPrefix))
}

// Get creator of the item
// 根据key = id获取whois的创建者， 但是可能会返回err
func (k Keeper) GetWhoisOwner(ctx sdk.Context, key string) sdk.AccAddress {
	whois, err := k.GetWhois(ctx, key)
	if err != nil {
		return nil
	}
	return whois.Creator
}

// Check if the key exists in the store
// 根绝key查询是否存在， key是域名结构体中的id
func (k Keeper) WhoisExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.WhoisPrefix + key))
}
