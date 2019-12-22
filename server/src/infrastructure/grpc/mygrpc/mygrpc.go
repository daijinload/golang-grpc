package mygrpc

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/armon/go-proxyproto"
)

// Serve .
func Serve(port string, isTLS bool) {

	grpcServer := setup(isTLS)

	go func() {
		lis, err := net.Listen("tcp4", ":"+port)
		if err != nil {
			log.Fatalf("Failed to parse listener: %v", err)
		}
		// 本番はLBからプロキシ情報が送られてくるため、先頭にあるプロキシ情報を取得＆削除する必要がある
		if isTLS {
			lis = &proxyproto.Listener{Listener: lis}
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("grpc server serve error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Gracefull Shutdown start ...")

	grpcServer.GracefulStop()
}
