package dogapm

import (
	"os"
	"os/signal"
	"syscall"
)

type starter interface {
	Start()
}

type closer interface {
	Close()
}

var (
	globalStarters = make([]starter, 0)
	globalCloses   = make([]closer, 0)
)

type EndPoint struct {
	Stop chan int
}

var EndPointInstance = &EndPoint{Stop: make(chan int)}

func (e *EndPoint) Start() {
	for _, com := range globalStarters {
		com.Start()
	}
	go func() {
		//等待关闭信号
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit
		e.ShutDown()

	}()
	<-e.Stop
}
func (e *EndPoint) ShutDown() {
	for _, com := range globalCloses {
		com.Close()
	}
	e.Stop <- 1
}

func (e *EndPoint) Close() {
	for _, com := range globalCloses {
		com.Close()
	}
}
