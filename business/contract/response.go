package contract

type ChoiceResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PlayResponse struct {
	Results  string `json:"results"`
	Player   int    `json:"player"`
	Computer int    `json:"computer"`
}
