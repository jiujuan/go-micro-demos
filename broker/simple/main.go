package main

import (
    "fmt"
    "github.com/micro/go-micro/v2/broker"
    "log"
    "time"
)

var (
    topic = "go.micro.topic.foo"
)

func main() {

    if err := broker.Init(); err != nil {
        log.Fatalf("Broker Init error: %v", err)
    }
    if err := broker.Connect(); err != nil {
        log.Fatalf("Broker Connect error: %v", err)
    }

    go sub()
    go pub()

    <-time.After(time.Second * 10)
}

func pub() {
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
        if err := broker.Publish(topic, msg); err != nil {
            log.Printf("[pub] failed: %v", err)
        } else {
            fmt.Println("[pub] published message:", string(msg.Body))
        }
        i++
    }
}

func sub() {
    // 订阅消息，通过指定的 topic，一般是客户端订阅，这里写到一个测试程序里
    _, err := broker.Subscribe(topic, func(p broker.Event) error {
        fmt.Println("[sub] received message: ", string(p.Message().Body), "header", p.Message().Header)
        return nil
    })
    if err != nil {
        fmt.Println(err)
    }
}