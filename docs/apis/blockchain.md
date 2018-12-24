
# Blockchain
> 一次最多返回20个块,如果 maxHeight-minHeight > 20, 则返回 maxHeight-maxHeight-19之间的块

* Uri:

   `/nodes/{chainId}/blockchain`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
| minHeight | number | false | 默认返回当前最新20块|最小高度|
| maxHeight |number  |false|默认返回当前最新20块|最大高度|


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

