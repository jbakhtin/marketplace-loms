package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/domain"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/ports"
)

type OrderUseCase struct {
	logger          ports.Logger
	orderRepository ports.OrderRepository
	stockRepository ports.StockRepository
}

func NewOrderUseCase(
	logger ports.Logger,
	orderRepository ports.OrderRepository,
	stockRepository ports.StockRepository,
) (OrderUseCase, error) {
	return OrderUseCase{
		logger:          logger,
		orderRepository: orderRepository,
		stockRepository: stockRepository,
	}, nil
}

func (o *OrderUseCase) CancelOrder(ctx context.Context, ID domain.OrderID) error {
	return nil
}

func (o *OrderUseCase) GetOrderInfo(ctx context.Context, orderID domain.OrderID) (domain.Order, error) {
	return domain.Order{}, nil
}

func (o *OrderUseCase) CreateOrder(ctx context.Context, ID domain.OrderID) error {
	return nil
}

func (o *OrderUseCase) PayOrder(ctx context.Context, ID domain.OrderID) error {
	return nil
}
