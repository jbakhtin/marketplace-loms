package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/app"
	"github.com/jbakhtin/marketplace-loms/pkg/stock/domain"
)

type Config interface {
}

type Handler struct {
	cfg     Config
	log     domain.Logger
	useCase app.StockUseCase
}

func NewStockHandler(cfg Config, logger domain.Logger, useCase app.StockUseCase) (Handler, error) {
	return Handler{
		cfg:     cfg,
		log:     logger,
		useCase: useCase,
	}, nil
}

func (h *Handler) RegisterRoutes(r chi.Router) error {
	r.Route("/stocks", func(r chi.Router) {
		r.Get("/info", h.Info)
	})

	return nil
}
