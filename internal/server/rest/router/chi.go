package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jbakhtin/marketplace-loms/pkg/order"
	"github.com/jbakhtin/marketplace-loms/pkg/stock"
)

type Config interface {
}

func NewRouter(
	orderModule order.Module,
	stockModule stock.Module,
) (*chi.Mux, error) {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)

	// Test endpoint
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	err := orderModule.RegisterRoutes(router)
	if err != nil {
		return nil, err
	}

	err = stockModule.RegisterRoutes(router)
	if err != nil {
		return nil, err
	}

	return router, nil
}
