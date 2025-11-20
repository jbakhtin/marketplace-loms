package entity

const (
	New             string = "new"
	AwaitingPayment string = "awaiting_payment"
	Failed          string = "failed"
	Payed           string = "payed"
	Cancelled       string = "cancelled"
)

type Order struct {
	ID     int64
	UserID int64
	Status string
	Items  []OrderItem
}
