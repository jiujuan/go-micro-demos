package main

import (
    "fmt"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/broker"
    "github.com/micro/go-plugins/broker/rabbitmq/v2"
    "go-micro-demos/broker/rabbitmq/config"
    "go-micro-demos/broker/rabbitmq/subscriber"
    "log"
)

var (
    conf = config.Config()
    topic = conf.Topics["rabbitmq"].Name
    amqp = conf.Rabbitmq.Amqp
    serviceName = conf.Services["rabbitmq"].Name
)

func main() {
    fmt.Println("topic: ", topic, "serviceName: ", serviceName)

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
        fmt.Printf("failed to connect broker: %v", err)
        return
    }

    subscriber.Sub(topic, brk)

    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
