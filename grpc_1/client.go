package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	proto "go-rpc/proto/helloworld"
	"context"
	"fmt"
)

const ServerAddress = "127.0.0.1:50052"

func main() {
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	req := new(proto.HelloRequest)
	req.Name = "hello request"
	resp, err := c.SayHello(context.Background(), req)
	fmt.Println(resp.Message)
}
