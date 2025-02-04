package models

type OrderStatus int

type OrderID int64

const (
	_               = iota
	New OrderStatus = 1
	AwaitingPayment
	Failed
	Payed
	Cancelled
)

type Order struct {
	OrderID
	status OrderStatus
}
