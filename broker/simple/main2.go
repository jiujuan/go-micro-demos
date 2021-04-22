package main

import (
    "fmt"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/broker"
    "time"
    "log"
)

var (
    topic2 = "go.micro.topic.foo"
)

func main() {
    // new service
    service := micro.NewService(
        micro.Name("com.brk.demo"),
    )
    service.Init(micro.AfterStart(func() error {
        brk := service.Options().Broker

        go sub2(brk)
        go pub2(brk)
        return nil
    }))

    service.Run()
}

func pub2(brk broker.Broker) {
    tick := time.NewTicker(time.Second)
    i := 0

    for _ = range tick.C {
        // 构建一个消息
        msg := &broker.Message{
            Header: map[string]string{
                "id": fmt.Sprintf("%d", i),
            },
            Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
        }

        // 发布构建的消息到某个 topic
        if err := brk.Publish(topic2, msg); err != nil {
            log.Printf("[pub] failed: %v", err)
        } else {
            fmt.Println("[pub] published message:", string(msg.Body))
        }
        i++
    }
}

func sub2(brk broker.Broker) {
    // 订阅消息，通过指定的 topic，一般是客户端订阅，这里写到一个测试程序里
    _, err := brk.Subscribe(topic2, func(p broker.Event) error {
        fmt.Println("[sub] received message: ", string(p.Message().Body), "header", p.Message().Header)
        return nil
    }, broker.Queue(topic2))
    if err != nil {
        fmt.Println(err)
    }
}
