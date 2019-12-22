package bidirectionalstream

import (
	"io"
	"log"
	"strings"

	"golang-grpc/server/src/infrastructure/grpc/gproto/bidirectionalstream"
	"golang-grpc/server/src/infrastructure/grpc/gproto/helloworld"
)

// BidirectionalStreamController .
type BidirectionalStreamController struct{}

// SayHello .
func (s *BidirectionalStreamController) SayHello(stream bidirectionalstream.BidirectionalStreamTest_SayHelloServer) error {
	var messages []string
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Print(messages)
			stream.Send(&helloworld.HelloReply{
				Message: strings.Join(messages, ","),
			})
			return nil
		}
		if err != nil {
			return err
		}
		//messages = append(messages, in.Name)
		stream.Send(&helloworld.HelloReply{
			Message: in.Name,
		})
	}
}
