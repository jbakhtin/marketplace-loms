package app

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain/models"

	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain"
)

type StockUseCase struct {
	logger          domain.Logger
	stockRepository domain.StockRepository
}

func NewStockUseCase(
	logger domain.Logger,
	stockRepository domain.StockRepository,
) (StockUseCase, error) {
	return StockUseCase{
		logger:          logger,
		stockRepository: stockRepository,
	}, nil
}

func (s *StockUseCase) StockInfo(ctx context.Context, sku int32) (models.StockItem, error) {
	return s.stockRepository.GetBySKU(ctx, sku)
}

func (s *StockUseCase) CancelReservation(ctx context.Context, SKU int32) error {
	return nil
}

// Реализация StockService интерфейса
func (s *StockUseCase) Reserve(ctx context.Context, SKU int32, qty uint16) error {
	return s.stockRepository.Reserve(ctx, SKU, qty)
}

func (s *StockUseCase) ReserveCancel(ctx context.Context, SKU int32, qty uint16) error {
	return s.stockRepository.ReserveCancel(ctx, SKU, qty)
}

func (s *StockUseCase) ReserveRemove(ctx context.Context, SKU int32, qty uint16) error {
	return s.stockRepository.ReserveRemove(ctx, SKU, qty)
}
