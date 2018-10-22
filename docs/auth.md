
# 签名算法
## 签名说明
> 公链和联盟数据查询的接口，qmoon基于密钥 HMAC (Hash Message Authentication Code) 的自定义 HTTP 方案进行身份验证。
用户可以去qos浏览器申请SecretID和SecretKey

## 签名生成

```
// 拼接签名串
orignal = 'k=[SecretID]&e=[expiredTime]&t=[currentTime]&r=[rand]'

// 加密拼接的签名串
signTmp = HMAC-SHA1(SecretKey, orignal)

// 加密后的加密串Base64编码
sign = Base64(SecretID + orignal)

// curl -H 'Authorization: ' 'https://api.qmoon.com/xxx'
```

字段说明:

| 字段 |      解释      |
|:----:|:-------------|
| k    | Secret ID |
| e    | 签名的有效期，是一个符合 UNIX Epoch 时间戳规范的数值，单位为秒|
| t    | 当前时间戳，是一个符合 UNIX Epoch 时间戳规范的数值，单位为秒 |
| r    | 随机串，无符号10进制整数，用户需自行生成，最长10 位 |

qmoon使用 HMAC-SHA1 算法对请求进行加密（SHA1 算法加密后的输出必须是原始的二进制数据，否则签名失败）；
签名串需要使用 Base64 编码
首先对 orignal 使用 HMAC-SHA1 算法进行签名，然后将 orignal 附加到签名结果的末尾，再进行 Base64 编码，得到最终的 sign。