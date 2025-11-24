package app

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
	"github.com/pkg/errors"

	"github.com/jbakhtin/marketplace-loms/pkg/order/domain"
)

type OrderUseCase struct {
	logger          domain.Logger
	orderRepository domain.OrderRepository
	stockService    domain.StockService
}

func NewOrderUseCase(
	logger domain.Logger,
	orderRepository domain.OrderRepository,
	stockService domain.StockService,
) (OrderUseCase, error) {
	return OrderUseCase{
		logger:          logger,
		orderRepository: orderRepository,
		stockService:    stockService,
	}, nil
}
func (o *OrderUseCase) CreateOrder(ctx context.Context, items []models.OrderItem) error {
	order, err := o.orderRepository.Create(ctx, items)
	if err != nil {
		return errors.Wrap(err, "create order")
	}

	err = o.stockService.Reserve(ctx, order.ID, order.Items)
	if err != nil {
		return errors.Wrap(err, "reserve stocks")
	}

	return nil
}

func (o *OrderUseCase) CancelOrder(ctx context.Context, ID int64) error {
	return nil
}

func (o *OrderUseCase) GetOrderInfo(ctx context.Context, orderID int64) (models.Order, error) {
	return models.Order{}, nil
}

func (o *OrderUseCase) PayOrder(ctx context.Context, ID int64) error {
	return nil
}
