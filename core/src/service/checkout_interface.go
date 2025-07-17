package service

type ICheckout interface {
	Scan(SKU string) error
	GetTotalPrice() (int, error)
}
