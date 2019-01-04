

# Sequence
> qos, qsc sequence查询

* Uri:

   `/nodes/{chainId}/sequence`

* Method:

    `GET`

* Param

| 参数名   |      类型      |是否必须|默认值|说明|
|----------|:-------------:|:-----:|:---:|:--:|
| chainId | string | yes | - | 链ID|

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

