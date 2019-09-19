package main

import (
	"sss/GetUserInfo/handler"
	"sss/GetUserInfo/subscriber"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	GETUSERINFO "sss/GetUserInfo/proto/GetUserInfo"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetUserInfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GETUSERINFO.RegisterGetUserInfoHandler(service.Server(), new(handler.GetUserInfo))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetUserInfo", service.Server(), new(subscriber.GetUserInfo))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetUserInfo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
