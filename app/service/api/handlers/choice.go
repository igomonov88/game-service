package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.uber.org/zap"

	"githib.com/igomonov88/game-service/business/contract"
	"githib.com/igomonov88/game-service/business/services/choice"
)

type ChoiceHandler struct {
	logger  *zap.SugaredLogger
	service *choice.Service
}

func NewChoiceHandler(logger *zap.SugaredLogger, service *choice.Service) ChoiceHandler {
	return ChoiceHandler{
		logger:  logger,
		service: service,
	}
}

func (ch ChoiceHandler) RegisterRoutes(router chi.Router) {
	router.Get("/choice", ch.getChoice)
	router.Get("/choices", ch.getChoices)
}

func (ch ChoiceHandler) getChoice(w http.ResponseWriter, r *http.Request) {
	resp, err := ch.service.GetChoice(r.Context())
	if err != nil {
		render.Render(w, r, contract.ErrorInternal)
		return
	}

	render.JSON(w, r, resp)
}

func (ch ChoiceHandler) getChoices(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, ch.service.GetChoices(r.Context()))
}
