package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/v2/web"
)

func main() {
    // gin 作为路由
    ginRouter := gin.Default()
    ginRouter.GET("/hello", func(ctx *gin.Context) {
        ctx.String(http.StatusOK, "hello gin!")
    })
    // 初始化 go-micro
    service := web.NewService(
        web.Address(":8080"),  // http 端口
        web.Metadata(map[string]string{"data": "hello world"}), // 可以携带一些信息
        web.Handler(ginRouter), // gin 路由
    )
    // 开始运行
    if err := service.Run(); err != nil {
        log.Println(err.Error())
    }
}
