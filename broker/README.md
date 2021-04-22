# demo05: broker 异步消息，消息中间件

**异步消息处理**：
把消息暂时存储到某一个位置，比如 kafka，nats，redis，rabbitmq，rocketmq 等，
然后等待其他客户端消费。

go-micro 的 pub/sub 发布订阅功能是在 broker 接口基础上构建的。

```go
broker.Publish
broker.Subscribe
```

## 例子一：simple/main.go

### 1.sub() 函数，订阅

```go
broker.Subscribe(topic, func(p broker.Event) error{}) 
```

上面这个 Subscribe() 函数在 `micro/go-micro/broker/broker.go`

```go
func Subscribe(topic string, handler Handler, opts ...SubscribeOption) (Subscriber, error) {
	return DefaultBroker.Subscribe(topic, handler, opts...)
}
```

第一个参数：订阅的主题 ，topic。

第二个参数：事件处理 Handler，它是一个回调函数，收到新事件时，此函数会执行。

第三个参数：option，订阅选项。这里可以实现自己的选项功能，从而扩展插件功能。

### 2.pub() 函数，发布

里面的 Publish() 函数也是在 `micro/go-micro/broker/broker.go` 文件中：

```go
func Publish(topic string, msg *Message, opts ...PublishOption) error {
	return DefaultBroker.Publish(topic, msg, opts...)
}
```

第一个参数：主题 topic。

第二个参数：消息 message。这个消息结构定义在 `micro/go-micro/broker/broker.go` 文件中，包含消息头和消息体：

```go
type Message struct {
	Header map[string]string
	Body   []byte
}
```

第三个参数：发布选项，可选参数



追踪 `broker.Init()` 会发现这个例子用的 HTTP Broker，那它数据存在哪里呢？总要有一个暂存数据的地方。

到源码里去看看 `micro/go-mciro/broker/http.go` ，有一个 `httpBroker struct` 的数据结构：

```go
// HTTP Broker is a point to point async broker
type httpBroker struct {
	id      string
	address string
	opts    Options

	mux *http.ServeMux

	c *http.Client
	r registry.Registry

	sync.RWMutex
	subscribers map[string][]*httpSubscriber // 订阅
	running     bool
	exit        chan chan error

	// offline message inbox
	mtx   sync.RWMutex
	inbox map[string][][]byte
}
```

存数据的地方已经做了说明，就是 inbox 字段。

### 3.代码例子

main.go

main2.go

main3.go



这 3 个例子大同小异，主要都是参考官方的例子，然后根据自己需要做一点变通。

但是 main2.go 运行后，出现一个 missing service port 的错误，一直 publish message，
而 sub 没有订阅到消息。暂时没时间研究它，TODO 。- -!

## 例子二：rabbitmq 发布/订阅

### rabbitmq 安装

用 docker 直接安装。

查询 rabbitmq 安装版本，2 种方法：

```shell
// 第一种方法, 命令查询
docker search rabbitmq

// 第二种方法，浏览器查看
https://hub.docker.com/_/rabbitmq
```

在浏览器上查看 rabbitmq 的 tag 信息：

> Supported tags and respective Dockerfile links
> 3.8.14, 3.8, 3, latest
> 3.8.14-management, 3.8-management, 3-management, management
> 3.8.14-alpine, 3.8-alpine, 3-alpine, alpine
> 3.8.14-management-alpine, 3.8-management-alpine, 3-management-alpine, management-alpine


我们安装 3.8-management 这个 tag 的版本，命令：

```shel
docker pull docker.io/rabbitmq:3.8-management
```

安装完成后查看下载的 image： `docker ps`

然后根据 image id 来启动 rabbitmq：

```shell
docker run --name rabbitmq -d -p 15672:15672 -p 5672:5672 edd581f9
aaa6b70f9f542b1986b7b2af1a2adf7f

// 查看启动的 rqbbitmq 是否启动
docker ps 
```

我这里已经启动，在浏览器上查看 web 界面，http://192.168.1.109:15672/#/ ，出来 web 界面启动成功。

### 代码例子

1. 配置文件

用 toml 格式存储配置文件，rabbitmq/config/config.toml。

toml 配置文件：config.toml。

读取配置文件程序：config.go。

具体代码见 github 。

2. 数据格式文件

定义传输的数据 proto/message.proto：

```protobuf
syntax = "proto3";

message Message {
    string id = 1;
    string message = 2;
}
```

然后用命令解析：`protoc --go_out=. --micro_out=. ./message.proto` ，生成对应的 go 文件

3. 订阅程序

先编写 subscriber/subscriber.go， 具体代码见 github。

main.go  函数调用订阅程序

4. 发布程序

publisher/publisher.go

这个程序里函数 `publisher(topic string, brk broker.Broker){}`

```go
msg := &broker.Message{
    Header: map[string]string{
        "id": strconv.Itoa(uuid.ClockSequence()),
    },
    Body: msgBody,
}
```

这里的 broker.Message{} 和 broker.Event 啥关系？

可以到 micro/go-micro/broker/broker.go 看看，它们的结构：

```go
type Message struct {
	Header map[string]string
	Body   []byte
}

// Event is given to a subscription handler for processing
type Event interface {
	Topic() string
	Message() *Message
	Ack() error
	Error() error
}
```

`Event` 是一个 interface，里面的 Message() *Message 方法返回值是 broker.Message 结构体




5. 测试运行

先确定前面安装的 rabbitmq 运行正常。

在运行 `go run main.go`， 启动订阅端。

最后运行 `go run publisher.go` ，运行发布端。



如果运行正常，就可以看到发布的消息被订阅端消费了。

每次运行 `go run publisher.go`，订阅端都会显示消费发布信息。














