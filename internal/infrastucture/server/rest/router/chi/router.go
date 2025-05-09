package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/server/rest/handler/order"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/server/rest/handler/stock"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/ports"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/use_case"
)

type Config interface {
}

func NewRouter(
	cfg Config,
	logger ports.Logger,
	orderUseCase use_case.OrderUseCase,
	stockUseCase use_case.StockUseCase,
) (*chi.Mux, error) {
	orderHandler, err := order.NewOrderHandler(cfg, logger, orderUseCase)
	if err != nil {
		return nil, err
	}

	stockHandler, err := stock.NewStockHandler(cfg, logger, stockUseCase)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)

	router.Route("/orders", func(r chi.Router) {
		r.Post("/create", orderHandler.Create)

		r.Route("/{OrderID}", func(r chi.Router) {
			r.Get("/info", orderHandler.Info)
			r.Put("/pay", orderHandler.Pay)
			r.Put("/cancel", orderHandler.Cancel)
		})
	})

	router.Route("/stocks", func(r chi.Router) {
		r.Route("/{StockID}", func(r chi.Router) {
			r.Get("/info", stockHandler.Info)
		})
	})

	return router, nil
}
