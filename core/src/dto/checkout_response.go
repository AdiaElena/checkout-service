package dto

// CheckoutResponse is returned by the /checkout endpoint.
// swagger:model CheckoutResponse
type CheckoutResponse struct {
	// Total price for all scanned items
	// required: true
	Total int `json:"total"`
}
