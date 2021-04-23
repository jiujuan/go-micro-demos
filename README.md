# go-micro 微服务框架练习的例子

>go v1.13
>
>go-micro v2.9.1

golang 微服务框架 go-micro 一步一步学习的小例子。

## 目录

- **01：greeter** 手写一个例子

  手动开始写一个简单 go-micro 例子，结合 etcd。

- **02：hello** 自动生成代码

  用 micro 自带命令生成服务例子。

- **03：http-gin** http api 网关

  go-micro 结合 http，gin 的例子。

  gin2，gin 提供 http 服务，相当于 http 网关，后面在调用一些服务。

- **04：gin-consul** 服务注册与发现

  gin 结合 consul 实现服务注册和服务发现的小例子。也可以换成 etcd 代码是一样的。

  负载均衡，可以随机、轮询选择一个服务。

- **05：broker** 消息中间件
  
  异步消息处理，结合一些消息队列（消息中间件）来处理消息。比如 kafka，redis，rabbitmq，nats 等消息中间件。
  
- **06：pubsub** 发布订阅程序

- **07：wrapper** 包装功能

  一个扩展功能，如果你想对程序进一步扩展，可以使用这个功能

- **08：circuitbreaker** 熔断功能

  