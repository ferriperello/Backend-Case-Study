package repos

import (
	"Challenge/internal/entity"
	"context"
)

type MemStore struct {
	objects []entity.Object
	jobs    []entity.Job
	reports []entity.Report
}

func (m *MemStore) GetReport(ctx context.Context, reportID string) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemStore) CreateJob(ctx context.Context, objectID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemStore) GetObject(ctx context.Context, objectID string) error {
	//TODO implement me
	panic("implement me")
}
