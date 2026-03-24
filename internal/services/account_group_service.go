package services

import (
	"context"
	"fmt"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
)

// AccountGroupService is the HTTP-facing API for account groups (DTO responses).
type AccountGroupService interface {
	CreateAccountGroup(ctx context.Context, req *dto.CreateAccountGroupRequest) (*dto.AccountGroupResponse, error)
	GetAccountGroup(ctx context.Context, id string) (*dto.AccountGroupResponse, error)
	ListAccountGroupsByBook(ctx context.Context, bookID string) ([]*dto.AccountGroupResponse, error)
	UpdateAccountGroup(ctx context.Context, id string, req *dto.UpdateAccountGroupRequest) (*dto.AccountGroupResponse, error)
	DeleteAccountGroup(ctx context.Context, id string) error
}

type accountGroupService struct {
	repos *repository.RepositoryContainer
}

func NewAccountGroupService(repos *repository.RepositoryContainer) AccountGroupService {
	return &accountGroupService{repos: repos}
}

func (s *accountGroupService) CreateAccountGroup(ctx context.Context, req *dto.CreateAccountGroupRequest) (*dto.AccountGroupResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	g := &models.AccountGroup{
		BookID:        req.BookID,
		Code:          req.Code,
		Name:          req.Name,
		ParentGroupID: req.ParentGroupID,
		AccountNature: models.AccountNature(req.AccountNature),
		SortOrder:     req.SortOrder,
		IsSystem:      req.IsSystem,
	}
	if err := s.repos.AccountGroupRepo.Create(ctx, g); err != nil {
		return nil, err
	}
	created, err := s.repos.AccountGroupRepo.GetByID(ctx, g.ID)
	if err != nil {
		return nil, err
	}
	return accountGroupModelToDTO(created), nil
}

func (s *accountGroupService) GetAccountGroup(ctx context.Context, id string) (*dto.AccountGroupResponse, error) {
	g, err := s.repos.AccountGroupRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return accountGroupModelToDTO(g), nil
}

func (s *accountGroupService) ListAccountGroupsByBook(ctx context.Context, bookID string) ([]*dto.AccountGroupResponse, error) {
	list, err := s.repos.AccountGroupRepo.GetByBookID(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return accountGroupsModelToDTO(list), nil
}

func (s *accountGroupService) UpdateAccountGroup(ctx context.Context, id string, req *dto.UpdateAccountGroupRequest) (*dto.AccountGroupResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	existing, err := s.repos.AccountGroupRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("account group not found")
	}
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.ParentGroupID != nil {
		existing.ParentGroupID = req.ParentGroupID
	}
	if req.SortOrder != nil {
		existing.SortOrder = *req.SortOrder
	}
	if err := s.repos.AccountGroupRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	updated, err := s.repos.AccountGroupRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return accountGroupModelToDTO(updated), nil
}

func (s *accountGroupService) DeleteAccountGroup(ctx context.Context, id string) error {
	return s.repos.AccountGroupRepo.Delete(ctx, id)
}
