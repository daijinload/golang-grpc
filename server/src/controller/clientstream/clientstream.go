package clientstream

import (
	"io"
	"log"
	"strings"

	"golang-grpc/server/src/infrastructure/grpc/gproto/clientstream"
	"golang-grpc/server/src/infrastructure/grpc/gproto/helloworld"
)

// ClientStreamController .
type ClientStreamController struct{}

// SayHello .
func (s *ClientStreamController) SayHello(stream clientstream.ClientStreamTest_SayHelloServer) error {
	var messages []string
	for {
		msg, err := stream.Recv()
		log.Print("cliant stream server recived:", msg, err)
		if err == io.EOF {
			return stream.SendAndClose(&helloworld.HelloReply{
				Message: strings.Join(messages, ","),
			})
		}

		messages = append(messages, msg.String())

		if err != nil {
			return err
		}
	}
}
