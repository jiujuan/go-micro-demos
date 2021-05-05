package main

import (
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
    "go-micro-demos/tracer/tracer"
    "log"
)

const (
    ServerName = "go.micro.service.task"
    JaegerAddress = "127.0.0.1:6831"
)

func main() {

    // 配置 jaeger 连接
    jaegerTracer, closer, err := tracer.NewJaeger(ServerName, JaegerAddress)
    if err != nil {
        log.Fatal(err)
    }
    defer closer.Close()

    // New Service
    service := micro.NewService(
        micro.Name(ServerName),
        //...
        // 配置 jaeger
        micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),

     )

    service.Init()

    service.Run()
}
