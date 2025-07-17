package dto

// ErrorResponse is returned when an endpoint fails.
// swagger:model ErrorResponse
type ErrorResponse struct {
	// Error message
	// required: true
	Error string `json:"error"`
}
