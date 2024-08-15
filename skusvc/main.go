package main

import (
	"dogapm"
	protos "proto"
	"skusvc/grpc"
)

func main() {
	dogapm.Infra.Init(
		dogapm.InfraDbOption("root:root@tcp(127.0.0.1:3306)/skusvc"),
	)
	grpcServer := dogapm.NewGrpcServer(":8081")
	protos.RegisterSkuServiceServer(grpcServer, &grpc.SkuServer{})
	dogapm.EndPointInstance.Start()
}
