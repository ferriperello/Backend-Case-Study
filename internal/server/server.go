package server

import (
	"Challenge/package/jsonApi"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("v1", func(r chi.Router) {
		r.Post("complex_report", jsonApi.JSON(CreateReport))
		r.Get("status/{job_id}", jsonApi.JSON(JobStatus))
	})

	return r
}

func JobStatus(ctx context.Context, r *http.Request) (interface{}, error) {

	return nil, nil
}

func CreateReport(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}