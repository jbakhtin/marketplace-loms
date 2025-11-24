package models

import "time"

type ReservationItem struct {
	OrderID   uint
	SKU       int32
	Reserved  uint64
	CreatedAt time.Time
}
