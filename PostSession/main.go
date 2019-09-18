package main

import (
	"sss/PostSession/handler"
	"sss/PostSession/subscriber"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	POSTSESSION "sss/PostSession/proto/PostSession"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostSession"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTSESSION.RegisterPostSessionHandler(service.Server(), new(handler.PostSession))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostSession", service.Server(), new(subscriber.PostSession))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostSession", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
