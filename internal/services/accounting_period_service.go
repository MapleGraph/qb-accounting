package services

import (
	"context"
	"fmt"
	"time"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
)

// AccountingPeriodService defines accounting period operations for HTTP handlers (DTO responses).
type AccountingPeriodService interface {
	CreateAccountingPeriod(ctx context.Context, req *dto.CreateAccountingPeriodRequest) (*dto.AccountingPeriodResponse, error)
	GetAccountingPeriod(ctx context.Context, id string) (*dto.AccountingPeriodResponse, error)
	ListAccountingPeriodsByBook(ctx context.Context, bookID string) ([]*dto.AccountingPeriodResponse, error)
	ListAccountingPeriodsByFiscalYear(ctx context.Context, fiscalYearID string) ([]*dto.AccountingPeriodResponse, error)
	UpdateAccountingPeriod(ctx context.Context, id string, req *dto.UpdateAccountingPeriodRequest) (*dto.AccountingPeriodResponse, error)
	DeleteAccountingPeriod(ctx context.Context, id string) error
}

type accountingPeriodService struct {
	repos *repository.RepositoryContainer
}

func NewAccountingPeriodService(repos *repository.RepositoryContainer) AccountingPeriodService {
	return &accountingPeriodService{repos: repos}
}

func (s *accountingPeriodService) CreateAccountingPeriod(ctx context.Context, req *dto.CreateAccountingPeriodRequest) (*dto.AccountingPeriodResponse, error) {
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
	status := models.PeriodStatusDraft
	if req.Status != "" {
		status = models.PeriodStatus(req.Status)
	}
	p := &models.AccountingPeriod{
		BookID:             req.BookID,
		FiscalYearID:       req.FiscalYearID,
		CompanyID:          req.CompanyID,
		PeriodNo:           req.PeriodNo,
		PeriodName:         req.PeriodName,
		StartDate:          start,
		EndDate:            end,
		Status:             status,
		IsAdjustmentPeriod: req.IsAdjustmentPeriod,
	}
	if err := s.repos.AccountingPeriodRepo.Create(ctx, p); err != nil {
		return nil, err
	}
	created, err := s.repos.AccountingPeriodRepo.GetByID(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	return accountingPeriodModelToDTO(created), nil
}

func (s *accountingPeriodService) GetAccountingPeriod(ctx context.Context, id string) (*dto.AccountingPeriodResponse, error) {
	p, err := s.repos.AccountingPeriodRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return accountingPeriodModelToDTO(p), nil
}

func (s *accountingPeriodService) ListAccountingPeriodsByBook(ctx context.Context, bookID string) ([]*dto.AccountingPeriodResponse, error) {
	list, err := s.repos.AccountingPeriodRepo.GetByBookID(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return accountingPeriodsModelToDTO(list), nil
}

func (s *accountingPeriodService) ListAccountingPeriodsByFiscalYear(ctx context.Context, fiscalYearID string) ([]*dto.AccountingPeriodResponse, error) {
	list, err := s.repos.AccountingPeriodRepo.GetByFiscalYearID(ctx, fiscalYearID)
	if err != nil {
		return nil, err
	}
	return accountingPeriodsModelToDTO(list), nil
}

func (s *accountingPeriodService) UpdateAccountingPeriod(ctx context.Context, id string, req *dto.UpdateAccountingPeriodRequest) (*dto.AccountingPeriodResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	existing, err := s.repos.AccountingPeriodRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("accounting period not found")
	}
	now := time.Now()
	if req.PeriodName != nil {
		existing.PeriodName = *req.PeriodName
	}
	if req.Status != nil {
		existing.Status = models.PeriodStatus(*req.Status)
	}
	if req.LockReason != nil {
		existing.LockReason = req.LockReason
	}
	if req.ApplySoftLock != nil && *req.ApplySoftLock {
		existing.SoftLockedAt = &now
		if req.SoftLockedBy != nil {
			existing.SoftLockedBy = req.SoftLockedBy
		}
	}
	if req.ReleaseSoftLock != nil && *req.ReleaseSoftLock {
		existing.SoftLockedAt = nil
		existing.SoftLockedBy = nil
	}
	if req.ApplyHardLock != nil && *req.ApplyHardLock {
		existing.HardLockedAt = &now
		if req.HardLockedBy != nil {
			existing.HardLockedBy = req.HardLockedBy
		}
	}
	if req.ReleaseHardLock != nil && *req.ReleaseHardLock {
		existing.HardLockedAt = nil
		existing.HardLockedBy = nil
	}
	if req.SoftLockedBy != nil && (req.ApplySoftLock == nil || !*req.ApplySoftLock) {
		existing.SoftLockedBy = req.SoftLockedBy
	}
	if req.HardLockedBy != nil && (req.ApplyHardLock == nil || !*req.ApplyHardLock) {
		existing.HardLockedBy = req.HardLockedBy
	}
	if req.ClosedBy != nil {
		existing.ClosedBy = req.ClosedBy
		existing.ClosedAt = &now
	}
	if err := s.repos.AccountingPeriodRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	updated, err := s.repos.AccountingPeriodRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return accountingPeriodModelToDTO(updated), nil
}

func (s *accountingPeriodService) DeleteAccountingPeriod(ctx context.Context, id string) error {
	return s.repos.AccountingPeriodRepo.Delete(ctx, id)
}
