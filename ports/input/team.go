package input

type TeamRequest struct {
	Name     string `json:"name"`
	Universe string `json:"universe"`
}

type TeamResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
}
