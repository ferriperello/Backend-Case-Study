package repos

import "context"

type Store interface {
	GetReport
	CreateJob
	GetObject
}

type GetReport interface {
	GetReport(ctx context.Context, reportID string) error
}

type CreateJob interface {
	CreateJob(ctx context.Context, objectID string) (string, error)
}

type GetObject interface {
	GetObject(ctx context.Context, objectID string) error
}
