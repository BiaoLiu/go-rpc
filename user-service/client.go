package main

import (
	"context"
	"fmt"
	"go-rpc/user-service/proto"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50053",grpc.WithInsecure())
	if err != nil {
		//log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	resp, err := client.NewUser(ctx, &pb.NewUserRequest{Username: "LBI-GRPC", Email: "45238107@qq.com", Password: "test"})
	fmt.Println("new user id:", resp.Id)
}
