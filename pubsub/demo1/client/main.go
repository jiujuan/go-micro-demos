package main

import (
    "context"
    "fmt"
    "time"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/util/log"
    "github.com/google/uuid"
    proto "go-micro-demos/pubsub/demo1/proto"
    "github.com/micro/go-micro/v2/metadata"
)

func sendMessage(ctx context.Context, topic string, p micro.Publisher) {
    t := time.NewTicker(time.Second)

    for _ = range t.C {
        msg := &proto.Message{
            Id:      uuid.New().String(),
            Message: fmt.Sprintf("message %s", topic),
        }
        log.Logf("publishing %+v\n", msg)

        // publish msg
        if err := p.Publish(ctx, msg); err != nil {
            log.Logf("publish error: %+v", err)
        }
    }
}

func main() {
    service := micro.NewService(
        micro.Name("go.micro.cli.pubsub"),
    )
    // 解析命令行
    service.Init()

    // created publisher
    topic1 := "example.topic.pubsub.1"
    topic2 := "example.topic.pubsub.2"
    pub1 := micro.NewEvent(topic1, service.Client())
    pub2 := micro.NewEvent(topic2, service.Client())

    ctx := metadata.NewContext(context.Background(), map[string]string{
        "id": "1",
        "content": "this is metadata test",
    })

    go sendMessage(ctx, topic1, pub1)
    go sendMessage(ctx, topic2, pub2)

    select{}
}

