package use_cases

import (
	"Challenge/internal/entity"
	"Challenge/internal/repos"
	"time"
)

func CreateReport(objectID string, mem repos.Store) (string, error) {
	// check if this object has a rerun in less than 5 minutes
	report, err := mem.GetReport(objectID)
	if err != nil {
		return "", err
	}
	if report.FinishedAt.After(time.Now().Add(-5 * time.Minute)) {
		return report.ID, nil
	}
	// create the job
	report = &entity.Report{
		ID:        "reporte",
		CreatedAt: time.Now(),
		Status:    "OnGoing",
	}

	job := &entity.Job{
		JobID:    "asd",
		ObjectID: objectID,
		ReportID: report.ID,
		Finished: false,
	}

	jobID, err := mem.StoreJob(*job)
	if err != nil {
		return "", nil
	}

	go func(report *entity.Report, job *entity.Job) {

		time.Sleep(time.Second * 5) //make it random

		job.Finished = true

		report.FinishedAt = time.Now()
		report.Status = "Successful"

		_, err := mem.StoreJob(*job)
		if err != nil {
			return
		}

		err = mem.StoreReport(*report)
		if err != nil {
			return
		}

	}(report, job)

	return jobID, nil
}

func StatusReport(jobID string, mem repos.Store) (string, error) {
	job, err := mem.GetJob(jobID)
	if err != nil {
		return "", err
	}

	report, err := mem.GetReport(job.ReportID)
	if err != nil {
		return "", err
	}

	return report.Status, nil
}
