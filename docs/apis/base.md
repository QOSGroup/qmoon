[[toc]]

# API Instructions

## Authorization
所有请求需要在header中增加，QToken:xxx

token在[explorer](http://explorer.qoschain.info/developer)中申请



## Response Code
http 状态码为200
业务是否正常根据返回结果json中error对象,具体错误原因详见错误吗

| 状态码   |      说明      |
|----------|:-------------:|
|  |  |
|  |  |
|  |  |


## Response Struct
> 后续接口描述中，Response只描述 result 中数据结构 \
> 同时提供原始tendermint的接口,需要增加前缀/nodes/{chainId}/tendermint,比如需要访问QOS的genesis,则只需要访问 /nodes/QOS/genesis



```
{
    "jsonrpc": "2.0",
    "id": "",
    "result": {},
    "error": {
        "code": 000,
        "message": "错误信息",
        "data": "错误详细描述"
    }
}
```



