
# Installation

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

初始化必要的配置文件
./qmoon init

初始化数据库表
./qmoon migration up

启动服务
./qmoon server
```

