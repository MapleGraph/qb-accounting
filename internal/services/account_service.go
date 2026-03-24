package services

import (
	"context"
	"fmt"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
)

// AccountService is the HTTP-facing API for ledger accounts (DTO responses).
type AccountService interface {
	CreateAccount(ctx context.Context, req *dto.CreateAccountRequest) (*dto.AccountResponse, error)
	GetAccount(ctx context.Context, id string) (*dto.AccountResponse, error)
	ListAccountsByBook(ctx context.Context, bookID string) ([]*dto.AccountResponse, error)
	ListAccountsByCompany(ctx context.Context, companyID string) ([]*dto.AccountResponse, error)
	UpdateAccount(ctx context.Context, id string, req *dto.UpdateAccountRequest) (*dto.AccountResponse, error)
	DeleteAccount(ctx context.Context, id string) error
}

type accountService struct {
	repos *repository.RepositoryContainer
}

func NewAccountService(repos *repository.RepositoryContainer) AccountService {
	return &accountService{repos: repos}
}

func (s *accountService) CreateAccount(ctx context.Context, req *dto.CreateAccountRequest) (*dto.AccountResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	a := &models.Account{
		BookID:             req.BookID,
		CompanyID:          req.CompanyID,
		Code:               req.Code,
		Name:               req.Name,
		DisplayName:        req.DisplayName,
		GroupID:            req.GroupID,
		ParentAccountID:    req.ParentAccountID,
		AccountNature:      models.AccountNature(req.AccountNature),
		Usage:              models.AccountUsage(req.Usage),
		NormalBalance:      req.NormalBalance,
		ControlType:        req.ControlType,
		AllowManualPosting: req.AllowManualPosting,
		RequireParty:       req.RequireParty,
		RequireBranch:      req.RequireBranch,
		RequireCostCenter:  req.RequireCostCenter,
		RequireEmployee:    req.RequireEmployee,
		RequireTaxBreakup:  req.RequireTaxBreakup,
		IsSystem:           req.IsSystem,
		IsActive:           req.IsActive,
	}
	if err := s.repos.AccountRepo.Create(ctx, a); err != nil {
		return nil, err
	}
	created, err := s.repos.AccountRepo.GetByID(ctx, a.ID)
	if err != nil {
		return nil, err
	}
	return accountModelToDTO(created), nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*dto.AccountResponse, error) {
	a, err := s.repos.AccountRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return accountModelToDTO(a), nil
}

func (s *accountService) ListAccountsByBook(ctx context.Context, bookID string) ([]*dto.AccountResponse, error) {
	list, err := s.repos.AccountRepo.GetByBookID(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return accountsModelToDTO(list), nil
}

func (s *accountService) ListAccountsByCompany(ctx context.Context, companyID string) ([]*dto.AccountResponse, error) {
	list, err := s.repos.AccountRepo.GetByCompanyID(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return accountsModelToDTO(list), nil
}

func (s *accountService) UpdateAccount(ctx context.Context, id string, req *dto.UpdateAccountRequest) (*dto.AccountResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	existing, err := s.repos.AccountRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("account not found")
	}
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.DisplayName != nil {
		existing.DisplayName = req.DisplayName
	}
	if req.ParentAccountID != nil {
		existing.ParentAccountID = req.ParentAccountID
	}
	if req.ControlType != nil {
		existing.ControlType = req.ControlType
	}
	if req.AllowManualPosting != nil {
		existing.AllowManualPosting = *req.AllowManualPosting
	}
	if req.RequireParty != nil {
		existing.RequireParty = *req.RequireParty
	}
	if req.RequireBranch != nil {
		existing.RequireBranch = *req.RequireBranch
	}
	if req.RequireCostCenter != nil {
		existing.RequireCostCenter = *req.RequireCostCenter
	}
	if req.RequireEmployee != nil {
		existing.RequireEmployee = *req.RequireEmployee
	}
	if req.RequireTaxBreakup != nil {
		existing.RequireTaxBreakup = *req.RequireTaxBreakup
	}
	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}
	if err := s.repos.AccountRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	updated, err := s.repos.AccountRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return accountModelToDTO(updated), nil
}

func (s *accountService) DeleteAccount(ctx context.Context, id string) error {
	return s.repos.AccountRepo.Delete(ctx, id)
}
