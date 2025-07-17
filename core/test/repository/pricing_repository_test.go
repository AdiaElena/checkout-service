package repository_test

import (
	"testing"

	"github.com/AdiaElena/checkout-service/core/src/repository"
	"github.com/stretchr/testify/assert"
)

func TestPricingRepository_GetPricingRule_and_IsValidSKU(t *testing.T) {
	repo := repository.NewPricingRepository()

	tests := []struct {
		sku         string
		expectValid bool
		expectErr   bool
		expectPrice int
	}{
		{"A", true, false, 50},
		{"B", true, false, 30},
		{"C", true, false, 20},
		{"D", true, false, 15},
		{"X", false, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.sku, func(t *testing.T) {
			rule, err := repo.GetPricingRule(tt.sku)
			if tt.expectErr {
				// Unknown SKUs should produce an error and no rule
				assert.Error(t, err)
				assert.Nil(t, rule)
			} else {
				// Known SKUs should return no error and a non‑nil rule
				assert.NoError(t, err)
				assert.NotNil(t, rule)
				assert.Equal(t, tt.expectPrice, rule.UnitPrice)
			}
			// IsValidSKU should agree with whether we expected an error
			assert.Equal(t, tt.expectValid, repo.IsValidSKU(tt.sku))
		})
	}
}
