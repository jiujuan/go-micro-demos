package main

import (
    "context"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/registry/etcd"
    proto "go-micro-demos/greeter/proto"
    "log"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, resp *proto.HelloResponse) error {
    resp.Greeting = "Hello " + req.Name
    return nil
}

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
    proto.RegisterGreeterHandler(service.Server(), new(Greeter))

    if err := service.Run();err!=nil {
        log.Fatal(err)
    }
}