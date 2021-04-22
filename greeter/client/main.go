package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/registry/etcd"
    proto "go-micro-demos/greeter/proto"
)

func main() {
    reg := etcd.NewRegistry(func(op *registry.Options){
        op.Addrs = []string{
            "127.0.0.1:2379",
        }
    })

    service := micro.NewService(
        micro.Name("greeter"),
        micro.Registry(reg),
        )

    service.Init()

    greeter := proto.NewGreeterService("greeter", service.Client())

    resp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name:"Micro World"})
    if err != nil {
        panic(err)
    }

    fmt.Println(resp.GetGreeting())
}
