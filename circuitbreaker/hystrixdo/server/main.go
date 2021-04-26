package main

import(
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-plugins/registry/consul/v2"
    proto "go-micro-demos/circuitbreaker/hystrixdo/proto/house"
    hservice "go-micro-demos/circuitbreaker/hystrixdo/server/service"
)

func main() {
    consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

    service := micro.NewService(
        micro.Name("go.micro.service.house"),
        micro.Registry(consulReg),
    )

    service.Init()

    err := proto.RegisterHouseHandler(service.Server(), &hservice.HouseService{})
    if err != nil {
        panic(err)
    }

    service.Run()
}
