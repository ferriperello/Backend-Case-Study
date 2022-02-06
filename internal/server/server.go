package server

import (
	use_cases "Challenge/internal/use-cases"
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

type createReportRequest struct {
	ObjectID string `json:"object_id"`
}

type createReportResponse struct {
	JobID string `json:"job_id"`
}

func CreateReport(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createReportRequest
	if err := jsonApi.UnmarshalJSONRequest(&req, r); err != nil {
		return nil, err
	}

	reportID, err := use_cases.CreateReport(req.ObjectID)
	if err != nil {
		return nil, err
	}

	return createReportResponse{JobID: reportID}, nil
}

type jobStatusResponse struct {
	JobID  string `json:"job_id"`
	Status string `json:"status"`
}

func JobStatus(ctx context.Context, r *http.Request) (interface{}, error) {
	jobID := chi.URLParam(r, "{job_id}")

	status, err := use_cases.StatusReport(jobID)
	if err != nil {
		return nil, err
	}

	return jobStatusResponse{JobID: jobID, Status: status}, nil
}
