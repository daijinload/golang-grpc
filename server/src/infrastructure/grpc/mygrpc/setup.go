package mygrpc

import (
	"log"
	"time"

	c4 "golang-grpc/server/src/controller/bidirectionalstream"
	c2 "golang-grpc/server/src/controller/clientstream"
	c1 "golang-grpc/server/src/controller/helloworld"
	c3 "golang-grpc/server/src/controller/serverstream"

	"golang-grpc/server/src/infrastructure/grpc/gproto/bidirectionalstream"
	"golang-grpc/server/src/infrastructure/grpc/gproto/clientstream"
	"golang-grpc/server/src/infrastructure/grpc/gproto/helloworld"
	"golang-grpc/server/src/infrastructure/grpc/gproto/serverstream"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// setup .
func setup(isTLS bool) *grpc.Server {

	// gRPCオプション設定
	opts := make([]grpc.ServerOption, 0)
	if isTLS { // TLSで起動する場合、credentialsを設定
		creds, err := credentials.NewServerTLSFromFile(
			"server.crt",
			"server.key",
		)
		if err != nil {
			log.Fatalf("creds error!!!: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	// opts = append(opts, grpc.UnaryInterceptor(
	// 	grpc_middleware.ChainUnaryServer(
	// 		unaryinterceptor,
	// 	),
	// ))

	// opts = append(opts, grpc.StreamInterceptor(
	// 	grpc_middleware.ChainStreamServer(
	// 		streaminterceptor,
	// 	),
	// ))

	// gRPCの接続パラメータ
	// keepaliveの設定は、下記を参照してください
	// See https://github.com/grpc/grpc-go/tree/master/examples/features/keepalive
	opts = append(
		opts,
		grpc.ConnectionTimeout(120*time.Second),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
			PermitWithoutStream: true,            // Allow pings even when there are no active streams
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			// リクエストの処理途中でcontextのcancelが発生するのを防ぐため、タイムアウトの時間は長めに取っている。
			// しかし、タイムアウト60秒だとコネクション時間が長過ぎるため、GOAWAYは早めに送る設定にしている。
			// （GOAWAYはサーバが送るし、処理中にcontext cancelが出ることも無いため）
			MaxConnectionIdle:     15 * time.Second, // クライアントidle時間（GOAWAY送る）
			MaxConnectionAge:      15 * time.Second, // コネクション維持時間（GOAWAY送る）
			MaxConnectionAgeGrace: 60 * time.Second, // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
			Time:                  60 * time.Second, // サーバからクライアントにPingを送る間隔（このpingで生存確認している）
			Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
		}),
	)

	grpcServer := grpc.NewServer(opts...)

	helloworld.RegisterGreeterServer(grpcServer, &c1.GreeterController{})
	clientstream.RegisterClientStreamTestServer(grpcServer, &c2.ClientStreamController{})
	serverstream.RegisterServerStreamTestServer(grpcServer, &c3.ServerStreamController{})
	bidirectionalstream.RegisterBidirectionalStreamTestServer(grpcServer, &c4.BidirectionalStreamController{})

	// ※ grpcurlで、メソッドリストを表示するため必要（本番では不要）
	reflection.Register(grpcServer)

	return grpcServer
}
