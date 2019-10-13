package web

import (
	"go.uber.org/zap"
	"log"
	"social/cmd"
	"social/internal/webserver"
	"time"
)

func Server() {
	time.Sleep(5 * time.Second)
	container := cmd.BuildContainer()
	err := container.Invoke(func(serverWeb *webserver.HttpServer, logger *zap.Logger) {
		serverWeb.Run()
	})
	if err != nil {
		log.Println(err)
	}
}
