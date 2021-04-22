package main

import (
	micro "github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"go-micro-demos/hello/handler"
	"go-micro-demos/hello/subscriber"

	hello "go-micro-demos/hello/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.hello"),
		micro.Version("latest"),
	) // Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.hello", service.Server(), new(subscriber.Hello))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
