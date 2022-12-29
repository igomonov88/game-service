package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.uber.org/zap"

	"githib.com/igomonov88/game-service/business/contract"
	"githib.com/igomonov88/game-service/business/services/play"
	"githib.com/igomonov88/game-service/internal/validation"
)

type PlayHandler struct {
	logger  *zap.SugaredLogger
	service *play.Service
}

func NewPlayHandler(logger *zap.SugaredLogger, service *play.Service) PlayHandler {
	return PlayHandler{
		logger:  logger,
		service: service,
	}
}

func (ph PlayHandler) RegisterRoutes(router chi.Router) {
	router.Post("/play", ph.play)
	router.Get("/scoreboard", ph.scoreboard)
	router.Delete("/scoreboard", ph.resetScoreboard)
}

func (ph PlayHandler) play(w http.ResponseWriter, r *http.Request) {
	var request contract.PlayRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		render.Render(w, r, contract.ErrorBadRequest)
		return
	}

	if err := validation.Check(request); err != nil {
		ph.logger.Infof("play request validation error: %v, requested choice: %v", err, request.Player)
		render.Render(w, r, contract.ErrorBadRequest)
		return
	}

	response, err := ph.service.Play(r.Context(), request.Player)
	if err != nil {
		render.Render(w, r, contract.ErrorInternal)
		return
	}

	render.JSON(w, r, response)
}

func (ph PlayHandler) scoreboard(w http.ResponseWriter, r *http.Request) {
	results, err := ph.service.GetResults(r.Context())
	if err != nil {
		render.Render(w, r, contract.ErrorInternal)
		return
	}

	render.JSON(w, r, results)
}

func (ph PlayHandler) resetScoreboard(w http.ResponseWriter, r *http.Request) {
	if err := ph.service.ResetScoreboard(r.Context()); err != nil {
		render.Render(w, r, contract.ErrorInternal)
		return
	}

	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, nil)
}
