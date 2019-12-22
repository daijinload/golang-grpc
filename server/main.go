package main

import (
	"flag"
	"golang-grpc/server/src/infrastructure/grpc/mygrpc"
	"log"
)

var (
	port  = flag.String("port", "8443", "grpc server port")
	isTLS = flag.Bool("istls", false, "true is tls on")
)

func main() {
	flag.Parse()

	log.Print("start")

	mygrpc.Serve("8443", *isTLS)
}
