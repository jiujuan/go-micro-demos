package main

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/web"
    "github.com/micro/go-plugins/registry/consul/v2"
    "github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
    house "go-micro-demos/circuitbreaker/wrapper/proto/house"
    "log"
    "net/http"
    "time"
)

func main() {
    consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
    // new service
    service := micro.NewService(
        micro.Name("go.micro.srv.client.house"),
        micro.Registry(consulReg),
        micro.WrapClient(
            // hystrix 的 wrapper 包
            hystrix.NewClientWrapper(),
        ),
    )
    houseService := house.NewHouseService("go.micro.srv.house", service.Client())

    ginHandler := gin.Default()
    // 真正运行的服务
    webServer := web.NewService(
        web.Name("go.micro.web.house"),
        web.Address(":8001"),
        web.Handler(ginHandler),
        web.Registry(consulReg),
    )

    // web 路由
    webRouter(ginHandler, houseService)

    webServer.Init()
    if err := webServer.Run(); err != nil{
        log.Fatal(err)
    }
}

var hservice house.HouseService
func webRouter(gin *gin.Engine, houseService house.HouseService){
    hservice = houseService
    v1 := gin.Group("/v1/house")
    {
        v1.POST("/create", buildHouse)
        v1.GET("/get", getHouse)
    }
}

func buildHouse(ctx *gin.Context) {
    req := new(house.RequestData)

    if err := ctx.BindJSON(req); err != nil {
        log.Println("request param: ", err)
        return
    }

    // 这里调用 HouseService 的方法(proto 生成的方法)
    // 这里模拟调用，并没有真正调用
    hservice.Build(context.Background(), req)

    resp := house.ResponseMsg{Msg: "build one house 1"}
    ctx.JSON(http.StatusOK, gin.H{
        "message": resp.Msg,
    })
}

func getHouse(ctx *gin.Context) {
    req := new(house.RequestData)
    if err := ctx.BindJSON(req); err != nil {
        log.Println("get house param error: ", err.Error())
        return
    }

    hservice.GetHouse(context.Background(), req)

    resp := house.ResponseMsg{Msg: "get house: "+req.Name}

    time.Sleep( time.Second * 5)

    ctx.JSON(http.StatusOK, gin.H{
        "message": resp.Msg,
    })
}
