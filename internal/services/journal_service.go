package services

import (
	"context"
	"fmt"
	"time"

	"qb-accounting/internal/dto"
	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"
	"qb-accounting/internal/utils"

	qbpostgres "github.com/MapleGraph/qb-core/v2/pkg/postgres"
)

const journalVoucherTypeJV = "JV"

// JournalService is the HTTP-facing API for journals (DTO responses).
type JournalService interface {
	CreateJournal(ctx context.Context, req *dto.CreateJournalRequest) (*dto.JournalResponse, error)
	GetJournal(ctx context.Context, id string, includeLines bool) (*dto.JournalResponse, error)
	ListJournalsByBook(ctx context.Context, bookID string) ([]*dto.JournalResponse, error)
	ListJournalsByPeriod(ctx context.Context, periodID string) ([]*dto.JournalResponse, error)
	PostJournal(ctx context.Context, req *dto.PostJournalRequest) (*dto.JournalResponse, error)
	ReverseJournal(ctx context.Context, req *dto.ReverseJournalRequest) (*dto.JournalResponse, error)
	DeleteJournal(ctx context.Context, id string) error
}

type journalService struct {
	repos *repository.RepositoryContainer
}

func NewJournalService(repos *repository.RepositoryContainer) JournalService {
	return &journalService{repos: repos}
}

func (s *journalService) getJournalDTO(ctx context.Context, id string, includeLines bool) (*dto.JournalResponse, error) {
	if includeLines {
		j, lines, err := s.getJournalWithLines(ctx, id)
		if err != nil {
			return nil, err
		}
		return journalModelToDTO(j, lines), nil
	}
	j, err := s.repos.JournalRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return journalModelToDTO(j, nil), nil
}

func formatJournalNumber(seq *models.VoucherSequence, n int64) string {
	prefix := ""
	if seq.Prefix != nil {
		prefix = *seq.Prefix
	}
	suffix := ""
	if seq.Suffix != nil {
		suffix = *seq.Suffix
	}
	num := fmt.Sprintf("%0*d", seq.Padding, n)
	return prefix + num + suffix
}

func journalLineTotals(lines []dto.JournalLineRequest) (debit, credit float64) {
	for _, ln := range lines {
		debit += ln.DebitAmountTxn
		credit += ln.CreditAmountTxn
	}
	return debit, credit
}

func (s *journalService) CreateJournal(ctx context.Context, req *dto.CreateJournalRequest) (*dto.JournalResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	dr, cr := journalLineTotals(req.Lines)
	if dr != cr {
		return nil, fmt.Errorf("journal lines not balanced: debit %v != credit %v", dr, cr)
	}
	journalDate, err := parseDateTime(req.JournalDate)
	if err != nil {
		return nil, err
	}
	postingDate, err := parseDateTime(req.PostingDate)
	if err != nil {
		return nil, err
	}

	var created *models.Journal
	err = qbpostgres.WithTransaction(ctx, s.repos.DB, func(txCtx context.Context) error {
		seq, err := s.repos.VoucherSequenceRepo.GetByBookIDAndVoucherType(txCtx, req.BookID, journalVoucherTypeJV)
		if err != nil {
			return err
		}
		if seq == nil {
			return fmt.Errorf("voucher sequence for type %s not found for book", journalVoucherTypeJV)
		}
		if !seq.IsActive {
			return fmt.Errorf("voucher sequence %s is inactive", seq.ID)
		}
		nextNo, err := s.repos.VoucherSequenceRepo.GetNextSequence(txCtx, seq.ID)
		if err != nil {
			return err
		}
		journalNo := formatJournalNumber(seq, nextNo)

		j := &models.Journal{
			BookID:             req.BookID,
			FiscalYearID:       req.FiscalYearID,
			PeriodID:           req.PeriodID,
			CompanyID:          req.CompanyID,
			BranchID:           req.BranchID,
			JournalNo:          journalNo,
			JournalKind:        models.JournalKind(req.JournalKind),
			Status:             models.JournalStatusDraft,
			SourceModule:       req.SourceModule,
			SourceDocumentType: req.SourceDocumentType,
			SourceDocumentID:   req.SourceDocumentID,
			SourceEventID:      req.SourceEventID,
			IdempotencyKey:     req.IdempotencyKey,
			JournalDate:        journalDate,
			PostingDate:        postingDate,
			CurrencyCode:       req.CurrencyCode,
			ExchangeRate:       req.ExchangeRate,
			ReferenceNo:        req.ReferenceNo,
			Narration:          req.Narration,
			Metadata:           utils.JSONB{},
		}
		if err := s.repos.JournalRepo.Create(txCtx, j); err != nil {
			return err
		}
		for i, line := range req.Lines {
			jl := &models.JournalLine{
				JournalID:        j.ID,
				LineNo:           i + 1,
				AccountID:        line.AccountID,
				CompanyID:        req.CompanyID,
				BranchID:         line.BranchID,
				PartyID:          line.PartyID,
				EmployeeID:       line.EmployeeID,
				CostCenterID:     line.CostCenterID,
				IncomeHeadID:     line.IncomeHeadID,
				TaxCodeID:        line.TaxCodeID,
				Description:      line.Description,
				DebitAmountTxn:   line.DebitAmountTxn,
				CreditAmountTxn:  line.CreditAmountTxn,
				DebitAmountBase:  line.DebitAmountBase,
				CreditAmountBase: line.CreditAmountBase,
				ExchangeRate:     line.ExchangeRate,
				LineMetadata:     utils.JSONB{},
			}
			if err := s.repos.JournalLineRepo.Create(txCtx, jl); err != nil {
				return err
			}
		}
		var getErr error
		created, getErr = s.repos.JournalRepo.GetByID(txCtx, j.ID)
		return getErr
	})
	if err != nil {
		return nil, err
	}
	return s.getJournalDTO(ctx, created.ID, true)
}

