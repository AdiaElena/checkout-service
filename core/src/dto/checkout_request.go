package dto

// CheckoutRequest is the JSON body for the /checkout endpoint.
// swagger:model CheckoutRequest
type CheckoutRequest struct {
	// List of SKUs to scan at checkout (each must be a single uppercase letter)
	// required: true
	SKUs []string `json:"skus" binding:"required,dive,required,uppercase,len=1"`
}
