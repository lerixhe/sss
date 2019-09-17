package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	GETIMAGECD "sss/GetImageCd/proto/GetImageCd"
)

type GetImageCd struct{}

func (e *GetImageCd) Handle(ctx context.Context, msg *GETIMAGECD.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *GETIMAGECD.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
