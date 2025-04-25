package rest

import (
	"context"
	"net/http"
)

type Server struct {
	http.Server
}

type Config interface {
	GetServerHTTPAddress() string
}

func NewServer(cfg Config, handler http.Handler) (Server, error) {
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
