package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	PostUserAuth "sss/PostUserAuth/proto/PostUserAuth"
)

type PostUserAuth struct{}

func (e *PostUserAuth) Handle(ctx context.Context, msg *PostUserAuth.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *PostUserAuth.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
