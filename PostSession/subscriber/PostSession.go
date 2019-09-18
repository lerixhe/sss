package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	POSTSESSION "sss/PostSession/proto/PostSession"
)

type PostSession struct{}

func (e *PostSession) Handle(ctx context.Context, msg *POSTSESSION.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *POSTSESSION.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
