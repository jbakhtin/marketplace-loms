package use_case

import (
	"context"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/domain"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/ports"
)

type StockUseCase struct {
	logger          ports.Logger
	stockRepository ports.StockRepository
}

func NewStockUseCase(
	logger ports.Logger,
	stockRepository ports.StockRepository,
) (OrderUseCase, error) {
	return OrderUseCase{
		logger:          logger,
		stockRepository: stockRepository,
	}, nil
}

func (s *StockUseCase) StockInfo(ctx context.Context, SKU domain.StockItemSKU) (domain.StockItem, error) {
	return domain.StockItem{}, nil
}

func (s *StockUseCase) CancelReservation(ctx context.Context, SKU domain.StockItemSKU) error {
	return nil
}
