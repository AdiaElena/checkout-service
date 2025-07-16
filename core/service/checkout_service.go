package service

type CheckoutService interface {
	Scan(sku string) error
	GetTotalPrice() (int, error)
	Reset()
}
