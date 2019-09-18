package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	DeleteSession "sss/DeleteSession/proto/DeleteSession"
)

type DeleteSession struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *DeleteSession) Call(ctx context.Context, req *DeleteSession.Request, rsp *DeleteSession.Response) error {
	log.Log("Received DeleteSession.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *DeleteSession) Stream(ctx context.Context, req *DeleteSession.StreamingRequest, stream DeleteSession.DeleteSession_StreamStream) error {
	log.Logf("Received DeleteSession.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&DeleteSession.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *DeleteSession) PingPong(ctx context.Context, stream DeleteSession.DeleteSession_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&DeleteSession.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
