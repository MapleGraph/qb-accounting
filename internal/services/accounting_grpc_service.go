package services

import (
	"context"
	"fmt"

	"qb-accounting/internal/repository"
	"qb-accounting/internal/repository/remote"
)

type AccountingGRPCService interface {
	GetOrganizationConfig(ctx context.Context, orgID string) (*remote.OrganizationConfig, error)
}

type accountingGRPCService struct {
	repos *repository.RepositoryContainer
}

func NewAccountingGRPCService(repos *repository.RepositoryContainer) AccountingGRPCService {
	return &accountingGRPCService{repos: repos}
}

func (s *accountingGRPCService) GetOrganizationConfig(ctx context.Context, orgID string) (*remote.OrganizationConfig, error) {
	if s.repos.SetupService == nil {
		return nil, fmt.Errorf("setup service not available")
	}
	return s.repos.SetupService.GetOrganizationConfig(ctx, orgID)
}
