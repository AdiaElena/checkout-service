package mocks

import (
	"github.com/AdiaElena/checkout-service/core/src/service"
)

// CheckoutServiceMock is a testify‑style mock for service.ICheckout
type CheckoutServiceMock struct {
	// ScanFunc is called when CheckoutServiceMock.Scan is invoked
	ScanFunc func(sku string) error
	// GetTotalPriceFunc is called when CheckoutServiceMock.GetTotalPrice is invoked
	GetTotalPriceFunc func() (int, error)
}

// Enforce interface
var _ service.ICheckout = (*CheckoutServiceMock)(nil)

// NewCheckoutServiceMock returns a mock with no‑ops by default
func NewCheckoutServiceMock() *CheckoutServiceMock {
	return &CheckoutServiceMock{
		ScanFunc:          func(string) error { return nil },
		GetTotalPriceFunc: func() (int, error) { return 0, nil },
	}
}

func (m *CheckoutServiceMock) Scan(sku string) error {
	return m.ScanFunc(sku)
}

func (m *CheckoutServiceMock) GetTotalPrice() (int, error) {
	return m.GetTotalPriceFunc()
}
