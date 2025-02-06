package models

type OrderSKU int32
type OrderCount uint16

type OrderItem struct {
	SKU   OrderSKU
	count OrderCount
}
