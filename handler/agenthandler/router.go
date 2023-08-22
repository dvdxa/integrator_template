package agenthandler

import (
	"integrator_template/config"
	"integrator_template/domain/agent"
	"integrator_template/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service agent.AgentService
	cfg     *config.Config
	log     *logger.Logger
}

func NewHandler(service agent.AgentService, cfg *config.Config, log *logger.Logger) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
		log:     log,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	newRouter := mux.NewRouter()
	mainRoute := newRouter.PathPrefix("/api").Subrouter()
	routeVer := mainRoute.PathPrefix("/v1").Subrouter()

	routeVer.HandleFunc("/", h.ProcessorHandler).Methods(http.MethodPost)

	return routeVer
}
