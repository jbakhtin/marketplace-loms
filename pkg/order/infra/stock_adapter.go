package infra

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
	stockModule "github.com/jbakhtin/marketplace-loms/pkg/stock/app"
	stockmodels "github.com/jbakhtin/marketplace-loms/pkg/stock/domain/models"
)

type StockAdapter struct {
	useCase stockModule.UseCase
}

func NewStockAdapter(useCase stockModule.UseCase) (StockAdapter, error) {
	return StockAdapter{
		useCase: useCase,
	}, nil
}

func (a *StockAdapter) Reserve(ctx context.Context, orderId uint, items []models.OrderItem) error {
	// 1. Преобразуем OrderItem → ReservationRequest
	req := make([]stockmodels.ReservationItem, 0, len(items))

	for _, it := range items {
		req = append(req, stockmodels.ReservationItem{
			OrderID:  orderId,
			SKU:      it.SKU,
			Reserved: uint64(it.Count),
		})
	}

	// 2. Вызываем модуль Stock
	_, err := a.useCase.Reserve(ctx, orderId, req)
	return err
}

func (a *StockAdapter) ReserveCancel(ctx context.Context, SKU int32, qty uint16) error {
	//TODO implement me
	panic("implement me")
}

func (a *StockAdapter) ReserveRemove(ctx context.Context, SKU int32, qty uint16) error {
	//TODO implement me
	panic("implement me")
}
