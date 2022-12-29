package game

import "go.uber.org/zap"

const (
	rock = iota + 1
	paper
	scissors
	lizard
	spock
)

type Service struct {
	choices map[int]string
	logger  *zap.SugaredLogger
}

func NewService(logger *zap.SugaredLogger) *Service {
	return &Service{
		choices: map[int]string{
			rock:     "rock",
			paper:    "paper",
			scissors: "scissors",
			lizard:   "lizard",
			spock:    "spock",
		},
		logger: logger,
	}
}
