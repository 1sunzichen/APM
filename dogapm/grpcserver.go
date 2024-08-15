package dogapm

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	*grpc.Server
	addr string
}

func NewGrpcServer(addr string) *GrpcServer {
	svc := grpc.NewServer()
	server := &GrpcServer{
		Server: svc,
		addr:   addr,
	}
	globalCloses = append(globalCloses, server)
	globalStarters = append(globalStarters, server)
	return server
}

func (g *GrpcServer) Close() {
	g.Server.GracefulStop()
}

func unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		return handler(ctx, req)
	}
}

func (g *GrpcServer) Start() {

	l, err := net.Listen("tcp", g.addr)
	if err != nil {
		panic(err)
	}
	go func() {
		err = g.Serve(l)
		if err != nil {
			panic(err)
		}
	}()
}
