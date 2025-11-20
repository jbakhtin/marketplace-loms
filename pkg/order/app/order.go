package app

import (
	"context"

	"github.com/jbakhtin/marketplace-loms/pkg/order/domain"
	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/entity"
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
func (o *OrderUseCase) CreateOrder(ctx context.Context, order entity.Order) error {
	// заказ получает статус "new"
	// резервирует нужное количество единиц товара
	// если удалось зарезервировать стоки, заказ получает статус "awaiting payment"
	// если не удалось зарезервировать стоки, заказ получает статус "failed"
	order, err := o.orderRepository.Create(ctx, order)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		err = o.stockService.Reserve(ctx, item.SKU, item.Count)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *OrderUseCase) CancelOrder(ctx context.Context, ID int64) error {
	return nil
}

func (o *OrderUseCase) GetOrderInfo(ctx context.Context, orderID int64) (entity.Order, error) {
	return entity.Order{}, nil
}

func (o *OrderUseCase) PayOrder(ctx context.Context, ID int64) error {
	return nil
}
