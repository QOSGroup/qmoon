# Account Query
> 查询账户余额

* Uri:

   `/nodes/{chainId}/accounts/{address}`

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

