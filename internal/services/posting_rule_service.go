package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
	"qb-accounting/internal/utils"
)

// PostingRuleService is the HTTP-facing API for posting rule versions (DTO responses).
type PostingRuleService interface {
	CreatePostingRuleVersion(ctx context.Context, req *dto.CreatePostingRuleVersionRequest) (*dto.PostingRuleVersionResponse, error)
	GetPostingRuleVersion(ctx context.Context, id string) (*dto.PostingRuleVersionResponse, error)
	GetActivePostingRule(ctx context.Context, sourceModule, sourceDocType string) (*dto.PostingRuleVersionResponse, error)
	UpdatePostingRuleVersion(ctx context.Context, id string, req *dto.UpdatePostingRuleVersionRequest) (*dto.PostingRuleVersionResponse, error)
	DeletePostingRuleVersion(ctx context.Context, id string) error
}

type postingRuleService struct {
	repos *repository.RepositoryContainer
}

func NewPostingRuleService(repos *repository.RepositoryContainer) PostingRuleService {
	return &postingRuleService{repos: repos}
}

func (s *postingRuleService) CreatePostingRuleVersion(ctx context.Context, req *dto.CreatePostingRuleVersionRequest) (*dto.PostingRuleVersionResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	var payload utils.JSONB
	if err := json.Unmarshal(req.RulePayload, &payload); err != nil {
		return nil, fmt.Errorf("rule_payload: %w", err)
	}
	status := models.PostingRuleStatusDraft
	if req.Status != "" {
		status = models.PostingRuleStatus(req.Status)
	}
	effectiveFrom, err := parseDateTime(req.EffectiveFrom)
	if err != nil {
		return nil, err
	}
	var effectiveTo *time.Time
	if req.EffectiveTo != nil {
		t, err := parseDateTime(*req.EffectiveTo)
		if err != nil {
			return nil, err
		}
		effectiveTo = &t
	}
	rule := &models.PostingRuleVersion{
		SourceModule:       req.SourceModule,
		SourceDocumentType: req.SourceDocumentType,
		VersionNo:          req.VersionNo,
		Name:               req.Name,
		Status:             status,
		EffectiveFrom:      effectiveFrom,
		EffectiveTo:        effectiveTo,
		RulePayload:        payload,
		Notes:              req.Notes,
	}
	if err := s.repos.PostingRuleVersionRepo.Create(ctx, rule); err != nil {
		return nil, err
	}
	created, err := s.repos.PostingRuleVersionRepo.GetByID(ctx, rule.ID)
	if err != nil {
		return nil, err
	}
	return postingRuleVersionModelToDTO(created)
}

func (s *postingRuleService) GetPostingRuleVersion(ctx context.Context, id string) (*dto.PostingRuleVersionResponse, error) {
	m, err := s.repos.PostingRuleVersionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return postingRuleVersionModelToDTO(m)
}

func (s *postingRuleService) GetActivePostingRule(ctx context.Context, sourceModule, sourceDocType string) (*dto.PostingRuleVersionResponse, error) {
	if sourceModule == "" || sourceDocType == "" {
		return nil, fmt.Errorf("source_module and source_document_type are required")
	}
	m, err := s.repos.PostingRuleVersionRepo.GetActiveBySourceDocType(ctx, sourceModule, sourceDocType)
	if err != nil {
		return nil, err
	}
	return postingRuleVersionModelToDTO(m)
}

func (s *postingRuleService) UpdatePostingRuleVersion(ctx context.Context, id string, req *dto.UpdatePostingRuleVersionRequest) (*dto.PostingRuleVersionResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	existing, err := s.repos.PostingRuleVersionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("posting rule not found")
	}
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Status != nil {
		existing.Status = models.PostingRuleStatus(*req.Status)
	}
	if req.EffectiveTo != nil {
		t, err := parseDateTime(*req.EffectiveTo)
		if err != nil {
			return nil, err
		}
		existing.EffectiveTo = &t
	}
	if req.RulePayload != nil {
		var payload utils.JSONB
		if err := json.Unmarshal(*req.RulePayload, &payload); err != nil {
			return nil, fmt.Errorf("rule_payload: %w", err)
		}
		existing.RulePayload = payload
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}
	if err := s.repos.PostingRuleVersionRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	updated, err := s.repos.PostingRuleVersionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return postingRuleVersionModelToDTO(updated)
}

func (s *postingRuleService) DeletePostingRuleVersion(ctx context.Context, id string) error {
	return s.repos.PostingRuleVersionRepo.Delete(ctx, id)
}
