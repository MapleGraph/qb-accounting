package services

import (
	"context"
	"fmt"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
)

// JournalBatchService is the HTTP-facing API for journal batches (DTO responses).
type JournalBatchService interface {
	CreateJournalBatch(ctx context.Context, req *dto.CreateJournalBatchRequest) (*dto.JournalBatchResponse, error)
	GetJournalBatch(ctx context.Context, id string) (*dto.JournalBatchResponse, error)
	ListJournalBatchesByBook(ctx context.Context, bookID string) ([]*dto.JournalBatchResponse, error)
	UpdateJournalBatch(ctx context.Context, id string, req *dto.UpdateJournalBatchRequest) (*dto.JournalBatchResponse, error)
	DeleteJournalBatch(ctx context.Context, id string) error
}

type journalBatchService struct {
	repos *repository.RepositoryContainer
}

func NewJournalBatchService(repos *repository.RepositoryContainer) JournalBatchService {
	return &journalBatchService{repos: repos}
}

func (s *journalBatchService) CreateJournalBatch(ctx context.Context, req *dto.CreateJournalBatchRequest) (*dto.JournalBatchResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	postingDate, err := parseDateTime(req.PostingDate)
	if err != nil {
		return nil, err
	}
	b := &models.JournalBatch{
		BookID:       req.BookID,
		CompanyID:    req.CompanyID,
		BranchID:     req.BranchID,
		BatchNo:      req.BatchNo,
		SourceModule: req.SourceModule,
		BatchType:    req.BatchType,
		PostingDate:  postingDate,
		Status:       models.JournalBatchStatusOpen,
		Narration:    req.Narration,
	}
	if err := s.repos.JournalBatchRepo.Create(ctx, b); err != nil {
		return nil, err
	}
	created, err := s.repos.JournalBatchRepo.GetByID(ctx, b.ID)
	if err != nil {
		return nil, err
	}
	return journalBatchModelToDTO(created), nil
}

func (s *journalBatchService) GetJournalBatch(ctx context.Context, id string) (*dto.JournalBatchResponse, error) {
	b, err := s.repos.JournalBatchRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return journalBatchModelToDTO(b), nil
}

func (s *journalBatchService) ListJournalBatchesByBook(ctx context.Context, bookID string) ([]*dto.JournalBatchResponse, error) {
	list, err := s.repos.JournalBatchRepo.GetByBookID(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return journalBatchesModelToDTO(list), nil
}

func (s *journalBatchService) UpdateJournalBatch(ctx context.Context, id string, req *dto.UpdateJournalBatchRequest) (*dto.JournalBatchResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	existing, err := s.repos.JournalBatchRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("journal batch not found")
	}
	if req.SourceModule != nil {
		existing.SourceModule = req.SourceModule
	}
	if req.BatchType != nil {
		existing.BatchType = *req.BatchType
	}
	if req.PostingDate != nil {
		t, err := parseDateTime(*req.PostingDate)
		if err != nil {
			return nil, err
		}
		existing.PostingDate = t
	}
	if req.Status != nil {
		existing.Status = models.JournalBatchStatus(*req.Status)
	}
	if req.Narration != nil {
		existing.Narration = req.Narration
	}
	if err := s.repos.JournalBatchRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	updated, err := s.repos.JournalBatchRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return journalBatchModelToDTO(updated), nil
}

func (s *journalBatchService) DeleteJournalBatch(ctx context.Context, id string) error {
	return s.repos.JournalBatchRepo.Delete(ctx, id)
}
