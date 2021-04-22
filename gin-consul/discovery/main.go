package main

import (
    "fmt"
    "github.com/micro/go-plugins/registry/consul/v2"
    "github.com/micro/go-micro/v2/registry"
    "log"
    "github.com/micro/go-micro/v2/client/selector"
)

func main() {
    consulReg := consul.NewRegistry(func(op *registry.Options){
        op.Addrs = []string{
            "127.0.0.1:8500", // consul 地址，可以添加多个
            // "192.168.1.1:8500",
        }
    })

    // 根据微服务名获取微服务列表
    services, err := consulReg.GetService("hello-service2")
    if err != nil {
        log.Fatal("failed to get service list")
    }

    // 随机获取一个服务
    //randomsvc := selector.Random(services)   // 随机选择
    randomsvc := selector.RoundRobin(services) // 轮询
    svc, err := randomsvc()
    if err != nil {
        log.Fatal("failed to get service: ", err.Error())
    }

    fmt.Println("get service: ", svc.Id, svc.Address, svc.Metadata)
}
