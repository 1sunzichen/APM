package main

import (
	"dogapm"
	"net/http"
	"ordersvc/api"
	"ordersvc/grpclient"
	protos "proto"
)

func main() {

	dogapm.Infra.Init(
		dogapm.InfraDbOption("root:root@tcp(127.0.0.1:3306)/ordersvc"),
	)

	grpclient.SkuClient = protos.NewSkuServiceClient(dogapm.NewGrpcClient(":8081"))
	grpclient.UserClient = protos.NewUserServiceClient(dogapm.NewGrpcClient(":8082"))
	httpServer := dogapm.NewHttpServer(":8123")
	dogapm.Logger.Info(nil, "start http server", map[string]interface{}{"addr": ":8080"})
	httpServer.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
		return
	})

	httpServer.HandleFunc("/order/add", api.Order.Add)
	// httpServer.Start()
	// time.Sleep(time.Second * 300)
	dogapm.EndPointInstance.Start()

}
