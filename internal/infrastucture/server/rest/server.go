package rest

import (
	"context"
	router "github.com/jbakhtin/marketplace-loms/internal/infrastucture/server/rest/router/chi"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms"
	"github.com/jbakhtin/marketplace-loms/internal/modules/loms/ports"
	"net/http"
)

type Server struct {
	http.Server
}

type Config interface {
	GetServerHTTPAddress() string
}

func NewWebServer(
	cfg Config,
	logger ports.Logger,
	module loms.Module,
) (Server, error) {
	handler, err := router.NewRouter(&cfg, logger, orderUseCase, stockUseCase)
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

func (s *Server) Start(ctx context.Context) (err error) {
	go func() {
		err = s.ListenAndServe()
	}()

	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
