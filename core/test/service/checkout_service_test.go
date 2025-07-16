package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AdiaElena/checkout-service/core/src/service"
	"github.com/AdiaElena/checkout-service/core/test/testdata/mocks"
)

func TestCheckoutService(t *testing.T) {
	tests := []struct {
		name            string
		scannedSKUs     []string
		expectScanError bool
		expectedTotal   int
	}{
		{
			name:          "single A",
			scannedSKUs:   []string{"A"},
			expectedTotal: 50,
		},
		{
			name:          "three A with special pricing",
			scannedSKUs:   []string{"A", "A", "A"},
			expectedTotal: 130,
		},
		{
			name:          "two B with special pricing",
			scannedSKUs:   []string{"B", "B"},
			expectedTotal: 45,
		},
		{
			name:          "mixed A, B, B, C",
			scannedSKUs:   []string{"A", "B", "B", "C"},
			expectedTotal: 50 + 45 + 20,
		},
		{
			name:            "invalid SKU midsequence",
			scannedSKUs:     []string{"A", "X", "B"},
			expectScanError: true,
		},
		{
			name:          "empty cart",
			scannedSKUs:   []string{},
			expectedTotal: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// each test gets a fresh service instance
			mockRepo := mocks.NewPricingRepositoryMock()
			checkout := service.NewCheckoutService(mockRepo)

			var scanErr error
			for _, sku := range tc.scannedSKUs {
				scanErr = checkout.Scan(sku)
				if scanErr != nil {
					break
				}
			}

			if tc.expectScanError {
				assert.Error(t, scanErr, "expected a scan error for invalid SKU")
				return
			}
			assert.NoError(t, scanErr, "unexpected scan error")

			total, err := checkout.GetTotalPrice()
			assert.NoError(t, err, "unexpected error computing total")
			assert.Equal(t, tc.expectedTotal, total, "total price mismatch")
		})
	}
}
