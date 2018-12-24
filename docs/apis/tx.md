
# Tx
> 交易详情

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