func (s *journalService) GetJournal(ctx context.Context, id string, includeLines bool) (*dto.JournalResponse, error) {
	return s.getJournalDTO(ctx, id, includeLines)
}

func (s *journalService) getJournalWithLines(ctx context.Context, id string) (*models.Journal, []*models.JournalLine, error) {
	j, err := s.repos.JournalRepo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	if j == nil {
		return nil, nil, nil
	}
	lines, err := s.repos.JournalLineRepo.GetByJournalID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	return j, lines, nil
}

func (s *journalService) ListJournalsByBook(ctx context.Context, bookID string) ([]*dto.JournalResponse, error) {
	list, err := s.repos.JournalRepo.GetByBookID(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return journalsModelToDTO(list), nil
}

func (s *journalService) ListJournalsByPeriod(ctx context.Context, periodID string) ([]*dto.JournalResponse, error) {
	list, err := s.repos.JournalRepo.GetByPeriodID(ctx, periodID)
	if err != nil {
		return nil, err
	}
	return journalsModelToDTO(list), nil
}

func (s *journalService) PostJournal(ctx context.Context, req *dto.PostJournalRequest) (*dto.JournalResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	j, err := s.repos.JournalRepo.GetByID(ctx, req.JournalID)
	if err != nil {
		return nil, err
	}
	if j == nil {
		return nil, fmt.Errorf("journal not found")
	}
	if j.Status != models.JournalStatusDraft {
		return nil, fmt.Errorf("journal must be in DRAFT status to post")
	}
	now := time.Now()
	j.Status = models.JournalStatusPosted
	j.PostedAt = &now
	if req.PostingDate != nil {
		pd, err := parseDateTime(*req.PostingDate)
		if err != nil {
			return nil, err
		}
		j.PostingDate = pd
	}
	if err := s.repos.JournalRepo.Update(ctx, j); err != nil {
		return nil, err
	}
	return s.getJournalDTO(ctx, j.ID, false)
}

func (s *journalService) ReverseJournal(ctx context.Context, req *dto.ReverseJournalRequest) (*dto.JournalResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}
	var reversal *models.Journal
	err := qbpostgres.WithTransaction(ctx, s.repos.DB, func(txCtx context.Context) error {
		orig, err := s.repos.JournalRepo.GetByID(txCtx, req.JournalID)
		if err != nil {
			return err
		}
		if orig == nil {
			return fmt.Errorf("journal not found")
		}
		if orig.Status != models.JournalStatusPosted {
			return fmt.Errorf("only posted journals can be reversed")
		}
		origLines, err := s.repos.JournalLineRepo.GetByJournalID(txCtx, orig.ID)
		if err != nil {
			return err
		}
		seq, err := s.repos.VoucherSequenceRepo.GetByBookIDAndVoucherType(txCtx, orig.BookID, journalVoucherTypeJV)
		if err != nil {
			return err
		}
		if seq == nil {
			return fmt.Errorf("voucher sequence for type %s not found for book", journalVoucherTypeJV)
		}
		nextNo, err := s.repos.VoucherSequenceRepo.GetNextSequence(txCtx, seq.ID)
		if err != nil {
			return err
		}
		journalNo := formatJournalNumber(seq, nextNo)
		revDate := time.Now()
		if req.ReversalDate != nil {
			revDate, err = parseDateTime(*req.ReversalDate)
			if err != nil {
				return err
			}
		}
		narration := req.Narration
		if narration == nil {
			n := fmt.Sprintf("Reversal of journal %s", orig.JournalNo)
			narration = &n
		}
		now := time.Now()
		rev := &models.Journal{
			BookID:              orig.BookID,
			FiscalYearID:        orig.FiscalYearID,
			PeriodID:            orig.PeriodID,
			CompanyID:           orig.CompanyID,
			BranchID:            orig.BranchID,
			JournalNo:           journalNo,
			JournalKind:         models.JournalKindReversal,
			Status:              models.JournalStatusPosted,
			SourceModule:        orig.SourceModule,
			SourceDocumentType:  orig.SourceDocumentType,
			SourceDocumentID:    orig.SourceDocumentID,
			SourceEventID:       orig.SourceEventID,
			JournalDate:         revDate,
			PostingDate:         revDate,
			CurrencyCode:        orig.CurrencyCode,
			ExchangeRate:        orig.ExchangeRate,
			Narration:           narration,
			ReversalOfJournalID: &orig.ID,
			Metadata:            utils.JSONB{},
			PostedAt:            &now,
		}
		if err := s.repos.JournalRepo.Create(txCtx, rev); err != nil {
			return err
		}
		for i, ln := range origLines {
			jl := &models.JournalLine{
				JournalID:        rev.ID,
				LineNo:           i + 1,
				AccountID:        ln.AccountID,
				CompanyID:        ln.CompanyID,
				BranchID:         ln.BranchID,
				PartyID:          ln.PartyID,
				EmployeeID:       ln.EmployeeID,
				CostCenterID:     ln.CostCenterID,
				IncomeHeadID:     ln.IncomeHeadID,
				TaxCodeID:        ln.TaxCodeID,
				Description:      ln.Description,
				DebitAmountTxn:   ln.CreditAmountTxn,
				CreditAmountTxn:  ln.DebitAmountTxn,
				DebitAmountBase:  ln.CreditAmountBase,
				CreditAmountBase: ln.DebitAmountBase,
				ExchangeRate:     ln.ExchangeRate,
				LineMetadata:     utils.JSONB{},
			}
			if err := s.repos.JournalLineRepo.Create(txCtx, jl); err != nil {
				return err
			}
		}
		orig.Status = models.JournalStatusReversed
		orig.ReversedByJournalID = &rev.ID
		orig.ReversedAt = &now
		if err := s.repos.JournalRepo.Update(txCtx, orig); err != nil {
			return err
		}
		var getErr error
		reversal, getErr = s.repos.JournalRepo.GetByID(txCtx, rev.ID)
		return getErr
	})
	if err != nil {
		return nil, err
	}
	return s.getJournalDTO(ctx, reversal.ID, true)
}

func (s *journalService) DeleteJournal(ctx context.Context, id string) error {
	j, err := s.repos.JournalRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if j == nil {
		return fmt.Errorf("journal not found")
	}
	if j.Status != models.JournalStatusDraft {
		return fmt.Errorf("only draft journals can be deleted")
	}
	lines, err := s.repos.JournalLineRepo.GetByJournalID(ctx, id)
	if err != nil {
		return err
	}
	for _, ln := range lines {
		if err := s.repos.JournalLineRepo.Delete(ctx, ln.ID); err != nil {
			return err
		}
	}
	return s.repos.JournalRepo.Delete(ctx, id)
}
