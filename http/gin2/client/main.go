package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro/v2/client"
    "github.com/micro/go-micro/v2/client/selector"
    "github.com/micro/go-micro/v2/registry"
    microhttp "github.com/micro/go-plugins/client/http/v2"
    "github.com/micro/go-plugins/registry/consul/v2"
    "log"
)

func main() {
    consulReg := consul.NewRegistry(
        registry.Addrs("127.0.0.1:8500"), // 如果只有一个地址，可以这样简单写，多个地址可以像server端那样写
     )

    mySelector := selector.NewSelector(
        // 服务注册中心地址
        selector.Registry(consulReg),
        // 选择轮询算法
        selector.SetStrategy(selector.RoundRobin),
     )

    //=================
    // 获取 house-service 服务
    services, err := consulReg.GetService("house-service")
    if err != nil {
        log.Fatal("failed to get service list")
    }
    //randomsrv := selector.Random(services) // 随机选择
    randomsrv := selector.RoundRobin(services) // 轮询
    srv, err := randomsrv()
    if err != nil {
        log.Fatal("random service error: " , err)
    }
    fmt.Println(srv.Id, srv.Metadata, srv.Address)
    //===========上面这段代码测试获取 house-service 服务，与下面代码没有关联=========

    // 下面利用 go-micro 的 plugins 里的 client http 插件随机选择服务，然后发起请求
    // 创建一个客户端，可以到注册中心选择服务
    newClient := microhttp.NewClient(
        client.Selector(mySelector),
        client.ContentType("application/json"),
    )

    // 发起 http 请求
    req := newClient.NewRequest(
        "house-service",
        "/v1/house",
        map[string]interface{}{"num":5},
    )
    var resp map[string]interface{}
    if err:=newClient.Call(context.Background(), req, &resp); err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp)
}
