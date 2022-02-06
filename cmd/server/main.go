package main

import (
	"Challenge/internal/server"
	"Challenge/package/jsonApi"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	// start http server service
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("v1", func(r chi.Router) {
		r.Post("complex_report", jsonApi.JSON(server.CreateReport))
		r.Get("status/{job_id}", jsonApi.JSON(server.JobStatus))
	})

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
