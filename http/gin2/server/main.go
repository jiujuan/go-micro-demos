package main

import (
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/web"
    "github.com/micro/go-plugins/registry/consul/v2"
    "go-micro-demos/http/gin2/service"
    "log"
    "net/http"
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
    // 路由分组
    GroupV1 := ginRouter.Group("/v1")
    GroupV1.Handle("POST", "/house", func(ctx *gin.Context) {
        req := service.Req{}

        err := ctx.Bind(&req) // 接收 body 里的信息，也可以有其他参数接收方式 ，比如：/house/:num ， param 接收
        if err != nil || req.Num < service.HouseMinNum || req.Num > service.HouseMaxNum {
            req = service.Req{service.HouseDefaultNum}
        }
        ctx.JSON(http.StatusOK, gin.H{
            "data": service.BuildHouse(req.Num),
            "msg": "success",
        })
    })

    // 用 go-micro 里的 web 创建 server
    server := web.NewService(
        web.Name("house-service"),
        web.Handler(ginRouter),
        web.Metadata(map[string]string{"protocol": "http"}),
        web.Registry(consulReg),
    )
    server.Init()

    if err := server.Run(); err != nil{
        log.Fatal(err)
    }
}
