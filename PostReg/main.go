package main

import (
	"sss/PostReg/handler"
	"sss/PostReg/subscriber"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	POSTREG "sss/PostReg/proto/PostReg"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostReg"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTREG.RegisterPostRegHandler(service.Server(), new(handler.PostReg))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostReg", service.Server(), new(subscriber.PostReg))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostReg", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
