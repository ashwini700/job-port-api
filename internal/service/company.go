package service

import (
	"context"

	"job-port-api/internal/models"
)

func (s *Service) AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error) {
	companyData, err := s.UserRepo.CreateCompany(ctx, companyData)
	if err != nil {
		return models.Company{}, err
	}
	return companyData, nil
}
func (s *Service) FetchCompByid(ctx context.Context, cid uint64) (models.Company, error) {
	companyData, err := s.UserRepo.FetchCompany(ctx, cid)
	if err != nil {
		return models.Company{}, err
	}
	return companyData, nil
}
func (s *Service) FetchAllCompanies(ctx context.Context) ([]models.Company, error) {
	companyDetails, err := s.UserRepo.FetchAllCompanies(ctx)
	if err != nil {
		return nil, err
	}
	return companyDetails, nil
}
