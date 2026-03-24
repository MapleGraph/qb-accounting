package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
	"qb-accounting/internal/utils"
)

type PostingRequestService interface {
	CreatePostingRequest(ctx context.Context, req *dto.CreatePostingRequest) (*dto.PostingRequestResponse, error)
	GetPostingRequest(ctx context.Context, id string) (*dto.PostingRequestResponse, error)
	GetPostingRequestByIdempotencyKey(ctx context.Context, bookID, key string) (*dto.PostingRequestResponse, error)
	GetPostingRequestsByStatus(ctx context.Context, status string) ([]*dto.PostingRequestResponse, error)
	UpdatePostingRequest(ctx context.Context, id string, req *dto.UpdatePostingRequest) (*dto.PostingRequestResponse, error)
}

type postingRequestService struct {
	repos *repository.RepositoryContainer
}

func NewPostingRequestService(repos *repository.RepositoryContainer) PostingRequestService {
	return &postingRequestService{repos: repos}
}

func jsonbToRawMessage(j utils.JSONB) json.RawMessage {
	if len(j) == 0 {
		return json.RawMessage("{}")
	}
	b, err := json.Marshal(j)
	if err != nil {
		return json.RawMessage("{}")
	}
	return json.RawMessage(b)
}

func postingRequestToDTO(m *models.PostingRequest) *dto.PostingRequestResponse {
	if m == nil {
		return nil
	}
	return &dto.PostingRequestResponse{
		ID:                   m.ID,
		BookID:               m.BookID,
		CompanyID:            m.CompanyID,
		BranchID:             m.BranchID,
		SourceModule:         m.SourceModule,
		SourceDocumentType:   m.SourceDocumentType,
		SourceDocumentID:     m.SourceDocumentID,
		SourceEventID:        m.SourceEventID,
		IdempotencyKey:       m.IdempotencyKey,
		Status:               string(m.RequestStatus),
		RequestedPostingDate: formatTime(m.RequestedPostingDate),
		RuleVersionID:        m.RuleVersionID,
		JournalID:            m.CurrentJournalID,
		ErrorCode:            m.ErrorCode,
		ErrorMessage:         m.ErrorMessage,
		RetryCount:           m.RetryCount,
		FirstReceivedAt:      formatTime(m.FirstReceivedAt),
		LastProcessedAt:      formatTimePtr(m.LastProcessedAt),
		LastFailedAt:         formatTimePtr(m.LastFailedAt),
		RequestPayload:       jsonbToRawMessage(m.RequestPayload),
	}
}

func postingRequestsToDTO(list []*models.PostingRequest) []*dto.PostingRequestResponse {
	out := make([]*dto.PostingRequestResponse, 0, len(list))
	for _, m := range list {
		out = append(out, postingRequestToDTO(m))
	}
	return out
}

func (s *postingRequestService) CreatePostingRequest(ctx context.Context, req *dto.CreatePostingRequest) (*dto.PostingRequestResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	sum := sha256.Sum256(req.RequestPayload)
	var payload utils.JSONB
	if err := json.Unmarshal(req.RequestPayload, &payload); err != nil {
		return nil, fmt.Errorf("request_payload: %w", err)
	}
	rd, err := parseDateTime(req.RequestedPostingDate)
	if err != nil {
		return nil, err
	}
	pr := &models.PostingRequest{
		BookID:               req.BookID,
		CompanyID:            req.CompanyID,
		BranchID:             req.BranchID,
		SourceModule:         req.SourceModule,
		SourceDocumentType:   req.SourceDocumentType,
		SourceDocumentID:     req.SourceDocumentID,
		SourceEventID:        req.SourceEventID,
		IdempotencyKey:       req.IdempotencyKey,
		RequestHash:          hex.EncodeToString(sum[:]),
		RequestStatus:        models.PostingRequestStatusReceived,
		RequestedPostingDate: rd,
		RequestPayload:       payload,
		RetryCount:           0,
	}
	if err := s.repos.PostingRequestRepo.Create(ctx, pr); err != nil {
		return nil, err
	}
	created, err := s.repos.PostingRequestRepo.GetByID(ctx, pr.ID)
	if err != nil {
		return nil, err
	}
	return postingRequestToDTO(created), nil
}

func (s *postingRequestService) GetPostingRequest(ctx context.Context, id string) (*dto.PostingRequestResponse, error) {
	m, err := s.repos.PostingRequestRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return postingRequestToDTO(m), nil
}

func (s *postingRequestService) GetPostingRequestByIdempotencyKey(ctx context.Context, bookID, key string) (*dto.PostingRequestResponse, error) {
	m, err := s.repos.PostingRequestRepo.GetByIdempotencyKey(ctx, bookID, key)
	if err != nil {
		return nil, err
	}
	return postingRequestToDTO(m), nil
}

func (s *postingRequestService) GetPostingRequestsByStatus(ctx context.Context, status string) ([]*dto.PostingRequestResponse, error) {
	list, err := s.repos.PostingRequestRepo.GetByStatus(ctx, status)
	if err != nil {
		return nil, err
	}
	return postingRequestsToDTO(list), nil
}

func (s *postingRequestService) UpdatePostingRequest(ctx context.Context, id string, req *dto.UpdatePostingRequest) (*dto.PostingRequestResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	existing, err := s.repos.PostingRequestRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("posting request not found")
	}
	now := time.Now()
	if req.RequestStatus != nil {
		existing.RequestStatus = models.PostingRequestStatus(*req.RequestStatus)
	}
	if req.RuleVersionID != nil {
		existing.RuleVersionID = req.RuleVersionID
	}
	if req.CurrentJournalID != nil {
		existing.CurrentJournalID = req.CurrentJournalID
	}
	if req.ErrorCode != nil {
		existing.ErrorCode = req.ErrorCode
	}
	if req.ErrorMessage != nil {
		existing.ErrorMessage = req.ErrorMessage
	}
	if req.RetryCount != nil {
		existing.RetryCount = *req.RetryCount
	}
	if req.RequestedPostingDate != nil {
		t, err := parseDateTime(*req.RequestedPostingDate)
		if err != nil {
			return nil, err
		}
		existing.RequestedPostingDate = t
	}
	if req.RequestPayload != nil {
		var payload utils.JSONB
		if err := json.Unmarshal(*req.RequestPayload, &payload); err != nil {
			return nil, fmt.Errorf("request_payload: %w", err)
		}
		existing.RequestPayload = payload
		sum := sha256.Sum256(*req.RequestPayload)
		existing.RequestHash = hex.EncodeToString(sum[:])
	}
	existing.LastProcessedAt = &now
	existing.UpdatedAt = &now
	if err := s.repos.PostingRequestRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	updated, err := s.repos.PostingRequestRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return postingRequestToDTO(updated), nil
}
