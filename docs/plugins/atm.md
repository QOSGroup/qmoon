
# ATM
> 领钱,一个地址一天只能领取一次，一次领取10万qos

* Uri:

   `/nodes/{chainId}/accounts/{address}/withdraw`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
| chainId | string | yes | - | 链ID|
| address | string | yes | - |地址|

* Response
```
"result": {
    hash   string
    height string
}
```
