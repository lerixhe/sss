package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	POSTHOUSESIMAGE "sss/PostHousesImage/proto/PostHousesImage"
)

type PostHousesImage struct{}

func (e *PostHousesImage) Handle(ctx context.Context, msg *POSTHOUSESIMAGE.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *POSTHOUSESIMAGE.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
