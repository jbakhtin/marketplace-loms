package services

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/order/models"
	"github.com/jbakhtin/marketplace-loms/internal/modules/order/ports/primary"
	"github.com/jbakhtin/marketplace-loms/internal/modules/order/ports/secondary"
)

var _ primary.UseCase = &UseCase{}

func NewUseCase(
	logger secondary.Logger,
	orderRepository secondary.OrderRepository,
	stockService secondary.StockService,
) (primary.UseCase, error) {
	return &UseCase{
		logger:          logger,
		orderRepository: orderRepository,
		stockService:    stockService,
	}, nil
}

type UseCase struct {
	logger          secondary.Logger
	orderRepository secondary.OrderRepository
	stockService    secondary.StockService
}

func (u *UseCase) CreateOrder(ctx context.Context, orderItems []models.OrderItem) (models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UseCase) OrderInfo(ctx context.Context, orderID models.OrderID) (models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UseCase) OrderPay(ctx context.Context, ID models.OrderID) error {
	//TODO implement me
	panic("implement me")
}

func (u *UseCase) CancelOrder(ctx context.Context, ID models.OrderID) error {
	//TODO implement me
	panic("implement me")
}
