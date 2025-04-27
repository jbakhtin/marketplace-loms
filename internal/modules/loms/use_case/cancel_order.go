package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/domain"
)

func NewCancelOrder(
	logger domain.Logger,
	orderRepository domain.OrderRepository,
	stockService domain.StockService,
) (CancelOrder, error) {
	return CancelOrder{
		logger:          logger,
		orderRepository: orderRepository,
		stockService:    stockService,
	}, nil
}

type CancelOrder struct {
	logger          domain.Logger
	orderRepository domain.OrderRepository
	stockService    domain.StockService
}

func (u *CancelOrder) Execute(ctx context.Context, orderItems []domain.OrderItem) (domain.Order, error) {
	//TODO implement me
	panic("implement me")
}
