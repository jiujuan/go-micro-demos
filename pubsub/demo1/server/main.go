package main

import (
    "context"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/metadata"
    "github.com/micro/go-micro/v2/server"
    "github.com/micro/go-micro/v2/util/log"
    proto "go-micro-demos/pubsub/demo1/proto"
    "net/http"
)

type Sub struct{}

func (s *Sub) Process(ctx context.Context, msg *proto.Message, req *http.Request) error {
    md, _ := metadata.FromContext(ctx)
    log.Logf("[pubsub.1] received msg %+v with metadata %+v", msg, md)

    return nil
}

// a function can be used
func subMsg(ctx context.Context, msg *proto.Message) error {
    md, _ := metadata.FromContext(ctx)
    log.Logf("[pubsub.2] received msg %+v with metadata %+v", msg, md)

    return nil
}

func main() {
    service := micro.NewService(
        micro.Name("go.micro.srv.pubsub"),
    )

    service.Init()

    // 注册订阅
    micro.RegisterSubscriber("example.topic.pubsub.1", service.Server(), new(Sub))
    // 把订阅注册到某个queue里，每一个分发的消息都会进入唯一的subscriber的queue
    micro.RegisterSubscriber("example.topic.pubsub.2", service.Server(), subMsg, server.SubscriberQueue("queue.pubsub"))

    if err := service.Run(); err != nil{
        log.Fatal(err)
    }
}
