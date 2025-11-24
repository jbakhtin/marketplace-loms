package rest

import (
	"context"
	"net/http"

	"github.com/jbakhtin/marketplace-loms/internal/server/rest/router"
	"github.com/jbakhtin/marketplace-loms/pkg/order"
	"github.com/jbakhtin/marketplace-loms/pkg/stock"
)

type Server struct {
	http.Server
}

type Config interface {
	GetServerHTTPAddress() string
}

func NewWebServer(
	cfg Config,
	orderModule order.Module,
	stockModule stock.Module,
) (Server, error) {
	handler, err := router.NewRouter(orderModule, stockModule)
	if err != nil {
		return Server{}, err
	}

	return Server{
		Server: http.Server{
			Addr:    cfg.GetServerHTTPAddress(),
			Handler: handler,
		},
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
