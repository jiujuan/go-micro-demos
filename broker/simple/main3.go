package main

import (
    "fmt"
    "github.com/micro/go-micro/v2/broker"
    "log"
    "time"
)

func main() {
    if err := broker.Connect(); err != nil {
        fmt.Println("Broker Connect error: %v", err)
    }

    go pub3("go.micro.topic.foo")
    go sub3("go.micro.topic.foo")

    time.Sleep(time.Second * 10)
}

func pub3(topic string) {
    tick := time.NewTicker(time.Second)
    i := 0
    for _ = range tick.C {
        msg := &broker.Message{
            Header: map[string]string{
                "id": fmt.Sprintf("%d", i),
            },
            Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
        }
        if err := broker.Publish(topic, msg); err != nil {
            log.Printf("[pub] failed: %v", err)
        } else {
            fmt.Println("[pub] pubbed message:", string(msg.Body))
        }
        i++
    }
}

func sub3(topic string) {
    _, err := broker.Subscribe(topic, func(p broker.Event) error {
        fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
        return nil
    })
    if err != nil {
        fmt.Println(err)
    }
}


