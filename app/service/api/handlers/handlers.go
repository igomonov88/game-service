package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"

	"githib.com/igomonov88/game-service/business/services/choice"
	"githib.com/igomonov88/game-service/business/services/health"
	"githib.com/igomonov88/game-service/business/services/play"
	"githib.com/igomonov88/game-service/internal/middlewaries"
)

func Handler(logger *zap.SugaredLogger, choiceService *choice.Service, playService *play.Service, healthService *health.Service) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middlewaries.RequestID)

	router.Route("/", func(r chi.Router) {
		choiceHandler := NewChoiceHandler(logger, choiceService)
		playHandler := NewPlayHandler(logger, playService)
		healthHandler := NewHealthHandler(healthService)
		choiceHandler.RegisterRoutes(r)
		playHandler.RegisterRoutes(r)
		healthHandler.RegisterRoutes(r)
	})

	return router
}
