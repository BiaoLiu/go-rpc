package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	greeter "go-rpc/greeter_service/proto"
	"go-rpc/user-service/proto"
	"google.golang.org/grpc"
	"io"
	"net"
	"time"
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

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor()))
	if err != nil {
		fmt.Println("connect greeter service fail...")
	}
	defer conn.Close()

	greeteService := greeter.NewGreeterClient(conn)
	resp, err := greeteService.SayHello(context.Background(), &greeter.HelloRequest{Name: username})
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
	tracer, closer := initJaeger("UserService")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	fmt.Println("jaeger init success")

	fmt.Println("user service start...")

	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_opentracing.UnaryServerInterceptor()))
	pb.RegisterUserServer(server, &srv)

	if err := server.Serve(listener); err != nil {
		fmt.Println("server start error...")
	}
}

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "140.143.56.14:6831",
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		fmt.Println("error:" + err.Error())
	}
	fmt.Println("tracer:", tracer)
	return tracer, closer
}
