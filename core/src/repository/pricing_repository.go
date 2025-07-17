package repository

import (
	"fmt"

	"github.com/AdiaElena/checkout-service/core/src/model"
)

// PricingRepository defines how to retrieve pricing rules
// and verify SKU validity.
type PricingRepository interface {
	// GetPricingRule looks up the rule for sku and returns an error if not found.
	GetPricingRule(sku string) (*model.PricingRule, error)
	// IsValidSKU returns true if sku exists in the pricing rules map.
	IsValidSKU(sku string) bool
}

type pricingRepo struct {
	rules map[string]*model.PricingRule
}

// NewPricingRepository constructs a PricingRepository pre‑populated
// with the default in‑memory pricing rules (A, B, C, D).
func NewPricingRepository() PricingRepository {
	return &pricingRepo{
		rules: map[string]*model.PricingRule{
			"A": {SKU: "A", UnitPrice: 50, SpecialQty: 3, SpecialPrice: 130},
			"B": {SKU: "B", UnitPrice: 30, SpecialQty: 2, SpecialPrice: 45},
			"C": {SKU: "C", UnitPrice: 20},
			"D": {SKU: "D", UnitPrice: 15},
		},
	}
}

func (p *pricingRepo) GetPricingRule(sku string) (*model.PricingRule, error) {
	rule, exists := p.rules[sku]
	if !exists {
		return nil, fmt.Errorf("pricing rule not found for SKU: %q", sku)
	}
	return rule, nil
}

func (p *pricingRepo) IsValidSKU(sku string) bool {
	_, exists := p.rules[sku]
	return exists
}
