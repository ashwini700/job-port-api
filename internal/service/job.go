package service

import (
	"context"

	"job-port-api/internal/models"

)

func (s *Service) AddJob(ctx context.Context, jobData models.Job) (models.Job, error) {
	jobData, err := s.UserRepo.CreateJob(ctx, jobData)
	if err != nil {
		return models.Job{}, err
	}
	return jobData, nil
}
func (s *Service) FetchJobDetailsById(ctx context.Context, cid uint64) (models.Job, error) {
	jobData, err := s.UserRepo.Fetchjob(ctx, cid)
	if err != nil {
		return models.Job{}, err
	}
	return jobData, nil

}

func (s *Service) FetchJobPosts(ctx context.Context) ([]models.Job, error) {
	jobData, err := s.UserRepo.FetchJobPosts(ctx)
	if err != nil {
		return nil, err
	}
	return jobData, nil

}

func (s *Service) FetchJobDetails(ctx context.Context, cid uint64) ([]models.Job, error) {
	jobData, err := s.UserRepo.FetchJobByCompanyId(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
