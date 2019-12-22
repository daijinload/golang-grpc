package serverstream

import (
	"log"
	"time"

	"golang-grpc/server/src/infrastructure/grpc/gproto/helloworld"
	"golang-grpc/server/src/infrastructure/grpc/gproto/serverstream"
)

// ServerStreamController .
type ServerStreamController struct{}

// SayHello .
func (s *ServerStreamController) SayHello(req *helloworld.HelloRequest, stream serverstream.ServerStreamTest_SayHelloServer) error {
	log.Print("server stream server recived:", req)

	var err error

	err = stream.Send(&helloworld.HelloReply{Message: req.GetName() + "aaa1"})
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)

	err = stream.Send(&helloworld.HelloReply{Message: req.GetName() + "aaa2"})
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)

	err = stream.Send(&helloworld.HelloReply{Message: req.GetName() + "aaa3"})
	if err != nil {
		return err
	}
	return nil
}
