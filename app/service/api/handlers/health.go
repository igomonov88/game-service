package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"githib.com/igomonov88/game-service/business/contract"
	"githib.com/igomonov88/game-service/business/services/health"
)

type HealthHandler struct {
	service *health.Service
}

func NewHealthHandler(service *health.Service) HealthHandler {
	return HealthHandler{service: service}
}

func (h HealthHandler) RegisterRoutes(router chi.Router) {
	router.Get("/health", h.health)
}

func (h HealthHandler) health(w http.ResponseWriter, r *http.Request) {
	if err := h.service.HealthCheck(r.Context()); err != nil {
		render.Render(w, r, contract.ErrorInternal)
		return
	}

	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, nil)
}
