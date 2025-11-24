package models

type Reservation struct {
	OrderId          uint
	ReservationItems []ReservationItem
}
