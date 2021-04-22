package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/client"
    proto "go-micro-demos/greeter/proto"
    "log"
)
// log wrapper logs every time a request is made
type logWrapper struct {
    client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
    fmt.Printf("[wrapper] client request service: %s method: %s\n", req.Service(), req.Endpoint())
    return l.Client.Call(ctx, req, rsp)
}

// Implements client.Wrapper as logWrapper
func logWrap(c client.Client) client.Client {
    return &logWrapper{c}
}

func main() {
    service := micro.NewService(
        micro.Name("greeter.client"),
        micro.WrapClient(logWrap),
    )

    service.Init()

    greeter := proto.NewGreeterService("greeter", service.Client())

    rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Jimmy"})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(rsp.Greeting)
}