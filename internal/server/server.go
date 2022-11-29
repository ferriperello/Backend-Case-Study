package server

import (
	"Challenge/internal/repos"
	use_cases "Challenge/internal/use-cases"
	"Challenge/package/jsonApi"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type router struct {
	Rout chi.Router
	Mem  repos.Store
}

func NewRouter(s repos.Store) *router {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	server := router{Rout: mux, Mem: s}

	server.Rout.Route("v1", func(r chi.Router) {
		r.Post("complex_report", jsonApi.JSON(server.CreateReport))
		r.Get("status/{job_id}", jsonApi.JSON(server.JobStatus))
	})

	return &server
}

type createReportRequest struct {
	ObjectID string `json:"object_id"`
}

type createReportResponse struct {
	JobID string `json:"job_id"`
}

func (rout *router) CreateReport(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createReportRequest
	if err := jsonApi.UnmarshalJSONRequest(&req, r); err != nil {
		return nil, err
	}

	reportID, err := use_cases.CreateReport(req.ObjectID, rout.Mem)
	if err != nil {
		return nil, err
	}

	return createReportResponse{JobID: reportID}, nil
}

type jobStatusResponse struct {
	JobID  string `json:"job_id"`
	Status string `json:"status"`
}

func (rout *router) JobStatus(ctx context.Context, r *http.Request) (interface{}, error) {
	jobID := chi.URLParam(r, "{job_id}")

	status, err := use_cases.StatusReport(jobID, rout.Mem)
	if err != nil {
		return nil, err
	}

	return jobStatusResponse{JobID: jobID, Status: status}, nil
}
