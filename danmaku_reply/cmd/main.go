package main

import (
	"danmaku/danmaku_reply/model"
	"danmaku/danmaku_reply/server/http"
	"danmaku/danmaku_reply/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := model.NewConfig()
	s, err := service.NewService(conf)
	if err != nil {
		panic(err)
	}
	defer s.Close()

	httpServer, err := http.StartServer(s, conf)
	if err != nil {
		panic(err)
	}
	defer httpServer.Close()

	err = s.StartJob()
	if err != nil {
		panic(err)
	}
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case s := <-c:
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}
}
