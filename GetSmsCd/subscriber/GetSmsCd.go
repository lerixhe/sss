package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	GETSMSCD "sss/GetSmsCd/proto/GetSmsCd"
)

type GetSmsCd struct{}

func (e *GetSmsCd) Handle(ctx context.Context, msg *GETSMSCD.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *GETSMSCD.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
