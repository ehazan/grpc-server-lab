package main

import (
	mygrpc "github.com/ehazan/grpc-server-lab/internal/adapter/grpc"
	app "github.com/ehazan/grpc-server-lab/internal/application"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	hs := &app.HelloService{}
	grpcAdapter := mygrpc.NewGrpcAdapter(hs, 9090)

	grpcAdapter.Run()
}
