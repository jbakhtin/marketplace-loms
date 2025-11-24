package app

import (
	"context"

	"github.com/jbakhtin/marketplace-loms/pkg/order/domain/models"
	"github.com/pkg/errors"

	"github.com/jbakhtin/marketplace-loms/pkg/order/domain"
)

// CreateOrderItemDTO - DTO для создания заказа (не доменная модель)
type CreateOrderItemDTO struct {
	SKU      int32
	Quantity uint16
}

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

func (o *OrderUseCase) CreateOrder(ctx context.Context, userID uint64, items []CreateOrderItemDTO) (int64, error) {
	// Бизнес-валидация
	if len(items) == 0 {
		return 0, errors.New("order must contain at least one item")
	}

	// Преобразуем DTO в доменные модели
	orderItems := make([]models.OrderItem, len(items))
	for i, item := range items {
		if item.SKU <= 0 {
			return 0, errors.Errorf("invalid SKU: %d", item.SKU)
		}
		if item.Quantity == 0 {
			return 0, errors.Errorf("quantity must be greater than 0 for SKU: %d", item.SKU)
		}
		orderItems[i] = models.OrderItem{
			SKU:   item.SKU,
			Count: item.Quantity,
		}
	}

	// Создаем заказ
	order, err := o.orderRepository.Create(ctx, userID, orderItems)
	if err != nil {
		return 0, errors.Wrap(err, "create order")
	}

	// Резервируем товары
	if err := o.stockService.Reserve(ctx, uint(order.ID), order.Items); err != nil {
		// Если резервирование не удалось, пытаемся отменить заказ
		// TODO: добавить транзакции для атомарности
		o.logger.Error("failed to reserve stocks, order created but not reserved",
			"order_id", order.ID, "error", err)
		return 0, errors.Wrap(err, "reserve stocks")
	}

	o.logger.Info("order created successfully", "order_id", order.ID, "user_id", userID)
	return int64(order.ID), nil
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
