package app

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain/models"
	"github.com/pkg/errors"

	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain"
)

type UseCase struct {
	logger     domain.Logger
	repository domain.StockRepository
}

func NewUseCase(
	logger domain.Logger,
	repository domain.StockRepository,
) (UseCase, error) {
	return UseCase{
		logger:     logger,
		repository: repository,
	}, nil
}

func (uc *UseCase) StockInfo(ctx context.Context, sku int32) (models.StockItem, error) {
	return uc.repository.GetStockItemBySKU(ctx, sku)
}

func (uc *UseCase) Reserve(
	ctx context.Context,
	orderId uint,
	reservationItems []models.ReservationItem,
) (models.Reservation, error) {
	reservation, err := uc.repository.GetReservationByOrderID(ctx, orderId)
	if err != nil {
		if !errors.Is(err, domain.NotFoundError) {
			return models.Reservation{}, err
		}
	} else {
		return reservation, nil
	}

	return uc.repository.Reserve(ctx, orderId, reservationItems)
}

func (uc *UseCase) ReserveCancel(ctx context.Context, orderId uint) (models.Reservation, error) {
	return uc.repository.ReserveCancel(ctx, orderId)
}

func (uc *UseCase) ReserveRemove(ctx context.Context, orderId uint) (models.Reservation, error) {
	return uc.repository.ReserveRemove(ctx, orderId)
}
