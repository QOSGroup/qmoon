
# Peers

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
