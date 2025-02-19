package dto

type OrderItem struct {
	SKU   int32
	count uint16
}

type CreateOrderRequest struct {
	UserID uint64
	Items  []OrderItem
}

type CreateOrderResponse struct {
	OrderID int64
}
