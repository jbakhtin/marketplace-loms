package http

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/config"
	"net/http"
)

type Server struct {
	http.Server
}

type Config interface {
	GetServerHTTPAddress() string
}

var _ Config = &config.Config{}

func NewServer(cfg Config, router chi.Router) (Server, error) {
	return Server{
		Server: http.Server{
			Addr:    cfg.GetServerHTTPAddress(),
			Handler: router,
		},
	}, nil
}

func (s *Server) Start(ctx context.Context) (err error) {
	go func() {
		err = s.ListenAndServe()
	}()

	return err
}

// Shutdown корректно завершает работу сервера
func (s *Server) Shutdown(ctx context.Context) err {
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
