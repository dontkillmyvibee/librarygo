package corehttpresponse

type ErrorResponse struct {
	Error   string `json:"error"   example:"invalid argument"`
	Message string `json:"message" example:"short human-readable message"`
}
