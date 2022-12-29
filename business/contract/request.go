package contract

type PlayRequest struct {
	Player int `json:"player" validate:"gte=1,lte=5"`
}
