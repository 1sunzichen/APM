package dogapm

import (
	"context"
	"net/http"
)

type HttpServer struct {
	mux *http.ServeMux
	*http.Server
}

func NewHttpServer(addr string) *HttpServer {
	mux := http.NewServeMux()
	server := &http.Server{Addr: addr, Handler: mux}
	s := &HttpServer{mux: mux, Server: server}
	globalStarters = append(globalStarters, s)
	globalCloses = append(globalCloses, s)
	return &HttpServer{mux: mux, Server: server}
}

func (h *HttpServer) Handle(parrern string, handler http.Handler) {
	h.mux.Handle(parrern, handler)
}

func (h *HttpServer) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	h.mux.HandleFunc(pattern, handler)
}

func (h *HttpServer) Start() {
	go func() {
		err := h.ListenAndServe()
		if err != nil {
			panic(err)
		}

	}()
}

func (h *HttpServer) Close() {
	h.Shutdown(context.TODO())
}
