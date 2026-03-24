package services

import (
	"context"
	"fmt"
	"time"

	qbpostgres "github.com/MapleGraph/qb-core/v2/pkg/postgres"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
)

type OpenItemService interface {
	CreateOpenItem(ctx context.Context, req *dto.CreateOpenItemRequest) (*dto.OpenItemResponse, error)
	GetOpenItem(ctx context.Context, id string) (*dto.OpenItemResponse, error)
	GetOpenItemsByParty(ctx context.Context, partyID string) ([]*dto.OpenItemResponse, error)
	GetOpenItemsByBook(ctx context.Context, bookID string) ([]*dto.OpenItemResponse, error)
	GetOpenItemsByStatus(ctx context.Context, bookID, side, status string) ([]*dto.OpenItemResponse, error)
	UpdateOpenItem(ctx context.Context, id string, req *dto.UpdateOpenItemRequest) (*dto.OpenItemResponse, error)
	DeleteOpenItem(ctx context.Context, id string) error
	CreateAllocation(ctx context.Context, req *dto.AllocationRequest) (*dto.AllocationResponse, error)
	CreateAdjustment(ctx context.Context, req *dto.AdjustmentRequest) (*dto.AdjustmentResponse, error)
	GetAllocationsByOpenItem(ctx context.Context, openItemID string) ([]*dto.AllocationResponse, error)
	GetAdjustmentsByOpenItem(ctx context.Context, openItemID string) ([]*dto.AdjustmentResponse, error)
}

type openItemService struct {
	repos *repository.RepositoryContainer
}

func NewOpenItemService(repos *repository.RepositoryContainer) OpenItemService {
	return &openItemService{repos: repos}
}

var _ = qbpostgres.WithTransaction

func openItemToDTO(o *models.OpenItem) *dto.OpenItemResponse {
	if o == nil {
		return nil
	}
	return &dto.OpenItemResponse{
		ID:                 o.ID,
		BookID:             o.BookID,
		CompanyID:          o.CompanyID,
		BranchID:           o.BranchID,
		ItemSide:           string(o.ItemSide),
		ItemStatus:         string(o.ItemStatus),
		PartyID:            o.PartyID,
		PartyType:          string(o.PartyType),
		ControlAccountID:   o.ControlAccountID,
		SourceModule:       o.SourceModule,
		SourceDocumentType: o.SourceDocumentType,
		SourceDocumentID:   o.SourceDocumentID,
		SourceLineRef:      o.SourceLineRef,
		JournalID:          o.JournalID,
		JournalLineID:      o.JournalLineID,
		DocumentNo:         o.DocumentNo,
		DocumentDate:       formatTime(o.DocumentDate),
		DueDate:            formatTimePtr(o.DueDate),
		CurrencyCode:       o.CurrencyCode,
		ExchangeRate:       o.ExchangeRate,
		OriginalAmountTxn:  o.OriginalAmountTxn,
		OriginalAmountBase: o.OriginalAmountBase,
		OpenAmountTxn:      o.OpenAmountTxn,
		OpenAmountBase:     o.OpenAmountBase,
		SettledAmountTxn:   o.SettledAmountTxn,
		SettledAmountBase:  o.SettledAmountBase,
		WriteoffAmountTxn:  o.WriteoffAmountTxn,
		WriteoffAmountBase: o.WriteoffAmountBase,
		Remarks:            o.Remarks,
	}
}

func openItemsToDTO(list []*models.OpenItem) []*dto.OpenItemResponse {
	out := make([]*dto.OpenItemResponse, 0, len(list))
	for _, o := range list {
		out = append(out, openItemToDTO(o))
	}
	return out
}

func allocationToDTO(a *models.OpenItemAllocation) *dto.AllocationResponse {
	if a == nil {
		return nil
	}
	return &dto.AllocationResponse{
		ID:                     a.ID,
		BookID:                 a.BookID,
		CompanyID:              a.CompanyID,
		AllocationStatus:       string(a.AllocationStatus),
		AllocationDate:         formatTime(a.AllocationDate),
		FromOpenItemID:         a.FromOpenItemID,
		ToOpenItemID:           a.ToOpenItemID,
		AllocationCurrencyCode: a.AllocationCurrencyCode,
		AllocationAmountTxn:    a.AllocationAmountTxn,
		AllocationAmountBase:   a.AllocationAmountBase,
		AllocationJournalID:    a.AllocationJournalID,
		ReferenceNo:            a.ReferenceNo,
		ReversalOfAllocationID: a.ReversalOfAllocationID,
		Notes:                  a.Notes,
	}
}

