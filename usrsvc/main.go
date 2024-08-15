package main

import (
	"dogapm"
	protos "proto"
	"usrsvc/grpc"
)

func main() {
	dogapm.Infra.Init(
		dogapm.InfraDbOption("root:root@tcp(127.0.0.1:3306)/usrsvc"),
		dogapm.InfraRDBOption("127.0.0.1:6379"),
	)
	grpcServer := dogapm.NewGrpcServer(":8082")
	protos.RegisterUserServiceServer(grpcServer, &grpc.UserServer{})
	dogapm.EndPointInstance.Start()
}
