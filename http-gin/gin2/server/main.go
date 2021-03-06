package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/web"
    "github.com/micro/go-plugins/registry/consul/v2"
    "go-micro-demos/http-gin/gin2/service"
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

    // gin 初始化路由
    ginRouter := gin.Default()
    // 路由分组
    GroupV1 := ginRouter.Group("/v1")
    GroupV1.Handle("POST", "/house", func(ctx *gin.Context) {
        houseNum := service.Num{}

        err := ctx.Bind(&houseNum) // 接收 body 里的信息，也可以有其他参数接收方式 ，比如：/house/:num ， param 接收
        if err != nil {
            houseNum = service.Num{service.HouseDefaultNum}
        }
        // 这里可以进行一些逻辑处理
        ctx.JSON(http.StatusOK, gin.H{
            "data": service.BuildHouse(houseNum.Num),
            "msg": "success",
        })
    })

    // 也可以是多路由分组
    GroupV2 := ginRouter.Group("/v2")
    {
        GroupV2.GET("/house/:name", func(ctx *gin.Context) {
            name := ctx.Param("name")

            // 这里进行逻辑处理

            ctx.String(http.StatusOK, "[GET] req name: ", name)
        })

        GroupV2.POST("/house", func(ctx *gin.Context) {
            house := new(service.House)
            err := ctx.Bind(house)
            if err != nil {
                fmt.Println(err)
                return
            }
            // 这里进行逻辑处理
            ctx.JSON(http.StatusOK, gin.H{
                "data": fmt.Sprintf("%d - %s", house.ID, house.Name),
            })
        })
    }

    // 用 go-micro 里的 web 创建 server
    server := web.NewService(
        web.Name("house-service"), // 注册服务名称到 consul
        web.Handler(ginRouter), // 配置 gin 路由
        web.Metadata(map[string]string{"protocol": "http"}),
        web.Registry(consulReg),
    )
    server.Init()

    if err := server.Run(); err != nil{
        log.Fatal(err)
    }
}
