package main

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
    "github.com/micro/go-micro/v2"
    "strconv"

    "github.com/micro/go-micro/v2/broker"
    "github.com/micro/go-plugins/broker/rabbitmq/v2"
    "go-micro-demos/broker/rabbitmq/config"
    proto "go-micro-demos/broker/rabbitmq/proto"
    "time"
)

var (
    conf = config.Config()
    topic = conf.Topics["rabbitmq"].Name
    amqp = conf.Rabbitmq.Amqp
    serviceName = conf.Services["rabbitmq"].Name
)

func main() {
    fmt.Println("topic: ", topic, ", serviceName: ", serviceName)

    // new service
    service := micro.NewService(
        micro.Name(serviceName),
        micro.Version("latest"),
        micro.Broker(
            rabbitmq.NewBroker(broker.Addrs(amqp)),
        ),
    )
    service.Init()

    brk := service.Server().Options().Broker
   if err := brk.Connect();err != nil {
       fmt.Printf("[pub]failed to connect broker: %v", err)
       return
   }

    go publisher(topic, brk)

    <-time.After(time.Second * 20)
}

// 发送消息
func publisher(topic string, brk broker.Broker) {
    t := time.NewTicker(time.Second)

    var i = 0
    for _ = range t.C {
        protoMsg := &proto.Message{
            Id     : uuid.New().String(),
            Message: fmt.Sprintf("topic: %s, %d:%s", topic, i, time.Now().String()),
        }
        msgBody, err := json.Marshal(protoMsg)
        if err != nil {
            fmt.Println(err)
        }

        msg := &broker.Message{
            Header: map[string]string{
                "id": strconv.Itoa(uuid.ClockSequence()),
            },
            Body: msgBody,
        }
        fmt.Println(string(msgBody))

        if err := brk.Publish(topic, msg); err != nil {
            fmt.Printf("publish error: %v", err)
        }
        i++
    }

    //defer brk.Disconnect()
}
