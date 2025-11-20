package domain

import "context"

type Logger interface {
	Debug(msg string, fields ...any)
	Info(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, fields ...any)
	Fatal(msg string, fields ...any)
}

// StockService - интерфейс для работы со складом (без зависимости от stock модуля)
type StockService interface {
	Reserve(ctx context.Context, SKU int32, qty uint16) error
	ReserveCancel(ctx context.Context, SKU int32, qty uint16) error
	ReserveRemove(ctx context.Context, SKU int32, qty uint16) error
}
