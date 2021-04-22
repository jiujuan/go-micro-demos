package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro/v2"
    hello "go-micro-demos/hello/proto/hello"
)

func main() {
    // new service
    service := micro.NewService(
        micro.Name("go.micro.client.hello"),
    )

    // init service
    service.Init()

    // create hello service client
    helloClient := hello.NewHelloService("go.micro.service.hello", service.Client())

    resp, err := helloClient.Call(context.TODO(), &hello.Request{Name: " world!!"})
    if err!=nil{
        fmt.Println(err)
        return
    }

    fmt.Println(resp.Msg)
}

