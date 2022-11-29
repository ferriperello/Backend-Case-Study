package repos

import (
	"Challenge/internal/entity"
)

func NewMemStore() *MemStore {
	q := &MemStore{
		objects: make([]entity.Object, 10),
		jobs:    make([]entity.Job, 10),
		reports: make([]entity.Report, 10),
	}
	return q
}

type MemStore struct {
	objects []entity.Object
	jobs    []entity.Job
	reports []entity.Report
}

func (m MemStore) GetReport(reportID string) (*entity.Report, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) StoreJob(job entity.Job) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) GetObject(objectID string) (*entity.Report, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) StoreReport(report entity.Report) error {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) GetJob(jobID string) (*entity.Job, error) {
	//TODO implement me
	panic("implement me")
}
