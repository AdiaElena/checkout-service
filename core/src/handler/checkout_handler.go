package handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/AdiaElena/checkout-service/core/src/dto"
	"github.com/AdiaElena/checkout-service/core/src/service"
	"github.com/gin-gonic/gin"
)

// Handler provides a single endpoint for supermarket checkout.
// @title           Supermarket Checkout API
// @version         1.0
// @description     One‑call endpoint to scan items and get total price.
// @host            localhost:8080
// @BasePath        /
type Handler struct {
	checkout func() service.ICheckout
}

// NewHandler constructs a new checkout Handler.
// @summary Create a new checkout handler
func NewHandler(checkout func() service.ICheckout) *Handler {
	return &Handler{checkout: checkout}
}

// Checkout godoc
// @Summary     Scan multiple items and calculate total
// @Description Accepts a list of SKUs, scans them all, and returns the total price.
// @Tags        checkout
// @Accept      json
// @Produce     json
// @Param       request body     dto.CheckoutRequest  true  "Basket SKUs"
// @Success     200     {object} dto.CheckoutResponse "Total price"
// @Failure     400     {object} dto.ErrorResponse    "Bad request (invalid SKU or JSON)"
// @Router      /checkout [post]
func (handler *Handler) Checkout(context *gin.Context) {
	var req dto.CheckoutRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			context.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: valErrs.Error()})
		} else {
			context.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid JSON"})
		}
		return
	}

	// fresh instance per request; no state leakage
	checkoutService := handler.checkout()

	// process all SKUs
	for _, sku := range req.SKUs {
		if err := checkoutService.Scan(sku); err != nil {
			context.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}
	}

	total, err := checkoutService.GetTotalPrice()
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, dto.CheckoutResponse{Total: total})
}
