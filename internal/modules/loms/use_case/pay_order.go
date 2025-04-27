package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/domain"
)

func NewPayOrder(
	logger domain.Logger,
	orderRepository domain.OrderRepository,
	stockService domain.StockService,
) (PayOrder, error) {
	return PayOrder{
		logger:          logger,
		orderRepository: orderRepository,
		stockService:    stockService,
	}, nil
}

type PayOrder struct {
	logger          domain.Logger
	orderRepository domain.OrderRepository
	stockService    domain.StockService
}

func (u *PayOrder) CancelOrder(ctx context.Context, ID domain.OrderID) error {
	//TODO implement me
	panic("implement me")
}
