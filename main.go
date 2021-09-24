package main

import (
	"gochi/adapters"
	"gochi/apps/routes"

	"go.uber.org/zap"
)

func main() {
	// configure logger
	log, _ := zap.NewProduction(zap.WithCaller(false))
	defer func() {
		_ = log.Sync()
	}()

	r := routes.CreateRoutes(log)
	adapters.CreateServer(r)
}
