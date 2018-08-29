package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "go-rpc/proto/helloworld"
	"google.golang.org/grpc/reflection"
	"time"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"fmt"
	"github.com/uber/jaeger-client-go"
	"io"
	"go-rpc/plugins"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := span.Tracer().StartSpan("SayHello", opentracing.ChildOf(span.Context()))
		span.SetTag("Hello", ctx.Value("Hello"))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	time.Sleep(time.Second * 2)

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tracer, closer := initJaeger("HelloServer")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	fmt.Println("jaeger init success")

	var serverOpts []grpc.ServerOption
	serverOpts = append(serverOpts, grpc.UnaryInterceptor(plugins.OpentracingServerInterceptor(tracer)))

	s := grpc.NewServer(serverOpts...)
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
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
