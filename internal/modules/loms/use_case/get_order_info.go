package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/domain"
)

func NewGetOrderInfo(
	logger domain.Logger,
	orderRepository domain.OrderRepository,
	stockService domain.StockService,
) (GetOrderInfo, error) {
	return GetOrderInfo{
		logger:          logger,
		orderRepository: orderRepository,
		stockService:    stockService,
	}, nil
}

type GetOrderInfo struct {
	logger          domain.Logger
	orderRepository domain.OrderRepository
	stockService    domain.StockService
}

func (g *GetOrderInfo) Execute(ctx context.Context, orderID domain.OrderID) (domain.Order, error) {
	//TODO implement me
	panic("implement me")
}
