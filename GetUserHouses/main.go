package main

import (
	"sss/GetUserHouses/handler"
	"sss/GetUserHouses/subscriber"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	GETUSERHOUSES "sss/GetUserHouses/proto/GetUserHouses"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetUserHouses"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETUSERHOUSES.RegisterGetUserHousesHandler(service.Server(), new(handler.GetUserHouses))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetUserHouses", service.Server(), new(subscriber.GetUserHouses))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetUserHouses", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
