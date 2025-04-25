package main

import (
	"context"
	"fmt"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/config"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/logger/zap"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/server/rest"
	"github.com/jbakhtin/marketplace-loms/internal/infrastucture/server/rest/router/chi"
	"github.com/jbakhtin/marketplace-loms/pkg/closer"
	"github.com/jbakhtin/marketplace-loms/pkg/starter"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

var err error
var lgr zap.Logger
var str starter.Starter
var clr closer.Closer
var cfg config.Config
var handler http.Handler
var restServer rest.Server

func init() {
	cfg, err = config.NewConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	lgr, err = zap.NewLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}

	starterBuilder := starter.New()
	closerBuilder := closer.New()

	handler, err = router.NewRouter(&cfg, lgr)
	if err != nil {
		log.Fatal(err)
	}

	restServer, err = rest.NewServer(&cfg, handler)
	if err != nil {
		log.Fatal(err)
	}
	starterBuilder.Add(restServer.Start)
	closerBuilder.Add(restServer.Shutdown)

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
