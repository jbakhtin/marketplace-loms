package models

type OrderID int64
type UserID int64
type OrderStatus string

const (
	New             OrderStatus = "new"
	AwaitingPayment OrderStatus = "awaiting_payment"
	Failed          OrderStatus = "failed"
	Payed           OrderStatus = "payed"
	Cancelled       OrderStatus = "cancelled"
)

type Order struct {
	ID     OrderID
	userID UserID
	status OrderStatus
	items  []OrderItem
}