func adjustmentToDTO(a *models.OpenItemAdjustment) *dto.AdjustmentResponse {
	if a == nil {
		return nil
	}
	return &dto.AdjustmentResponse{
		ID:                  a.ID,
		OpenItemID:          a.OpenItemID,
		AdjustmentType:      string(a.AdjustmentType),
		AdjustmentDate:      formatTime(a.AdjustmentDate),
		AdjustmentJournalID: a.AdjustmentJournalID,
		AmountTxn:           a.AmountTxn,
		AmountBase:          a.AmountBase,
		ReasonCode:          a.ReasonCode,
		Notes:               a.Notes,
	}
}

func (s *openItemService) CreateOpenItem(ctx context.Context, req *dto.CreateOpenItemRequest) (*dto.OpenItemResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	docDate, err := parseDateTime(req.DocumentDate)
	if err != nil {
		return nil, err
	}
	var due *time.Time
	if req.DueDate != nil {
		t, err := parseDateTime(*req.DueDate)
		if err != nil {
			return nil, err
		}
		due = &t
	}
	o := &models.OpenItem{
		BookID:             req.BookID,
		CompanyID:          req.CompanyID,
		BranchID:           req.BranchID,
		ItemSide:           models.OpenItemSide(req.ItemSide),
		ItemStatus:         models.OpenItemStatusOpen,
		PartyID:            req.PartyID,
		PartyType:          models.PartyType(req.PartyType),
		ControlAccountID:   req.ControlAccountID,
		SourceModule:       req.SourceModule,
		SourceDocumentType: req.SourceDocumentType,
		SourceDocumentID:   req.SourceDocumentID,
		JournalID:          req.JournalID,
		DocumentNo:         req.DocumentNo,
		DocumentDate:       docDate,
		DueDate:            due,
		CurrencyCode:       req.CurrencyCode,
		ExchangeRate:       req.ExchangeRate,
		OriginalAmountTxn:  req.OriginalAmountTxn,
		OriginalAmountBase: req.OriginalAmountBase,
		OpenAmountTxn:      req.OriginalAmountTxn,
		OpenAmountBase:     req.OriginalAmountBase,
	}
	if err := s.repos.OpenItemRepo.Create(ctx, o); err != nil {
		return nil, err
	}
	created, err := s.repos.OpenItemRepo.GetByID(ctx, o.ID)
	if err != nil {
		return nil, err
	}
	return openItemToDTO(created), nil
}

func (s *openItemService) GetOpenItem(ctx context.Context, id string) (*dto.OpenItemResponse, error) {
	o, err := s.repos.OpenItemRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return openItemToDTO(o), nil
}

func (s *openItemService) GetOpenItemsByParty(ctx context.Context, partyID string) ([]*dto.OpenItemResponse, error) {
	list, err := s.repos.OpenItemRepo.GetByPartyID(ctx, partyID)
	if err != nil {
		return nil, err
	}
	return openItemsToDTO(list), nil
}

func (s *openItemService) GetOpenItemsByBook(ctx context.Context, bookID string) ([]*dto.OpenItemResponse, error) {
	list, err := s.repos.OpenItemRepo.GetByBookID(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return openItemsToDTO(list), nil
}

func (s *openItemService) GetOpenItemsByStatus(ctx context.Context, bookID, side, status string) ([]*dto.OpenItemResponse, error) {
	list, err := s.repos.OpenItemRepo.GetByStatus(ctx, bookID, side, status)
	if err != nil {
		return nil, err
	}
	return openItemsToDTO(list), nil
}

func (s *openItemService) UpdateOpenItem(ctx context.Context, id string, req *dto.UpdateOpenItemRequest) (*dto.OpenItemResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	existing, err := s.repos.OpenItemRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, fmt.Errorf("open item not found")
	}
	if req.DueDate != nil {
		t, err := parseDateTime(*req.DueDate)
		if err != nil {
			return nil, err
		}
		existing.DueDate = &t
	}
	if req.ItemStatus != nil {
		existing.ItemStatus = models.OpenItemStatus(*req.ItemStatus)
	}
	if req.Remarks != nil {
		existing.Remarks = req.Remarks
	}
	now := time.Now()
	existing.UpdatedAt = &now
	if err := s.repos.OpenItemRepo.Update(ctx, existing); err != nil {
		return nil, err
	}
	updated, err := s.repos.OpenItemRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return openItemToDTO(updated), nil
}

