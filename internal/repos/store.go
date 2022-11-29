package repos

import "Challenge/internal/entity"

type Store interface {
	GetReport(reportID string) (*entity.Report, error)
	StoreJob(job entity.Job) (string, error)
	GetJob(jobID string) (*entity.Job, error)
	GetObject(objectID string) (*entity.Report, error)
	StoreReport(report entity.Report) error
}
