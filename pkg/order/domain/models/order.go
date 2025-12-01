package models

const (
	New             string = "new"
	AwaitingPayment string = "awaiting_payment"
	Failed          string = "failed"
	Payed           string = "payed"
	Cancelled       string = "cancelled"
)

type Order struct {
	ID     uint64
	UserID int64
	Status string
	Items  []OrderItem
}
