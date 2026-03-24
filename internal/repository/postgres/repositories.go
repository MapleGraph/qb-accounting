package postgres

import (
	"context"
	"errors"
	"fmt"

	"qb-accounting/internal/models"
	"qb-accounting/internal/repository"

	qbpostgres "github.com/MapleGraph/qb-core/v2/pkg/postgres"
)

func ptrSlice[T any](rows []T) []*T {
	out := make([]*T, len(rows))
	for i := range rows {
		out[i] = &rows[i]
	}
	return out
}

// --- Book ---

type bookRepository struct {
	repo qbpostgres.Repository[models.Book]
}

func NewBookRepository(db qbpostgres.DBHandler) (repository.BookRepository, error) {
	repo, err := qbpostgres.NewRepository[models.Book](db, "books")
	if err != nil {
		return nil, err
	}
	return &bookRepository{repo: repo}, nil
}

func (r *bookRepository) Create(ctx context.Context, book *models.Book) error {
	return r.repo.Create(writeContext(ctx), book)
}

func (r *bookRepository) GetByID(ctx context.Context, id string) (*models.Book, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *bookRepository) GetByCompanyID(ctx context.Context, companyID string) ([]*models.Book, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("company_id", "=", companyID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *bookRepository) Update(ctx context.Context, book *models.Book) error {
	return r.repo.Update(writeContext(ctx), book)
}

func (r *bookRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- FiscalYear ---

type fiscalYearRepository struct {
	repo qbpostgres.Repository[models.FiscalYear]
}

func NewFiscalYearRepository(db qbpostgres.DBHandler) (repository.FiscalYearRepository, error) {
	repo, err := qbpostgres.NewRepository[models.FiscalYear](db, "fiscal_years")
	if err != nil {
		return nil, err
	}
	return &fiscalYearRepository{repo: repo}, nil
}

func (r *fiscalYearRepository) Create(ctx context.Context, fy *models.FiscalYear) error {
	return r.repo.Create(writeContext(ctx), fy)
}

func (r *fiscalYearRepository) GetByID(ctx context.Context, id string) (*models.FiscalYear, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *fiscalYearRepository) GetByBookID(ctx context.Context, bookID string) ([]*models.FiscalYear, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("book_id", "=", bookID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *fiscalYearRepository) GetByCompanyID(ctx context.Context, companyID string) ([]*models.FiscalYear, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("company_id", "=", companyID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *fiscalYearRepository) Update(ctx context.Context, fy *models.FiscalYear) error {
	return r.repo.Update(writeContext(ctx), fy)
}

func (r *fiscalYearRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- AccountingPeriod ---

type accountingPeriodRepository struct {
	repo qbpostgres.Repository[models.AccountingPeriod]
}

func NewAccountingPeriodRepository(db qbpostgres.DBHandler) (repository.AccountingPeriodRepository, error) {
	repo, err := qbpostgres.NewRepository[models.AccountingPeriod](db, "accounting_periods")
	if err != nil {
		return nil, err
	}
	return &accountingPeriodRepository{repo: repo}, nil
}

func (r *accountingPeriodRepository) Create(ctx context.Context, period *models.AccountingPeriod) error {
	return r.repo.Create(writeContext(ctx), period)
}

func (r *accountingPeriodRepository) GetByID(ctx context.Context, id string) (*models.AccountingPeriod, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *accountingPeriodRepository) GetByBookID(ctx context.Context, bookID string) ([]*models.AccountingPeriod, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("book_id", "=", bookID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *accountingPeriodRepository) GetByFiscalYearID(ctx context.Context, fiscalYearID string) ([]*models.AccountingPeriod, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("fiscal_year_id", "=", fiscalYearID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *accountingPeriodRepository) GetByStatus(ctx context.Context, status string) ([]*models.AccountingPeriod, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("status", "=", status))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *accountingPeriodRepository) Update(ctx context.Context, period *models.AccountingPeriod) error {
	return r.repo.Update(writeContext(ctx), period)
}

func (r *accountingPeriodRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- AccountGroup ---

type accountGroupRepository struct {
	repo qbpostgres.Repository[models.AccountGroup]
}

func NewAccountGroupRepository(db qbpostgres.DBHandler) (repository.AccountGroupRepository, error) {
	repo, err := qbpostgres.NewRepository[models.AccountGroup](db, "account_groups")
	if err != nil {
		return nil, err
	}
	return &accountGroupRepository{repo: repo}, nil
}

func (r *accountGroupRepository) Create(ctx context.Context, group *models.AccountGroup) error {
	return r.repo.Create(writeContext(ctx), group)
}

func (r *accountGroupRepository) GetByID(ctx context.Context, id string) (*models.AccountGroup, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *accountGroupRepository) GetByBookID(ctx context.Context, bookID string) ([]*models.AccountGroup, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("book_id", "=", bookID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *accountGroupRepository) GetByParentGroupID(ctx context.Context, parentGroupID string) ([]*models.AccountGroup, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("parent_group_id", "=", parentGroupID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *accountGroupRepository) Update(ctx context.Context, group *models.AccountGroup) error {
	return r.repo.Update(writeContext(ctx), group)
}

func (r *accountGroupRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- Account ---

type accountRepository struct {
	repo qbpostgres.Repository[models.Account]
}

func NewAccountRepository(db qbpostgres.DBHandler) (repository.AccountRepository, error) {
	repo, err := qbpostgres.NewRepository[models.Account](db, "accounts")
	if err != nil {
		return nil, err
	}
	return &accountRepository{repo: repo}, nil
}

func (r *accountRepository) Create(ctx context.Context, account *models.Account) error {
	return r.repo.Create(writeContext(ctx), account)
}

func (r *accountRepository) GetByID(ctx context.Context, id string) (*models.Account, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *accountRepository) GetByBookID(ctx context.Context, bookID string) ([]*models.Account, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("book_id", "=", bookID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *accountRepository) GetByCompanyID(ctx context.Context, companyID string) ([]*models.Account, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("company_id", "=", companyID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *accountRepository) GetByGroupID(ctx context.Context, groupID string) ([]*models.Account, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("group_id", "=", groupID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *accountRepository) GetByControlType(ctx context.Context, bookID string, controlType string) ([]*models.Account, error) {
	rows, err := r.repo.List(readContext(ctx),
		qbpostgres.WithFilter("book_id", "=", bookID),
		qbpostgres.WithFilter("control_type", "=", controlType),
	)
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *accountRepository) Update(ctx context.Context, account *models.Account) error {
	return r.repo.Update(writeContext(ctx), account)
}

func (r *accountRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- VoucherSequence ---

type voucherSequenceRepository struct {
	repo qbpostgres.Repository[models.VoucherSequence]
	db   qbpostgres.DBHandler
}

func NewVoucherSequenceRepository(db qbpostgres.DBHandler) (repository.VoucherSequenceRepository, error) {
	repo, err := qbpostgres.NewRepository[models.VoucherSequence](db, "voucher_sequences")
	if err != nil {
		return nil, err
	}
	return &voucherSequenceRepository{repo: repo, db: db}, nil
}

func (r *voucherSequenceRepository) Create(ctx context.Context, seq *models.VoucherSequence) error {
	return r.repo.Create(writeContext(ctx), seq)
}

func (r *voucherSequenceRepository) GetByID(ctx context.Context, id string) (*models.VoucherSequence, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *voucherSequenceRepository) GetByBookIDAndVoucherType(ctx context.Context, bookID string, voucherType string) (*models.VoucherSequence, error) {
	row, err := r.repo.Get(readContext(ctx),
		qbpostgres.WithFilter("book_id", "=", bookID),
		qbpostgres.WithFilter("voucher_type", "=", voucherType),
	)
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

// GetNextSequence atomically increments next_number and returns the sequence value assigned for this call
// (same pattern as qb-billing transaction series; qb-core UpdatePartial cannot express column increments).
func (r *voucherSequenceRepository) GetNextSequence(ctx context.Context, id string) (int64, error) {
	query := `
		UPDATE voucher_sequences
		SET next_number = next_number + 1
		WHERE id = $1
		RETURNING next_number - 1
	`
	var seq int64
	writeCtx := writeContext(ctx)
	conn, err := r.db.GetConnection(writeCtx)
	if err != nil {
		return 0, fmt.Errorf("failed to get write connection: %w", err)
	}
	if err := conn.QueryRowContext(writeCtx, query, id).Scan(&seq); err != nil {
		return 0, fmt.Errorf("failed to get next sequence: %w", err)
	}
	return seq, nil
}

func (r *voucherSequenceRepository) Update(ctx context.Context, seq *models.VoucherSequence) error {
	return r.repo.Update(writeContext(ctx), seq)
}

func (r *voucherSequenceRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- JournalBatch ---

type journalBatchRepository struct {
	repo qbpostgres.Repository[models.JournalBatch]
}

func NewJournalBatchRepository(db qbpostgres.DBHandler) (repository.JournalBatchRepository, error) {
	repo, err := qbpostgres.NewRepository[models.JournalBatch](db, "journal_batches")
	if err != nil {
		return nil, err
	}
	return &journalBatchRepository{repo: repo}, nil
}

func (r *journalBatchRepository) Create(ctx context.Context, batch *models.JournalBatch) error {
	return r.repo.Create(writeContext(ctx), batch)
}

func (r *journalBatchRepository) GetByID(ctx context.Context, id string) (*models.JournalBatch, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *journalBatchRepository) GetByBookID(ctx context.Context, bookID string) ([]*models.JournalBatch, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("book_id", "=", bookID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *journalBatchRepository) GetByStatus(ctx context.Context, status string) ([]*models.JournalBatch, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("status", "=", status))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *journalBatchRepository) Update(ctx context.Context, batch *models.JournalBatch) error {
	return r.repo.Update(writeContext(ctx), batch)
}

func (r *journalBatchRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- Journal ---

type journalRepository struct {
	repo qbpostgres.Repository[models.Journal]
}

func NewJournalRepository(db qbpostgres.DBHandler) (repository.JournalRepository, error) {
	repo, err := qbpostgres.NewRepository[models.Journal](db, "journals")
	if err != nil {
		return nil, err
	}
	return &journalRepository{repo: repo}, nil
}

func (r *journalRepository) Create(ctx context.Context, journal *models.Journal) error {
	return r.repo.Create(writeContext(ctx), journal)
}

func (r *journalRepository) GetByID(ctx context.Context, id string) (*models.Journal, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *journalRepository) GetByBookID(ctx context.Context, bookID string) ([]*models.Journal, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("book_id", "=", bookID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *journalRepository) GetByPeriodID(ctx context.Context, periodID string) ([]*models.Journal, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("period_id", "=", periodID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *journalRepository) GetBySourceDocument(ctx context.Context, sourceModule, sourceDocType string, sourceDocID string) (*models.Journal, error) {
	row, err := r.repo.Get(readContext(ctx),
		qbpostgres.WithFilter("source_module", "=", sourceModule),
		qbpostgres.WithFilter("source_document_type", "=", sourceDocType),
		qbpostgres.WithFilter("source_document_id", "=", sourceDocID),
	)
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *journalRepository) GetByIdempotencyKey(ctx context.Context, bookID, key string) (*models.Journal, error) {
	row, err := r.repo.Get(readContext(ctx),
		qbpostgres.WithFilter("book_id", "=", bookID),
		qbpostgres.WithFilter("idempotency_key", "=", key),
	)
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *journalRepository) Update(ctx context.Context, journal *models.Journal) error {
	return r.repo.Update(writeContext(ctx), journal)
}

func (r *journalRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- JournalLine ---

type journalLineRepository struct {
	repo qbpostgres.Repository[models.JournalLine]
}

func NewJournalLineRepository(db qbpostgres.DBHandler) (repository.JournalLineRepository, error) {
	repo, err := qbpostgres.NewRepository[models.JournalLine](db, "journal_lines")
	if err != nil {
		return nil, err
	}
	return &journalLineRepository{repo: repo}, nil
}

func (r *journalLineRepository) Create(ctx context.Context, line *models.JournalLine) error {
	return r.repo.Create(writeContext(ctx), line)
}

func (r *journalLineRepository) GetByID(ctx context.Context, id string) (*models.JournalLine, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *journalLineRepository) GetByJournalID(ctx context.Context, journalID string) ([]*models.JournalLine, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("journal_id", "=", journalID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *journalLineRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- PostingRuleVersion ---

type postingRuleVersionRepository struct {
	repo qbpostgres.Repository[models.PostingRuleVersion]
}

func NewPostingRuleVersionRepository(db qbpostgres.DBHandler) (repository.PostingRuleVersionRepository, error) {
	repo, err := qbpostgres.NewRepository[models.PostingRuleVersion](db, "posting_rule_versions")
	if err != nil {
		return nil, err
	}
	return &postingRuleVersionRepository{repo: repo}, nil
}

func (r *postingRuleVersionRepository) Create(ctx context.Context, rule *models.PostingRuleVersion) error {
	return r.repo.Create(writeContext(ctx), rule)
}

func (r *postingRuleVersionRepository) GetByID(ctx context.Context, id string) (*models.PostingRuleVersion, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *postingRuleVersionRepository) GetActiveBySourceDocType(ctx context.Context, sourceModule, sourceDocType string) (*models.PostingRuleVersion, error) {
	row, err := r.repo.Get(readContext(ctx),
		qbpostgres.WithFilter("source_module", "=", sourceModule),
		qbpostgres.WithFilter("source_document_type", "=", sourceDocType),
		qbpostgres.WithFilter("status", "=", string(models.PostingRuleStatusActive)),
		qbpostgres.WithSort("version_no", "DESC"),
	)
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *postingRuleVersionRepository) Update(ctx context.Context, rule *models.PostingRuleVersion) error {
	return r.repo.Update(writeContext(ctx), rule)
}

func (r *postingRuleVersionRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- PostingRequest ---

type postingRequestRepository struct {
	repo qbpostgres.Repository[models.PostingRequest]
}

func NewPostingRequestRepository(db qbpostgres.DBHandler) (repository.PostingRequestRepository, error) {
	repo, err := qbpostgres.NewRepository[models.PostingRequest](db, "posting_requests")
	if err != nil {
		return nil, err
	}
	return &postingRequestRepository{repo: repo}, nil
}

func (r *postingRequestRepository) Create(ctx context.Context, req *models.PostingRequest) error {
	return r.repo.Create(writeContext(ctx), req)
}

func (r *postingRequestRepository) GetByID(ctx context.Context, id string) (*models.PostingRequest, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *postingRequestRepository) GetByIdempotencyKey(ctx context.Context, bookID, key string) (*models.PostingRequest, error) {
	row, err := r.repo.Get(readContext(ctx),
		qbpostgres.WithFilter("book_id", "=", bookID),
		qbpostgres.WithFilter("idempotency_key", "=", key),
	)
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *postingRequestRepository) GetBySourceDocument(ctx context.Context, sourceModule, sourceDocType string, sourceDocID string) (*models.PostingRequest, error) {
	row, err := r.repo.Get(readContext(ctx),
		qbpostgres.WithFilter("source_module", "=", sourceModule),
		qbpostgres.WithFilter("source_document_type", "=", sourceDocType),
		qbpostgres.WithFilter("source_document_id", "=", sourceDocID),
	)
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *postingRequestRepository) GetByStatus(ctx context.Context, status string) ([]*models.PostingRequest, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("request_status", "=", status))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *postingRequestRepository) Update(ctx context.Context, req *models.PostingRequest) error {
	return r.repo.Update(writeContext(ctx), req)
}

func (r *postingRequestRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- PostingRequestSnapshot ---

type postingRequestSnapshotRepository struct {
	repo qbpostgres.Repository[models.PostingRequestSnapshot]
}

func NewPostingRequestSnapshotRepository(db qbpostgres.DBHandler) (repository.PostingRequestSnapshotRepository, error) {
	repo, err := qbpostgres.NewRepository[models.PostingRequestSnapshot](db, "posting_request_snapshots")
	if err != nil {
		return nil, err
	}
	return &postingRequestSnapshotRepository{repo: repo}, nil
}

func (r *postingRequestSnapshotRepository) Create(ctx context.Context, snap *models.PostingRequestSnapshot) error {
	return r.repo.Create(writeContext(ctx), snap)
}

func (r *postingRequestSnapshotRepository) GetByID(ctx context.Context, id string) (*models.PostingRequestSnapshot, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *postingRequestSnapshotRepository) GetByPostingRequestID(ctx context.Context, postingRequestID string) ([]*models.PostingRequestSnapshot, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("posting_request_id", "=", postingRequestID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

// --- JournalSourceLink ---

type journalSourceLinkRepository struct {
	repo qbpostgres.Repository[models.JournalSourceLink]
}

func NewJournalSourceLinkRepository(db qbpostgres.DBHandler) (repository.JournalSourceLinkRepository, error) {
	repo, err := qbpostgres.NewRepository[models.JournalSourceLink](db, "journal_source_links")
	if err != nil {
		return nil, err
	}
	return &journalSourceLinkRepository{repo: repo}, nil
}

func (r *journalSourceLinkRepository) Create(ctx context.Context, link *models.JournalSourceLink) error {
	return r.repo.Create(writeContext(ctx), link)
}

func (r *journalSourceLinkRepository) GetByID(ctx context.Context, id string) (*models.JournalSourceLink, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *journalSourceLinkRepository) GetByPostingRequestID(ctx context.Context, postingRequestID string) ([]*models.JournalSourceLink, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("posting_request_id", "=", postingRequestID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *journalSourceLinkRepository) GetByJournalID(ctx context.Context, journalID string) ([]*models.JournalSourceLink, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("journal_id", "=", journalID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

// --- OpenItem ---

type openItemRepository struct {
	repo qbpostgres.Repository[models.OpenItem]
}

func NewOpenItemRepository(db qbpostgres.DBHandler) (repository.OpenItemRepository, error) {
	repo, err := qbpostgres.NewRepository[models.OpenItem](db, "open_items")
	if err != nil {
		return nil, err
	}
	return &openItemRepository{repo: repo}, nil
}

func (r *openItemRepository) Create(ctx context.Context, item *models.OpenItem) error {
	return r.repo.Create(writeContext(ctx), item)
}

func (r *openItemRepository) GetByID(ctx context.Context, id string) (*models.OpenItem, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *openItemRepository) GetByPartyID(ctx context.Context, partyID string) ([]*models.OpenItem, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("party_id", "=", partyID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *openItemRepository) GetByBookID(ctx context.Context, bookID string) ([]*models.OpenItem, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("book_id", "=", bookID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *openItemRepository) GetByStatus(ctx context.Context, bookID string, side string, status string) ([]*models.OpenItem, error) {
	rows, err := r.repo.List(readContext(ctx),
		qbpostgres.WithFilter("book_id", "=", bookID),
		qbpostgres.WithFilter("item_side", "=", side),
		qbpostgres.WithFilter("item_status", "=", status),
	)
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *openItemRepository) Update(ctx context.Context, item *models.OpenItem) error {
	return r.repo.Update(writeContext(ctx), item)
}

func (r *openItemRepository) Delete(ctx context.Context, id string) error {
	return r.repo.Delete(writeContext(ctx), id)
}

// --- OpenItemAllocation ---

type openItemAllocationRepository struct {
	repo qbpostgres.Repository[models.OpenItemAllocation]
}

func NewOpenItemAllocationRepository(db qbpostgres.DBHandler) (repository.OpenItemAllocationRepository, error) {
	repo, err := qbpostgres.NewRepository[models.OpenItemAllocation](db, "open_item_allocations")
	if err != nil {
		return nil, err
	}
	return &openItemAllocationRepository{repo: repo}, nil
}

func (r *openItemAllocationRepository) Create(ctx context.Context, alloc *models.OpenItemAllocation) error {
	return r.repo.Create(writeContext(ctx), alloc)
}

func (r *openItemAllocationRepository) GetByID(ctx context.Context, id string) (*models.OpenItemAllocation, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *openItemAllocationRepository) GetByFromOpenItemID(ctx context.Context, fromOpenItemID string) ([]*models.OpenItemAllocation, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("from_open_item_id", "=", fromOpenItemID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

func (r *openItemAllocationRepository) GetByToOpenItemID(ctx context.Context, toOpenItemID string) ([]*models.OpenItemAllocation, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("to_open_item_id", "=", toOpenItemID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}

// --- OpenItemAdjustment ---

type openItemAdjustmentRepository struct {
	repo qbpostgres.Repository[models.OpenItemAdjustment]
}

func NewOpenItemAdjustmentRepository(db qbpostgres.DBHandler) (repository.OpenItemAdjustmentRepository, error) {
	repo, err := qbpostgres.NewRepository[models.OpenItemAdjustment](db, "open_item_adjustments")
	if err != nil {
		return nil, err
	}
	return &openItemAdjustmentRepository{repo: repo}, nil
}

func (r *openItemAdjustmentRepository) Create(ctx context.Context, adj *models.OpenItemAdjustment) error {
	return r.repo.Create(writeContext(ctx), adj)
}

func (r *openItemAdjustmentRepository) GetByID(ctx context.Context, id string) (*models.OpenItemAdjustment, error) {
	row, err := r.repo.Get(readContext(ctx), qbpostgres.WithFilter("id", "=", id))
	if err != nil {
		if errors.Is(err, qbpostgres.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return row, nil
}

func (r *openItemAdjustmentRepository) GetByOpenItemID(ctx context.Context, openItemID string) ([]*models.OpenItemAdjustment, error) {
	rows, err := r.repo.List(readContext(ctx), qbpostgres.WithFilter("open_item_id", "=", openItemID))
	if err != nil {
		return nil, err
	}
	return ptrSlice(rows), nil
}
