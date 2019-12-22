package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"

	"golang-grpc/server/src/infrastructure/grpc/gproto/helloworld"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	isTLS   = flag.Bool("istls", false, "true is tls on")
	address = flag.String("address", "127.0.0.1:8443", "example -> 127.0.0.1:8443")
)

func main() {
	start := time.Now()
	fmt.Println(start)

	if *isTLS {
		tlsRequest(*address)
	} else {
		simpleRequest(*address)
	}

	fmt.Println(time.Now())
	fmt.Println(convertTimeToMilli(time.Now()) - convertTimeToMilli(start))
}

func simpleRequest(address string) {
	ctx := context.Background()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	defer conn.Close()

	c := helloworld.NewGreeterClient(conn)

	result, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: "simple-request"})
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Println(result)
}

func tlsRequest(address string) {
	ctx := context.Background()

	dialerFunc := func(ctx context.Context, address string) (net.Conn, error) {
		netConn, err := (&net.Dialer{}).DialContext(ctx, "tcp", address)
		if err != nil {
			return nil, err
		}
		creds := credentials.NewTLS(&tls.Config{})

		netConn2, _, err := creds.ClientHandshake(ctx, address, netConn)
		if err != nil {
			return nil, err
		}
		return netConn2, nil
	}

	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithBlock(),
		grpc.FailOnNonTempDialError(true),
		grpc.WithContextDialer(dialerFunc),
		grpc.WithInsecure(), // we are handling TLS, so tell grpc not to
	)
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		log.Fatalf("failed to DialContext: %+v\n", err)
	}
	defer conn.Close()

	c := helloworld.NewGreeterClient(conn)

	result, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: "tls-request"})
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Println(result)
}

func convertTimeToMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
