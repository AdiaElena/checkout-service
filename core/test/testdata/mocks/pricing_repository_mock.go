package mocks

import (
	"fmt"

	"github.com/AdiaElena/checkout-service/core/src/model"
	"github.com/AdiaElena/checkout-service/core/src/repository"
)

type PricingRepositoryMock struct {
	Rules map[string]*model.PricingRule
	// CalledWith records SKUs passed to GetPricingRule
	CalledWith []string
	// IsValidCalls records SKUs passed to IsValidSKU
	IsValidCalls []string
}

// Enforce that PricingRepositoryMock implements repository.PricingRepository.
var _ repository.PricingRepository = (*PricingRepositoryMock)(nil)

func NewPricingRepositoryMock() repository.PricingRepository {
	return &PricingRepositoryMock{
		Rules: map[string]*model.PricingRule{
			"A": {SKU: "A", UnitPrice: 50, SpecialQty: 3, SpecialPrice: 130},
			"B": {SKU: "B", UnitPrice: 30, SpecialQty: 2, SpecialPrice: 45},
			"C": {SKU: "C", UnitPrice: 20},
			"D": {SKU: "D", UnitPrice: 15},
		},
	}
}

func (m *PricingRepositoryMock) GetPricingRule(sku string) (*model.PricingRule, error) {
	m.CalledWith = append(m.CalledWith, sku)

	rule, exists := m.Rules[sku]
	if !exists {
		return nil, fmt.Errorf("pricing rule not found for SKU: %q", sku)
	}
	return rule, nil
}

func (m *PricingRepositoryMock) IsValidSKU(sku string) bool {
	m.IsValidCalls = append(m.IsValidCalls, sku)
	_, exists := m.Rules[sku]
	return exists
}
