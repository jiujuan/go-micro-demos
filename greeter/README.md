#demo01: greeter

> 说明：先自行安装 etcd 软件
>
> 这里有一个 etcd 安装说明：[etcd安装说明地址](https://www.bookstack.cn/read/docker_practice-1.2.0/etcd-install.md)

写一个 go-micro 的例子，结合 etcd 写一个简单的示例

1.编写 proto/greeter.proto，然后用 protoc 编译：
```shell script
protoc --go_out=. --micro_out=. ./greeter.proto
```

2.编写客户端代码
client/main.go

```go
//初始化一个服务
service := micro.NewService() 

// cmd 参数初始化
service.Init() 

// 初始化 greeter 服务
greeter := proto.NewGreeterService()
// 调用服务
resp, err := greeter.Hello() 

```

3.编写服务端代码
server/main.go

