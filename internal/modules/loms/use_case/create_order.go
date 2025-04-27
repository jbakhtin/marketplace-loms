package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/domain"
)

func NewOrderInfo(
	logger domain.Logger,
	orderRepository domain.OrderRepository,
	stockService domain.StockService,
) (OrderInfo, error) {
	return OrderInfo{
		logger:          logger,
		orderRepository: orderRepository,
		stockService:    stockService,
	}, nil
}

type OrderInfo struct {
	logger          domain.Logger
	orderRepository domain.OrderRepository
	stockService    domain.StockService
}

func (o *OrderInfo) Execute(ctx context.Context, ID domain.OrderID) (domain.Order, error) {
	//TODO implement me
	panic("implement me")
}
