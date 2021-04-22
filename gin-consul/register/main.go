package main

import (
    "github.com/gin-gonic/gin"
    "github.com/micro/go-plugins/registry/consul/v2"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/web"
    "log"
    //"time"
)

func main() {
    consulReg := consul.NewRegistry(func(op *registry.Options){
        op.Addrs = []string{
            "127.0.0.1:8500", // consul 地址，可以添加多个
           // "192.168.1.1:8500",
        }
    })

    // gin router 初始化路由
    ginRouter := gin.Default()
    ginRouter.GET("/hello", func(ctx *gin.Context) {
        ctx.String(200, "hello consul world!")
    })

    // 注册服务
    service := web.NewService(
        web.Name("hello-service2"), // 微服务的服务名称，把这个服务名称注册到了 consul 里
        web.Address(":8080"),      // http 端口
        web.Handler(ginRouter),       // gin 路由
        web.Registry(consulReg),      // 注册 consul 的地址
        web.Metadata(map[string]string{"data": "the first service test"}), // 携带额外信息到consul
        //web.RegisterTTL(time.Second * 30), // 设置主从服务的过期时间
        //web.RegisterInterval(time.Second * 20), // 设置间隔多久再次注册服务
    )

    // 运行
    if err := service.Run(); err != nil {
        log.Println(err.Error())
    }
}
