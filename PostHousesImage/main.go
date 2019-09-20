package main

import (
	"sss/PostHousesImage/handler"
	"sss/PostHousesImage/subscriber"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	POSTHOUSESIMAGE "sss/PostHousesImage/proto/PostHousesImage"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostHousesImage"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	POSTHOUSESIMAGE.RegisterPostHousesImageHandler(service.Server(), new(handler.PostHousesImage))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostHousesImage", service.Server(), new(subscriber.PostHousesImage))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.PostHousesImage", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
