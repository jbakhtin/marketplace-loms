package domain

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain/models"
)

type StockRepository interface {
	Reserve(ctx context.Context, orderId uint, reservationItems []models.ReservationItem) (models.Reservation, error)
	ReserveCancel(ctx context.Context, orderId uint) (models.Reservation, error)
	ReserveRemove(ctx context.Context, orderId uint) (models.Reservation, error)
	GetStockItemBySKU(ctx context.Context, SKU int32) (models.StockItem, error)
	GetReservationByOrderID(ctx context.Context, orderId uint) (models.Reservation, error)
}
