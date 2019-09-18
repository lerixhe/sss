package main

import (
	"sss/GetSession/handler"
	"sss/GetSession/subscriber"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	GETSESSION "sss/GetSession/proto/GetSession"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.GetSession"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETSESSION.RegisterGetSessionHandler(service.Server(), new(handler.GetSession))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetSession", service.Server(), new(subscriber.GetSession))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetSession", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
