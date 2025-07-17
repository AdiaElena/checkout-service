package service

import (
	"fmt"

	"github.com/AdiaElena/checkout-service/core/src/repository"
)

type checkoutService struct {
	pricingRepo repository.PricingRepository
	items       map[string]int
}

var _ ICheckout = (*checkoutService)(nil)

// NewCheckoutService returns an ICheckout configured with repo’s pricing rules.
func NewCheckoutService(repo repository.PricingRepository) ICheckout {
	return &checkoutService{
		pricingRepo: repo,
		items:       make(map[string]int),
	}
}

func (c *checkoutService) Scan(sku string) error {
	if !c.pricingRepo.IsValidSKU(sku) {
		return fmt.Errorf("invalid SKU: %q", sku)
	}
	c.items[sku]++
	return nil
}

func (c *checkoutService) GetTotalPrice() (int, error) {
	total := 0
	for sku, qty := range c.items {
		rule, err := c.pricingRepo.GetPricingRule(sku)
		if err != nil {
			return 0, err
		}
		if rule.SpecialQty > 0 {
			total += (qty/rule.SpecialQty)*rule.SpecialPrice + (qty%rule.SpecialQty)*rule.UnitPrice
		} else {
			total += qty * rule.UnitPrice
		}
	}
	return total, nil
}
