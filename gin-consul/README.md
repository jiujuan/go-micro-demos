# demo04: gin 结合 consul 服务注册和服务发现
> consul 软件请自行安装

## 服务注册

1. 先编写注册程序 register/main.go，具体代码见 github
2. 编写完后运行 go run main.go
3. 在浏览器上执行 http://localhost:8080/hello

然后打开 consul 的 web 界面 http://localhost:8500/ui/#/dc1/services，

就可以看到 hello-service2 成功注册了。

>consul 有一个 check service 的服务，你过一段时间再去查看 hello-service2，会发现它被删除了。

## 服务发现

```go
// 获取注册的服务
// 返回的 services 是一个 slice
services, err := consulReg.GetService("hello-service2")

// 随机获取一个服务
randomsvc := selector.Random(services)
svc, err := randomsvc()
```

具体代码见 github 。

编写完成后运行程序,go run main.go , 显示信息如下：
>get service:  704d5505-b3ac-4b71-8cdb-d374382f1fef 192.168.0.100:8080 map[data:the first service test]

id, address, metadata 信息获得了。