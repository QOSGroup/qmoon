
# Txs
> 交易列表

* Uri:

   `/nodes/{chainId}/txs`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
| chainId | string | yes | - | 链ID|
| minHeight | number | false | 默认返回最新20条记录|最小高度|
| maxHeight |number  |false|默认返回最新20条记录|最大高度|


* Response
```
"result":{
    "last_height": "170156",
    "blocks":[
        "chain_id": "qstars-test",
        "height":100,
        "block_id": "B17D877473E9FE91628BCF0885624A922D677D24",
        "time": "2018-11-06T05:59:15.420249826Z",
        "num_txs": "0",
        "total_txs": "2",
    ]
}
```