func (s *openItemService) DeleteOpenItem(ctx context.Context, id string) error {
	return s.repos.OpenItemRepo.Delete(ctx, id)
}

func (s *openItemService) CreateAllocation(ctx context.Context, req *dto.AllocationRequest) (*dto.AllocationResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	from, err := s.repos.OpenItemRepo.GetByID(ctx, req.FromOpenItemID)
	if err != nil {
		return nil, err
	}
	if from == nil {
		return nil, fmt.Errorf("from open item not found")
	}
	to, err := s.repos.OpenItemRepo.GetByID(ctx, req.ToOpenItemID)
	if err != nil {
		return nil, err
	}
	if to == nil {
		return nil, fmt.Errorf("to open item not found")
	}
	ad, err := parseDateTime(req.AllocationDate)
	if err != nil {
		return nil, err
	}
	a := &models.OpenItemAllocation{
		BookID:                 from.BookID,
		CompanyID:              from.CompanyID,
		AllocationStatus:       models.AllocationStatusApplied,
		AllocationDate:         ad,
		FromOpenItemID:         req.FromOpenItemID,
		ToOpenItemID:           req.ToOpenItemID,
		AllocationCurrencyCode: req.AllocationCurrencyCode,
		AllocationAmountTxn:    req.AllocationAmountTxn,
		AllocationAmountBase:   req.AllocationAmountBase,
	}
	if err := s.repos.OpenItemAllocationRepo.Create(ctx, a); err != nil {
		return nil, err
	}
	created, err := s.repos.OpenItemAllocationRepo.GetByID(ctx, a.ID)
	if err != nil {
		return nil, err
	}
	return allocationToDTO(created), nil
}

func (s *openItemService) CreateAdjustment(ctx context.Context, req *dto.AdjustmentRequest) (*dto.AdjustmentResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	ad, err := parseDateTime(req.AdjustmentDate)
	if err != nil {
		return nil, err
	}
	a := &models.OpenItemAdjustment{
		OpenItemID:     req.OpenItemID,
		AdjustmentType: models.OpenItemAdjustmentType(req.AdjustmentType),
		AdjustmentDate: ad,
		AmountTxn:      req.AmountTxn,
		AmountBase:     req.AmountBase,
		ReasonCode:     req.ReasonCode,
		Notes:          req.Notes,
	}
	if err := s.repos.OpenItemAdjustmentRepo.Create(ctx, a); err != nil {
		return nil, err
	}
	created, err := s.repos.OpenItemAdjustmentRepo.GetByID(ctx, a.ID)
	if err != nil {
		return nil, err
	}
	return adjustmentToDTO(created), nil
}

func (s *openItemService) GetAllocationsByOpenItem(ctx context.Context, openItemID string) ([]*dto.AllocationResponse, error) {
	from, err := s.repos.OpenItemAllocationRepo.GetByFromOpenItemID(ctx, openItemID)
	if err != nil {
		return nil, err
	}
	to, err := s.repos.OpenItemAllocationRepo.GetByToOpenItemID(ctx, openItemID)
	if err != nil {
		return nil, err
	}
	seen := make(map[string]struct{})
	var out []*dto.AllocationResponse
	for _, a := range from {
		if _, ok := seen[a.ID]; ok {
			continue
		}
		seen[a.ID] = struct{}{}
		out = append(out, allocationToDTO(a))
	}
	for _, a := range to {
		if _, ok := seen[a.ID]; ok {
			continue
		}
		seen[a.ID] = struct{}{}
		out = append(out, allocationToDTO(a))
	}
	return out, nil
}

func (s *openItemService) GetAdjustmentsByOpenItem(ctx context.Context, openItemID string) ([]*dto.AdjustmentResponse, error) {
	list, err := s.repos.OpenItemAdjustmentRepo.GetByOpenItemID(ctx, openItemID)
	if err != nil {
		return nil, err
	}
	out := make([]*dto.AdjustmentResponse, 0, len(list))
	for _, a := range list {
		out = append(out, adjustmentToDTO(a))
	}
	return out, nil
}
