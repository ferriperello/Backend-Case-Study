package main

import (
	"Challenge/internal/repos"
	"Challenge/internal/server"
	"context"
	"github.com/kelseyhightower/envconfig"
	"net/http"
	"os"
	"os/signal"
)

var conf config

func main() {

	ctx, cancelContext := context.WithCancel(context.Background())
	defer cancelContext()

	// open logger

	// load config
	if err := envconfig.Process("config", &conf); err != nil {
		panic(err)
	}

	// init persistence
	mem := repos.NewMemStore()
	ctx = context.WithValue(ctx, "mem", mem)

	// start http server service
	r := server.NewRouter()

	Httpserver := &http.Server{Addr: ":8080", Handler: r}
	go func() {
		if err := Httpserver.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// log the error
		}
	}()

	// gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	_ = Httpserver.Shutdown(ctx)

}
