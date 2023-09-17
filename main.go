package main

import (
	"integrator_template/adapter/vend"
	"integrator_template/config"
	"integrator_template/domain/agent"
	"integrator_template/handler/agenthandler"
	"integrator_template/pkg/logger"
	"integrator_template/server"
	"log"
)

func main() {
	cfg, err := config.InitConfigs()
	if err != nil {
		log.Fatalf("failed to init configs %v", err)
	}
	log := logger.GetLogger()
	adapters := vend.NewAdapter(&cfg.Adapter, log)
	services := agent.NewService(adapters)
	handlers := agenthandler.NewHandler(services, cfg, log)
	srv := server.NewServer(cfg, handlers)
	srv.Run()
	log.Println("server is running...")
}
