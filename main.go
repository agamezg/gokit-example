package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
	"kit-example/account"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const datasource = "postgres://postgres:postgres@localhost:5432/gokit_example?sslmode=disable"

func main(){
	var httpAdrr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller)
	}

	level.Info(logger).Log("msg", "service started")
	defer level .Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", datasource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	flag.Parse()
	ctx := context.Background()
	var service account.Service
	{
		repository := account.NewRepo(db, logger)
		service = account.NewService(repository, logger)
	}

	errors := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errors <- fmt.Errorf("%s", <-c)
	}()

	endpoints := account.MakeEndpoints(service)

	go func() {
		fmt.Println("listening on port", *httpAdrr)
		handler := account.NewHttpServer(ctx, endpoints)
		errors <- http.ListenAndServe(*httpAdrr, handler)
	}()

	level.Error(logger).Log("exit", <-errors)
}

