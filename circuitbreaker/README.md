# 服务治理-熔断

> go-micro v2.9.1

在 go-micro 里封装了 2 个熔断的包，位于 github.com\micro\go-plugins\wrapper\breaker 目录下:

[hystrix-go](https://github.com/afex/hystrix-go/) 和 [gobreaker](https://github.com/sony/gobreaker)

这 2 个插件都是在 wrapper 包下，它们使用包装机制实现的。

实现熔断机制，就可以利用 go-micro 提供的这些包。

> 说明：熔断功能是在客户端使用


## simple/demo.go

写一个简单的demo看看 hystrix 的使用。

使用方法：

```go
package main

import (
    "github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
    "github.com/micro/go-micro/v2"

    hysgo "github.com/afex/hystrix-go/hystrix"
)

func main() {
    service := micro.NewService(
        micro.Name("go.micro.foo.breaker.demo"),
        // 只需要在创建 service 时候指定 hystrix 插件，系统便有了自动熔断能力
        micro.WrapClient(hystrix.NewClientWrapper()),
    )

    service.Init()

    // 修改 hystrix 的默认配置
    hysgo.DefaultMaxConcurrent = 3  // 修改最大并发
    hysgo.DefaultTimeout = 300 // 修改 timeout 时间到 300 milliseconds（毫秒）
}
```

## wrapper
