package services

import (
	"context"
	"fmt"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
)

// FiscalYearService defines fiscal year operations for HTTP handlers (DTO responses).
type FiscalYearService interface {
	CreateFiscalYear(ctx context.Context, req *dto.CreateFiscalYearRequest) (*dto.FiscalYearResponse, error)
	GetFiscalYear(ctx context.Context, id string) (*dto.FiscalYearResponse, error)
	ListFiscalYearsByBook(ctx context.Context, bookID string) ([]*dto.FiscalYearResponse, error)
	UpdateFiscalYear(ctx context.Context, id string, req *dto.UpdateFiscalYearRequest) (*dto.FiscalYearResponse, error)
	DeleteFiscalYear(ctx context.Context, id string) error
}

type fiscalYearService struct {
	repos *repository.RepositoryContainer
}

func NewFiscalYearService(repos *repository.RepositoryContainer) FiscalYearService {
	return &fiscalYearService{repos: repos}
}

func (s *fiscalYearService) CreateFiscalYear(ctx context.Context, req *dto.CreateFiscalYearRequest) (*dto.FiscalYearResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	start, err := parseDateTime(req.StartDate)
	if err != nil {
		return nil, err
	}
	end, err := parseDateTime(req.EndDate)
	if err != nil {
		return nil, err
	}
	status := models.FiscalYearStatusDraft
	if req.Status != "" {
		status = models.FiscalYearStatus(req.Status)
	}
	fy := &models.FiscalYear{
		BookID:        req.BookID,
		CompanyID:     req.CompanyID,
		Code:          req.Code,
		Name:          req.Name,
		StartDate:     start,
		EndDate:       end,
		Status:        status,
		CloseSequence: 0,
	}
	if err := s.repos.FiscalYearRepo.Create(ctx, fy); err != nil {
		return nil, err
	}
	created, err := s.repos.FiscalYearRepo.GetByID(ctx, fy.ID)
	if err != nil {
		return nil, err
	}
	return fiscalYearModelToDTO(created), nil
}

func (s *fiscalYearService) GetFiscalYear(ctx context.Context, id string) (*dto.FiscalYearResponse, error) {
	fy, err := s.repos.FiscalYearRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return fiscalYearModelToDTO(fy), nil
}

func (s *fiscalYearService) ListFiscalYearsByBook(ctx context.Context, bookID string) ([]*dto.FiscalYearResponse, error) {
	list, err := s.repos.FiscalYearRepo.GetByBookID(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return fiscalYearsModelToDTO(list), nil
}

func (s *fiscalYearService) UpdateFiscalYear(ctx context.Context, id string, req *dto.UpdateFiscalYearRequest) (*dto.FiscalYearResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	existing, err := s.repos.FiscalYearRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("fiscal year not found")
	}
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.EndDate != nil {
		t, err := parseDateTime(*req.EndDate)
		if err != nil {
			return nil, err
		}
		existing.EndDate = t
	}
	if req.Status != nil {
		existing.Status = models.FiscalYearStatus(*req.Status)
	}
	if err := s.repos.FiscalYearRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	updated, err := s.repos.FiscalYearRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return fiscalYearModelToDTO(updated), nil
}

func (s *fiscalYearService) DeleteFiscalYear(ctx context.Context, id string) error {
	return s.repos.FiscalYearRepo.Delete(ctx, id)
}
