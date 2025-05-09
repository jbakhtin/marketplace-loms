package stock

import (
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/ports"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/use_case"
)

type Config interface {
}

type Handler struct {
	cfg     Config
	log     ports.Logger
	useCase use_case.StockUseCase
}

func NewStockHandler(cfg Config, logger ports.Logger, useCase use_case.StockUseCase) (Handler, error) {
	return Handler{
		cfg:     cfg,
		log:     logger,
		useCase: useCase,
	}, nil
}
