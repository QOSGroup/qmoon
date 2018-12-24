
## ATM
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
