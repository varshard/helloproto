package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/varshard/helloproto/rest/addressbook"
	"google.golang.org/grpc"
)

var grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9000", "gRPC server endpoint")

func runHttpServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterAddressBookHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8081", mux)
}

func StartHttpServer() error {
	flag.Parse()

	if err := runHttpServer(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := StartHttpServer(); err != nil {
		panic(err)
	}
}
