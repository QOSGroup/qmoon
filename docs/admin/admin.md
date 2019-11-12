# 管理后台相关API

## 注册
```
api: /admin/register
method: POST

post body
    mail 邮箱
    password 密码
```
## 登陆
```
api: /admin/login
method: POST

post body
    mail 邮箱
    password 密码
```
## 登出
```
api: /admin/logout
method: POST
```
## 查看app
```
api: /apps
method: GET
```
## 创建app
```
api: /apps
method: POST

post body
    name app名称
```
## 删除app
```
api: /apps/{appid}
method: DELETE
```

## 更新个人信息
```
api: /admin/register
method: GET

post body
    mail 邮箱
    password 密码
```
