
# Validator
> validator基础信息和参与打块
* Uri:

   `/nodes/{chainId}/validators/{address}`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
| chainId | string | yes | - | 链ID|
| address | string | yes | - | 验证者地址|
| minHeight | integer | no | - | 打块查询|
| maxHeight | integer | no | - | 打块查询|



* Response
```
result": {
    "validator": {
      "name": "string",
      "owner": "string",
      "chain_id": "string",
      "address": "string",
      "pub_key_type": "string",
      "pub_key_value": "string",
      "voting_power": 0,
      "accum": 0,
      "first_block_height": 0,
      "first_block_time": "string",
      "status": 0,
      "statusStr": "string",
      "inactiveCode": 0,
      "inactiveTime": "string",
      "inactiveHeight": 0,
      "bondHeight": 0,
      "percent": "string"
    },
    "blocks": [
      {
        "name": "string",
        "owner": "string",
        "chain_id": "string",
        "address": "string",
        "pub_key_type": "string",
        "pub_key_value": "string",
        "voting_power": 0,
        "accum": 0,
        "first_block_height": 0,
        "first_block_time": "string",
        "status": 0,
        "statusStr": "string",
        "inactiveCode": 0,
        "inactiveTime": "string",
        "inactiveHeight": 0,
        "bondHeight": 0,
        "percent": "string"
      }
    ]
  }
```