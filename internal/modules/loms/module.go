package loms

import (
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/ports"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/use_case"
)

type Module struct {
	orderUseCase use_case.OrderUseCase
	stockUseCase use_case.OrderUseCase
}

func InitModule(
	logger ports.Logger,
	orderRepository ports.OrderRepository,
	stockRepository ports.StockRepository,
) (Module, error) {
	orderUseCase, err := use_case.NewOrderUseCase(logger, orderRepository, stockRepository)
	if err != nil {
		return Module{}, err
	}

	logger.Debug("order use case successful initiated")

	stockUseCase, err := use_case.NewOrderUseCase(logger, orderRepository, stockRepository)
	if err != nil {
		return Module{}, err
	}

	logger.Debug("stock use case successful initiated")

	return Module{
		orderUseCase: orderUseCase,
		stockUseCase: stockUseCase,
	}, nil
}
