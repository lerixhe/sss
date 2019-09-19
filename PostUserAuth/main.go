package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"sss/PostUserAuth/handler"
	"sss/PostUserAuth/subscriber"

	PostUserAuth "sss/PostUserAuth/proto/PostUserAuth"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.PostUserAuth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	PostUserAuth.RegisterPostUserAuthHandler(service.Server(), new(handler.PostUserAuth))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostUserAuth", service.Server(), new(subscriber.PostUserAuth))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostUserAuth", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
