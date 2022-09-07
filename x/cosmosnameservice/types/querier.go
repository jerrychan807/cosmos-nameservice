package types

import "strings"

// 查询常量，对应客户端输入的参数
const QueryListWhois = "list-whois"
const QueryGetWhois = "get-whois"
const QueryResolveName = "resolve-name"

// QueryResResolve Queries Result Payload for a resolve query
// 查询域名解析结果结构体函数
// 因为MarshalJSONIndent解析需要一个结构体， 所以创建了这样的QueryResResolve结构体以赋值
type QueryResResolve struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
// 重写string方法
func (r QueryResResolve) String() string {
	return r.Value
}

// QueryResNames Queries Result Payload for a names query
// 查询域名群集合的解析结果，返回结果的切片
type QueryResNames []string

// implement fmt.Stringer
// 格式化输出解析结果切片的数据
func (n QueryResNames) String() string {
	return strings.Join(n[:], "\n")
}
