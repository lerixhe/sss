package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	PostUserAuth "sss/PostUserAuth/proto/PostUserAuth"
)

type PostUserAuth struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PostUserAuth) Call(ctx context.Context, req *PostUserAuth.Request, rsp *PostUserAuth.Response) error {
	log.Log("Received PostUserAuth.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *PostUserAuth) Stream(ctx context.Context, req *PostUserAuth.StreamingRequest, stream PostUserAuth.PostUserAuth_StreamStream) error {
	log.Logf("Received PostUserAuth.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&PostUserAuth.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *PostUserAuth) PingPong(ctx context.Context, stream PostUserAuth.PostUserAuth_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&PostUserAuth.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
