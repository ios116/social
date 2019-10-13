package rpc

import (
	"go.uber.org/zap"
	"log"
	"social/cmd"
	"social/internal/grpcserver"
)

func Server() {
	container := cmd.BuildContainer()
	err := container.Invoke(func(serverGRPS *grpcserver.RPCServer, logger *zap.Logger) {
		serverGRPS.Start()
	})
	if err != nil {
		log.Println(err)
	}
}
