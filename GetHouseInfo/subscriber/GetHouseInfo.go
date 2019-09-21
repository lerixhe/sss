package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	GETHOUSEINFO "sss/GetHouseInfo/proto/GetHouseInfo"
)

type GetHouseInfo struct{}

func (e *GetHouseInfo) Handle(ctx context.Context, msg *GETHOUSEINFO.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *GETHOUSEINFO.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
