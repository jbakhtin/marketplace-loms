package stock

import (
	"database/sql"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/presentation/http/handler"

	"github.com/go-chi/chi/v5"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/app"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/infra/postgres/repositories"
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
	stockHandler handler.Handler
	stockUseCase app.StockUseCase
	logger       domain.Logger
	cfg          Config
}

func InitModule(
	db *sql.DB,
	logger domain.Logger,
	cfg Config,
) (Module, error) {
	stockRepository, err := repositories.NewStockRepository(db)
	if err != nil {
		return Module{}, err
	}

	stockUseCase, err := app.NewStockUseCase(logger, &stockRepository)
	if err != nil {
		return Module{}, err
	}

	return Module{
		stockUseCase: stockUseCase,
		logger:       logger,
		cfg:          cfg,
	}, nil
}

func (m *Module) RegisterRoutes(r chi.Router) (err error) {
	m.stockHandler, err = handler.NewStockHandler(m.cfg, m.logger, m.stockUseCase)
	if err != nil {
		return nil
	}

	return m.stockHandler.RegisterRoutes(r)
}

func (m *Module) GetUseCases() app.StockUseCase {
	return m.stockUseCase
}
