
# Account Tx
> 查询账户交易
> 最大返回7天的数据，如果 endDate-startDate > 7天, 则返回 endDate - endDate-6天的数据。\
> 如果根据块查询，则默认查询20块的数据

* Uri:

   `/nodes/{chainId}/accounts/{address}/txs`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
| chainId | string | yes | - | 链ID|
| address | string | yes | - |地址|
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
