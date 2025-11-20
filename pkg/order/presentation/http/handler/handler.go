package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/jbakhtin/marketplace-loms/pkg/order/app"
	"github.com/jbakhtin/marketplace-loms/pkg/order/domain"
)

type Config interface {
}

type Handler struct {
	cfg     Config
	logger  domain.Logger
	useCase app.OrderUseCase
}

func NewOrderHandler(cfg Config, logger domain.Logger, useCase app.OrderUseCase) (Handler, error) {
	return Handler{
		cfg:     cfg,
		logger:  logger,
		useCase: useCase,
	}, nil
}

func (h *Handler) RegisterRoutes(r chi.Router) error {
	orderHandler, err := NewOrderHandler(h.cfg, h.logger, h.useCase)
	if err != nil {
		return err
	}

	r.Route("/orders", func(r chi.Router) {
		r.Post("/create", orderHandler.Create)

		r.Route("/{OrderID}", func(r chi.Router) {
			r.Get("/info", orderHandler.Info)
			r.Put("/pay", orderHandler.Pay)
			r.Put("/cancel", orderHandler.Cancel)
		})
	})

	return nil
}
