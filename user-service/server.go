package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"go-rpc/user-service/proto"
	"go-rpc/plugins"
	"github.com/opentracing/opentracing-go"
)

type userServie struct {
}

func (userServie) NewUser(ctx context.Context, req *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	username := req.Username
	email := req.Email
	password := req.Password

	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := span.Tracer().StartSpan("NewUser", opentracing.ChildOf(span.Context()))
		span.SetTag("User", ctx.Value("NewUser"))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	dialOpts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	dialOpts = append(dialOpts, grpc.WithUnaryInterceptor(plugins.OpenTracingClientInterceptor(opentracing.SpanFromContext(ctx).Tracer())))

	conn, err := grpc.Dial("localhost:50051", dialOpts...)
	if err != nil {
		fmt.Println("connect greeter service fail...")
	}

	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: username})
	if err != nil {
		fmt.Println("call sayhello service fail...")
	}
	fmt.Println("sayhello result", resp.Message)

	fmt.Println("username:", username, "email:", email, "password:", password)
	return &pb.NewUserResponse{Id: "1"}, nil
}

func (userServie) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.UserResponse, error) {
	email := req.Email
	return &pb.UserResponse{Id: "1", Username: "LBI", Email: email}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		fmt.Println("failed to listen")
	}
	srv := userServie{}
	tracer, closer := plugins.InitJaeger("UserServer")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	fmt.Println("jaeger init success")

	fmt.Println("user service start...")
	var serverOpts []grpc.ServerOption
	serverOpts = append(serverOpts, grpc.UnaryInterceptor(plugins.OpentracingServerInterceptor(tracer)))

	server := grpc.NewServer(serverOpts...)
	pb.RegisterUserServer(server, &srv)

	if err := server.Serve(listener); err != nil {
		fmt.Println("server start error...")
	}
}
