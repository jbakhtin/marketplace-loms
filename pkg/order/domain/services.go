package domain

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
)

type Logger interface {
	Debug(msg string, fields ...any)
	Info(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, fields ...any)
	Fatal(msg string, fields ...any)
}

// StockService - интерфейс для работы со складом (без зависимости от stock модуля)
type StockService interface {
	Reserve(ctx context.Context, orderId uint, orders []models.OrderItem) error
	ReserveCancel(ctx context.Context, SKU int32, qty uint16) error
	ReserveRemove(ctx context.Context, SKU int32, qty uint16) error
}
