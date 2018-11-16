# Qmoon

[![version](https://img.shields.io/github/tag/QOSGroup/qmoon.svg)](https://github.com/QOSGroup/qmoon/releases/latest)
[![Build Status](https://travis-ci.org/QOSGroup/qmoon.svg?branch=master)](https://travis-ci.org/QOSGroup/qmoon)
[![codecov](https://codecov.io/gh/QOSGroup/qmoon/branch/master/graph/badge.svg)](https://codecov.io/gh/QOSGroup/qmoon)
[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://godoc.org/github.com/QOSGroup/qmoon)
[![Go version](https://img.shields.io/badge/go-1.11.0-blue.svg)](https://github.com/moovweb/gvm)
[![license](https://img.shields.io/github/license/QOSGroup/qmoon.svg)](https://github.com/QOSGroup/qmoon/blob/master/LICENSE)
[![](https://tokei.rs/b1/github/QOSGroup/qmoon?category=lines)](https://github.com/QOSGroup/qmoon)


QOS公链,QSC联盟链数据查询服务

## 安装

* 下载和安装
```
go get -u github.com/QOSGroup/qmoon
cd $GOPATH/github.com/QOSGroup/qmoon/cmd
go build -o qmoon
```

* 启动服务
需要安装postgresql
```
create database qmoon ENCODING 'UTF8' TEMPLATE template0;

# 初始化必要的配置文件
./qmoon init

# 初始化数据库表
./qmoon migration up

# 启动服务
./qmoon server
```



## TODO

> - [x] 项目初始化，init和migration命令
> - [x] tendermint默认提供的API
> - [x] 账户及app相关
> - [x] API访问控制
> - [ ] 数据被动同步
> - [x] 数据主动同步
> - [x] qos数据查询API
> - [x] qsc数据查询API
> - [ ] 动态更新配置
