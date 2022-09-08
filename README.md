# cosmosnameservice

**cosmosnameservice** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

Blog:[StudyRecord-Cosmos官方NameService域名解析应用教程](https://jerrychan807.github.io/2022/37691.html)

# Snapshot:

```shell
# 查询 namechain.com 的信息
cosmos-nameservicecli query cosmosnameservice resolve namechain.com

[root@localhost cosmos-nameservice]# cosmos-nameservicecli query cosmosnameservice resolve namechain.com
{
  "value": "8.8.8.8"
}
```

# Refs:

- [cosmos官方nameservice测试项目详解（代码注释+官方文档错误纠正）](https://blog.csdn.net/weixin_43988498/article/details/114996094)
- [Cosmos SDK 中文文档](https://learnblockchain.cn/docs/cosmos/tutorial/06-set-name.html#msg)
- [cosmos-sdk-tutorials namechain 体验过程解析](https://www.jianshu.com/p/547fc70b8335)
- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)