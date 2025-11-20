package order

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/jbakhtin/marketplace-loms/pkg/order/presentation/http/handler"

	"github.com/jbakhtin/marketplace-loms/pkg/order/app"
	"github.com/jbakhtin/marketplace-loms/pkg/order/domain"
	"github.com/jbakhtin/marketplace-loms/pkg/order/infra/postgres/repositories"
)

type Config interface {
	GetDbDriver() string
	GetDbHost() string
	GetDbPort() string
	GetDbName() string
	GetDbUser() string
	GetDbPassword() string
}

type Module struct {
	orderHandler handler.Handler
	useCase      app.OrderUseCase
	logger       domain.Logger
	cfg          Config
}

func InitModule(
	db *sql.DB,
	logger domain.Logger,
	cfg Config,
	stockService domain.StockService,
) (Module, error) {
	orderRepository, err := repositories.NewOrderRepository(db)
	if err != nil {
		return Module{}, err
	}

	orderUseCase, err := app.NewOrderUseCase(logger, orderRepository, stockService)
	if err != nil {
		return Module{}, err
	}

	return Module{
		useCase: orderUseCase,
		logger:  logger,
		cfg:     cfg,
	}, nil
}

func (m *Module) RegisterRoutes(r chi.Router) (err error) {
	m.orderHandler, err = handler.NewOrderHandler(m.cfg, m.logger, m.useCase)
	if err != nil {
		return nil
	}

	return m.orderHandler.RegisterRoutes(r)
}
