package helloworld

import (
	"context"
	"log"

	"golang-grpc/server/src/infrastructure/grpc/gproto/helloworld"

	"google.golang.org/grpc/peer"
)

// GreeterController .
type GreeterController struct{}

// SayHello .
func (s *GreeterController) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	p, _ := peer.FromContext(ctx)
	response := new(helloworld.HelloReply)
	response.Message = req.Name + p.Addr.String()
	log.Print(p.Addr.String())
	return response, nil
}
