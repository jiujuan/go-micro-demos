package main

import (
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
    "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
    "go-micro-demos/tracer/tracer"
    "log"
)

const (
    ServerName = "go.micro.service.task"
    JaegerAddress = "127.0.0.1:6831"
)

func main() {
    jaegerTracer, closer, err := tracer.NewJaeger(ServerName, JaegerAddress)
    if err != nil {
        log.Fatal(err)
    }
    defer closer.Close()

    app := micro.NewService(
        micro.Name("go.micro.client.greeter"),

        // ...

        micro.WrapClient(
            // 配置断路器 hystrix
            hystrix.NewClientWrapper(),
            // 配置 jaeger
            opentracing.NewClientWrapper(jaegerTracer),
         ),
    )

    app.Init()

    app.Run()
}