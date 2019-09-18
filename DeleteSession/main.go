package main

import (
	"sss/DeleteSession/handler"
	"sss/DeleteSession/subscriber"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	DeleteSession "sss/DeleteSession/proto/DeleteSession"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.DeleteSession"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	DeleteSession.RegisterDeleteSessionHandler(service.Server(), new(handler.DeleteSession))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.DeleteSession", service.Server(), new(subscriber.DeleteSession))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.DeleteSession", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
