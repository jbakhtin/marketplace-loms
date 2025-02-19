package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/config"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/server/http"
	"github.com/jbakhtin/marketplace-loms/pkg/closer"
	"github.com/jbakhtin/marketplace-loms/pkg/starter"
	"os/signal"
	"syscall"
)

var err error
var str starter.Starter
var clr closer.Closer
var cfg config.Config
var router *chi.Mux
var server http.Server

func init() {
	starterBuilder := starter.New()
	closerBuilder := closer.New()

	cfg, err = config.NewConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	router, err = http.NewRouter()
	if err != nil {
		fmt.Println(err.Error())
	}

	server, err = http.NewServer(&cfg, router)
	if err != nil {
		fmt.Println(err.Error())
	}
	starterBuilder.Add(server.Start)
	closerBuilder.Add(server.Shutdown)

	str = starterBuilder.Build()
	clr = closerBuilder.Build()
}

func main() {
	osCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	err = str.Start(osCtx)
	if err != nil {
		fmt.Println(err.Error())
	}

	<-osCtx.Done()

	err = clr.Close(osCtx)
	if err != nil {
		fmt.Println(err.Error())
	}
}
