package types

const (
	// ModuleName is the name of the module
	// 模块名
	ModuleName = "cosmosnameservice"

	// StoreKey to be used when creating the KVStore
	// 模块存储空间的键/key
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	// 路由消息的键名
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	// 查询消息的查询路由名
	QuerierRoute = ModuleName
)

const (
	// whois结构体的前缀，即在模块空间中k/v的前缀key
	WhoisPrefix = "whois-value-"
	// 存储总记数的前缀key，是WhoisPrefix的特例
	WhoisCountPrefix = "whois-count-"
)
