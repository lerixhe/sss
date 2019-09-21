package main

import (
	"sss/GetHouseInfo/handler"
	"sss/GetHouseInfo/subscriber"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	GETHOUSEINFO "sss/GetHouseInfo/proto/GetHouseInfo"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetHouseInfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETHOUSEINFO.RegisterGetHouseInfoHandler(service.Server(), new(handler.GetHouseInfo))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetHouseInfo", service.Server(), new(subscriber.GetHouseInfo))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetHouseInfo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
