package main

import (
	proto "go-rpc/proto/helloworld"
	"context"
	"google.golang.org/grpc"
	"net"
	"google.golang.org/grpc/grpclog"
	"fmt"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

type Greeter struct{}

func (g *Greeter) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	resp := new(proto.HelloReply)
	resp.Message = "Hello grpc!"
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()


	proto.RegisterGreeterServer(s, new(Greeter))

	grpclog.Infoln("Listen on " + Address)

	fmt.Println("Listen on " + Address)

	s.Serve(listen)
}
