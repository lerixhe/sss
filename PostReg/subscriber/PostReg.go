package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	POSTREG "sss/PostReg/proto/PostReg"
)

type PostReg struct{}

func (e *PostReg) Handle(ctx context.Context, msg *POSTREG.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *POSTREG.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
