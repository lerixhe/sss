package main

import (
	"sss/PostUserAuth/handler"
	"sss/PostUserAuth/subscriber"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	POSTUSERAUTH "sss/PostUserAuth/proto/PostUserAuth"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostUserAuth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTUSERAUTH.RegisterPostUserAuthHandler(service.Server(), new(handler.PostUserAuth))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostUserAuth", service.Server(), new(subscriber.PostUserAuth))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostUserAuth", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
