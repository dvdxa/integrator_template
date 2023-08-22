package server

import (
	"context"
	"integrator_template/config"
	"integrator_template/handler/agenthandler"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewServer(cfg *config.Config, handler *agenthandler.Handler) *Server {
	return &Server{httpServer: &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      handler.InitRoutes(),
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}}
}
