package subscriber

import (
    "encoding/json"
    "fmt"
    "github.com/micro/go-micro/v2/broker"
    proto "go-micro-demos/broker/rabbitmq/proto"
)

// 订阅消息
func Sub(topic string, brk broker.Broker) {
    _, err := brk.Subscribe(topic, func(event broker.Event) error {
        var msg *proto.Message
        if err := json.Unmarshal(event.Message().Body, &msg); err != nil {
            return err
        }

        head := event.Message().Header
        fmt.Println(head["id"], "received message: ", msg.Message, "id: ", msg.Id)
        return nil
    })
    if err != nil {
        fmt.Println("receive message error")
    }
}
