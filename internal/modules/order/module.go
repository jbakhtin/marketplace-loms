package order

import (
	"github.com/jbakhtin/marketplace-loms/internal/modules/order/ports/primary"
	"github.com/jbakhtin/marketplace-loms/internal/modules/order/ports/secondary"
	"github.com/jbakhtin/marketplace-loms/internal/modules/order/services"
)

func NewModule(
	logger secondary.Logger,
	repository secondary.OrderRepository,
	service secondary.StockService,
) (primary.UseCase, error) {
	return services.NewUseCase(logger, repository, service)
}
