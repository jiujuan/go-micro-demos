package simple

import (
    "github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
    "github.com/micro/go-micro/v2"

    hystrixgo "github.com/afex/hystrix-go/hystrix"
)

func main() {
    service := micro.NewService(
        micro.Name("go.micro.service.hystrix.demo"),
        // 只需要在创建 service 时候指定 hystrix 插件，系统便有了自动熔断能力
        micro.WrapClient(hystrix.NewClientWrapper()),
    )

    service.Init()

    // 修改 hystrix 的默认配置
    hystrixgo.DefaultMaxConcurrent = 3  // 修改最大并发
    hystrixgo.DefaultTimeout = 300 // 修改 timeout 时间到 300 milliseconds（毫秒）

    service.Run()
}
