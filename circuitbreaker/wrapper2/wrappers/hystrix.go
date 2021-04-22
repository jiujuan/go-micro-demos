package wrappers

import (
    "context"
    "fmt"
    "github.com/afex/hystrix-go/hystrix"
    "github.com/micro/go-micro/v2/client"
)

type House struct {
    ID int
    Name string
}

func defaultHouse() {
    fmt.Println("circuit breaker method")
    house := House{ID: 0, Name: "default house"}
    fmt.Println(house)
}

type HouseWrapper struct {
    client.Client
}

// wrapper 方法
func (house *HouseWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
    // command 名称
    commandName := fmt.Sprintf("%s.%s", req.Service(), req.Endpoint())
    // 1.配置 config
    configOne := hystrix.CommandConfig{
        Timeout: 800,
    }

    // 2.配置 command
    hystrix.ConfigureCommand(commandName, configOne)

    // 3.执行，使用 Do 方法
    return hystrix.Do(
        commandName,
        func() error {
            // 正常就继续执行
            return house.Client.Call(ctx, req, resp, opts...)
        },
        func(err error) error {
            // 如果熔断了，调用默认函数
            defaultHouse()
            return nil
        },
    )
}

func NewHouseWrapper(c client.Client) client.Client {
    return &HouseWrapper{c}
}
