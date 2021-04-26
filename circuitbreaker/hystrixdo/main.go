package main

import (
    "context"
    "errors"
    "fmt"
    "github.com/afex/hystrix-go/hystrix"
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/web"
    "github.com/micro/go-plugins/registry/consul/v2"
    proto "go-micro-demos/circuitbreaker/hystrixdo/proto/house"
    "log"
    "net/http"
)

func main() {
    consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
    // new service
    service := micro.NewService(
        micro.Name("go.micro.service.client.house"),
        micro.Registry(consulReg),
    )
    // 这个 houseService 涉及到的是后面 rpc 的请求服务，houseService 起着http到rpc的连接作用
    // http param -> gin webRouter() -> gin getHouse() -> houseService.GetHouse()
    // 这里 hystrix 设置熔断也是对 GetHouse() 这个服务方法设置熔断，是对服务的熔断
    houseService := proto.NewHouseService("go.micro.service.house", service.Client())

    // 真正运行的服务
    webServer := web.NewService(
        web.Name("go.micro.web.house"),
        web.Address(":8001"),
        web.Handler(webRouter(houseService)),
        web.Registry(consulReg),
    )

    webServer.Init()
    if err := webServer.Run(); err != nil{
        log.Fatal(err)
    }
}

func webRouter(houseService proto.HouseService) *gin.Engine {
    router := gin.Default()
    router.Use(func(ctx *gin.Context) {
        ctx.Set("houseService", houseService)
    })

    v1 := router.Group("/v1/house")
    {
        v1.GET("/get", getHouse)
        v1.POST("/create", buildHouse)
    }

    return router
}

// 给这个方法设置 hystrix 处理命令
func getHouse(ctx *gin.Context) {
    req := new(proto.RequestData)
    if err := ctx.BindJSON(req); err != nil {// 获取客户端请求数据并绑定，如果不为 nil 则返回错误信息
        log.Println("get house param error: ", err.Error())
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": err.Error(),
        })
        return
    }

    // 设置一个超时时间
    configOne := hystrix.CommandConfig{
        Timeout: 3000,
    }
    // 配置 command
    hystrix.ConfigureCommand("getHouse", configOne)

    var houseService proto.HouseService
    if hservice, ok := ctx.Get("houseService"); ok {
        houseService = hservice.(proto.HouseService)
    }
    // 执行 Do 方法
    var respMsg *proto.ResponseMsg
    err := hystrix.Do(
        "getHouse",
        func() error {
            var err error
            respMsg, err = houseService.GetHouse(context.Background(), req)
            return err
        },
        func(err error) error {
            errstr := fmt.Sprintf("%s, %v","hystrix error msg: ", err.Error())
            return errors.New(errstr)
            //return nil
        },
    )

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": err.Error(),
        })
    } else {
        ctx.JSON(http.StatusOK, gin.H{
            "data": respMsg,
            "message" : "success",
        })
    }
}

func buildHouse(ctx *gin.Context) {
    req := new(proto.RequestData)

    if err := ctx.BindJSON(req); err != nil {
        log.Println("request param: ", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": err.Error(),
        })
        return
    }

    // 这里调用 HouseService 的方法(proto 生成的方法)
    // 这里模拟调用，并没有真正调用
    var houseService proto.HouseService
    if hservice, ok := ctx.Get("houseService"); ok {
        houseService = hservice.(proto.HouseService)
    }
    houseService.Build(context.Background(), req)

    resp := proto.ResponseMsg{Msg: "build one house 1"}
    ctx.JSON(http.StatusOK, gin.H{
        "message": resp.Msg,
    })
}
