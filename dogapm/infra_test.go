package dogapm

import (
	"context"
	"fmt"
	"net/http"
	protos "proto"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Infra.Init(InfraDbOption("root:root@tcp(127.0.0.1:3306)/mydb"),
		InfraRDBOption("127.0.0.1:6379"))
}

func TestNewHttpServer(t *testing.T) {
	s := NewHttpServer(":8123")
	s.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	s.Start()
	time.Sleep(time.Second * 300)
	// s.Stop()
}

type HelloSvc struct {
	protos.UnimplementedHelloServiceServer
}

func (h *HelloSvc) Receive(ctx context.Context, msg *protos.HelloMsg) (*protos.HelloMsg, error) {
	return msg, nil
}
func TestGrpc(t *testing.T) {
	go func() {
		s := NewGrpcServer(":8080")
		protos.RegisterHelloServiceServer(s, &HelloSvc{})
		s.Start()
	}()
	client := NewGrpcClient("127.0.0.1:8080")
	res, err := protos.NewHelloServiceClient(client).Receive(context.Background(), &protos.HelloMsg{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Msg)
}
