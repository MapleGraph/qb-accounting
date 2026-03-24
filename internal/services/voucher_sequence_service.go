package services

import (
	"context"
	"fmt"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
)

// VoucherSequenceService is the HTTP-facing API for voucher sequences (DTO responses).
type VoucherSequenceService interface {
	CreateVoucherSequence(ctx context.Context, req *dto.CreateVoucherSequenceRequest) (*dto.VoucherSequenceResponse, error)
	GetVoucherSequence(ctx context.Context, id string) (*dto.VoucherSequenceResponse, error)
	UpdateVoucherSequence(ctx context.Context, id string, req *dto.UpdateVoucherSequenceRequest) (*dto.VoucherSequenceResponse, error)
	DeleteVoucherSequence(ctx context.Context, id string) error
}

type voucherSequenceService struct {
	repos *repository.RepositoryContainer
}

func NewVoucherSequenceService(repos *repository.RepositoryContainer) VoucherSequenceService {
	return &voucherSequenceService{repos: repos}
}

func (s *voucherSequenceService) CreateVoucherSequence(ctx context.Context, req *dto.CreateVoucherSequenceRequest) (*dto.VoucherSequenceResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	v := &models.VoucherSequence{
		BookID:      req.BookID,
		CompanyID:   req.CompanyID,
		BranchID:    req.BranchID,
		VoucherType: req.VoucherType,
		Prefix:      req.Prefix,
		Suffix:      req.Suffix,
		Padding:     req.Padding,
		NextNumber:  req.NextNumber,
		ResetPolicy: req.ResetPolicy,
		IsActive:    req.IsActive,
	}
	if err := s.repos.VoucherSequenceRepo.Create(ctx, v); err != nil {
		return nil, err
	}
	created, err := s.repos.VoucherSequenceRepo.GetByID(ctx, v.ID)
	if err != nil {
		return nil, err
	}
	return voucherSequenceModelToDTO(created), nil
}

func (s *voucherSequenceService) GetVoucherSequence(ctx context.Context, id string) (*dto.VoucherSequenceResponse, error) {
	v, err := s.repos.VoucherSequenceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return voucherSequenceModelToDTO(v), nil
}

func (s *voucherSequenceService) UpdateVoucherSequence(ctx context.Context, id string, req *dto.UpdateVoucherSequenceRequest) (*dto.VoucherSequenceResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	existing, err := s.repos.VoucherSequenceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("voucher sequence not found")
	}
	if req.BranchID != nil {
		existing.BranchID = req.BranchID
	}
	if req.Prefix != nil {
		existing.Prefix = req.Prefix
	}
	if req.Suffix != nil {
		existing.Suffix = req.Suffix
	}
	if req.Padding != nil {
		existing.Padding = *req.Padding
	}
	if req.NextNumber != nil {
		existing.NextNumber = *req.NextNumber
	}
	if req.ResetPolicy != nil {
		existing.ResetPolicy = *req.ResetPolicy
	}
	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}
	if err := s.repos.VoucherSequenceRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	updated, err := s.repos.VoucherSequenceRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return voucherSequenceModelToDTO(updated), nil
}

func (s *voucherSequenceService) DeleteVoucherSequence(ctx context.Context, id string) error {
	return s.repos.VoucherSequenceRepo.Delete(ctx, id)
}
