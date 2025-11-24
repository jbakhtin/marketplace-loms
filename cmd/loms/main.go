package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jbakhtin/marketplace-loms/internal/storage/postgres"
	"github.com/jbakhtin/marketplace-loms/pkg/order/infra"
	"log"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jbakhtin/marketplace-loms/internal/config"
	"github.com/jbakhtin/marketplace-loms/internal/logger/zap"
	"github.com/jbakhtin/marketplace-loms/internal/server/rest"
	"github.com/jbakhtin/marketplace-loms/pkg/closer"
	"github.com/jbakhtin/marketplace-loms/pkg/order"
	"github.com/jbakhtin/marketplace-loms/pkg/starter"
	"github.com/jbakhtin/marketplace-loms/pkg/stock"
	"github.com/joho/godotenv"
)

var err error
var logger zap.Logger
var str starter.Starter
var clr closer.Closer
var cfg config.Config
var restServer rest.Server
var db *sql.DB

func init() {
	_ = godotenv.Load()

	cfg, err = config.NewConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	logger, err = zap.NewLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}

	starterBuilder := starter.New()
	closerBuilder := closer.New()

	db, err = postgres.NewSQLClient(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация модулей (каждый получает только db)
	stockModule, err := stock.InitModule(db, logger, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	stockUseCases := stockModule.GetUseCases()

	stockAdapter, err := infra.NewStockAdapter(stockUseCases)
	if err != nil {
		log.Fatal(err)
	}

	orderModule, err := order.InitModule(db, logger, &cfg, &stockAdapter)
	if err != nil {
		log.Fatal(err)
	}

	restServer, err = rest.NewWebServer(&cfg, orderModule, stockModule)
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

	if db != nil {
		db.Close()
	}
}
