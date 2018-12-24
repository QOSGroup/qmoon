# qmoon restful API

## 错误码说明
http 状态码一直为200
业务是否正常根据返回结果json中error对象,具体错误原因详见错误吗

| 状态码   |      说明      |
|----------|:-------------:|
|  |  |
|  |  |
|  |  |


## 响应体格式
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


## 接口说明
> 以下接口描述中，Response只描述 result 中数据结构 \
> 同时提供原始tendermint的接口,需要增加前缀/nodes/{chainId}/tendermint,比如需要访问QOS的genesis,则只需要访问 /nodes/QOS/genesis

### 块列表
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


### 块详情

* Uri:

   `/nodes/{chainId}/blocks/{height}`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|


* Response
```
"result":{
    "chain_id": "qstars-test",
    "height":100,
    "block_id": "B17D877473E9FE91628BCF0885624A922D677D24",
    "time": "2018-11-06T05:59:15.420249826Z",
    "num_txs": "2",
    "total_txs": "2",
    "data_hash": "2AD717CAED295643DDA6A14412BB28889D3C3860",
    "txs":["tx1","tx2"]
    "validators_hash":"BDB1BB8971FF0F97FA9AE00175CC37F897359865",
    "validators":["C95E740DBD2285BB2060630FDE24FCD971ADE4C7"],
}
```

### 交易列表

* Uri:

   `/nodes/{chainId}/txs`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
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


### 块详情

* Uri:

   `/nodes/{chainId}/txs/{id}`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|


* Response
```
"result":{
    "chain_id": "qstars-test",
    "height":100,
    "block_id": "B17D877473E9FE91628BCF0885624A922D677D24",
    "time": "2018-11-06T05:59:15.420249826Z",
    "num_txs": "2",
    "total_txs": "2",
    "data_hash": "2AD717CAED295643DDA6A14412BB28889D3C3860",
    "txs":["tx1","tx2"]
    "validators_hash":"BDB1BB8971FF0F97FA9AE00175CC37F897359865",
    "validators":["C95E740DBD2285BB2060630FDE24FCD971ADE4C7"],
}
```


### sequence查询

* Uri:

   `/nodes/{chainId}/peer`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|


* Response
```
"result":{
    "apps":[
    {
        "name":"qos-test",
         "in":"100",
         "out":"100",
    }
    ]

}
```


### validators

* Uri:

   `/nodes/{chainId}/validators`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
| height | number | false | 查询最新块的validators |块高度|


* Response
```
"result":{
    "validators": [
        {
            "address": "C95E740DBD2285BB2060630FDE24FCD971ADE4C7",
            "pub_key": {
                "type": "tendermint/PubKeyEd25519",
                "value": "XsVLCKdI6FMhQex5gPzZqunXli8fREkZrOgRj6Lipdg="
            },
            "voting_power": "10",
            "accum": "0",
        }
    ]
}
```

### validators detail(暂未提供)
> 查询validator参与打块信息
* Uri:

   `/nodes/{chainId}/validators/{addr}`

* Method:

    `GET`

* Param



* Response
```
"result":{

}
```


### 查询账户余额

* Uri:

   `/nodes/{chainId}/accounts/{address}`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|


* Response
```
"result": {
    {
      "type": "qbase/account/QOSAccount",
      "value": {
        "base_account": {
          "account_address": "address1k0m8ucnqug974maa6g36zw7g2wvfd4sug6uxay",
          "public_key": {
            "type": "tendermint/PubKeyEd25519",
            "value": "E/LYBhSS6fgUlRC1tn00DRmf8k80yF27vX4N94Dppsw="
          },
          "nonce": "33096"
        },
        "qos": "99999982",
        "qscs": [
          {
            "coin_name": "qstar",
            "amount": "99966922"
          }
        ]
      }
    }
}
```


### peers查询

* Uri:

   `/nodes/{chainId}/peers`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|


* Response
```
"result":{
    "n_peers": "0",
    "peers": [
    {
        "moniker": "qos-test1_node1",
        "id": "9283a73b40704e8aefaa82a40021ed31f6e29306",
        "listen_addr": "192.168.1.228:26656",
        "network": "qos-test1",
        "version": "0.23.1",
        "channels": "4020212223303800",
        "send_start": "2018-11-07T08:09:46.94Z",
        "recv_start": "2018-11-07T08:09:46.94Z",
        "create_at": "2018-11-09T08:47:32.029584Z"
    }
    ]
}
```


### 查询账户交易
> 最大返回7天的数据，如果 endDate-startDate > 7天, 则返回 endDate - endDate-6天的数据。\
> 如果根据块查询，则默认查询20块的数据

* Uri:

   `/nodes/{chainId}/accounts/{addr}/txs`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
| startDate | string | false | 最近24小时的交易 |开始时间|
| endDate | string | false | 最近24小时的交易 |结束时间|
| minHeight | number | false | |最小高度|
| maxHeight |number  |false||最大高度|

* Response
```
"result": {
    "txs":[]
}
```


### 查询账户转账记录
> 查询账户转账记录

* Uri:

   `/nodes/{chainId}/accounts/:address/transfer`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
| address | string | yes | - |地址|
| offset | number | false | 0|开始位置|
| limit |number  |false|20|一页大小|

* Response
```
"result": {
    "txs":[
    {
        ID      int64        `json:"id"`       // id
        ChainID string       `json:"chain_id"` // chain_id
        Height  int64        `json:"height"`   // height
        Hash    string       `json:"hash"`     // hash
        Address string       `json:"address"`  // address
        Coin    string       `json:"coin"`     // coin
        Amount  string       `json:"amount"`   // amount
        Type    TransferType `json:"type"`     // type
        Time    time.Time    `json:"time"`     // time
     }
    ]
}
```


### 领钱
> 领钱

* Uri:

   `/nodes/{chainId}/accounts/:address/withdraw`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
| address | string | yes | - |地址|

* Response
```
"result": {
    hash   string
    heigth string
}
```
