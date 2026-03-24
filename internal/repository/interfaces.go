package repository

import (
	"context"

	"qb-accounting/internal/models"
)

// BookRepository defines operations for accounting books.
type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
	GetByID(ctx context.Context, id string) (*models.Book, error)
	GetByCompanyID(ctx context.Context, companyID string) ([]*models.Book, error)
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id string) error
}

// FiscalYearRepository defines operations for fiscal years.
type FiscalYearRepository interface {
	Create(ctx context.Context, fy *models.FiscalYear) error
	GetByID(ctx context.Context, id string) (*models.FiscalYear, error)
	GetByBookID(ctx context.Context, bookID string) ([]*models.FiscalYear, error)
	GetByCompanyID(ctx context.Context, companyID string) ([]*models.FiscalYear, error)
	Update(ctx context.Context, fy *models.FiscalYear) error
	Delete(ctx context.Context, id string) error
}

// AccountingPeriodRepository defines operations for accounting periods.
type AccountingPeriodRepository interface {
	Create(ctx context.Context, period *models.AccountingPeriod) error
	GetByID(ctx context.Context, id string) (*models.AccountingPeriod, error)
	GetByBookID(ctx context.Context, bookID string) ([]*models.AccountingPeriod, error)
	GetByFiscalYearID(ctx context.Context, fiscalYearID string) ([]*models.AccountingPeriod, error)
	GetByStatus(ctx context.Context, status string) ([]*models.AccountingPeriod, error)
	Update(ctx context.Context, period *models.AccountingPeriod) error
	Delete(ctx context.Context, id string) error
}

// AccountGroupRepository defines operations for account groups.
type AccountGroupRepository interface {
	Create(ctx context.Context, group *models.AccountGroup) error
	GetByID(ctx context.Context, id string) (*models.AccountGroup, error)
	GetByBookID(ctx context.Context, bookID string) ([]*models.AccountGroup, error)
	GetByParentGroupID(ctx context.Context, parentGroupID string) ([]*models.AccountGroup, error)
	Update(ctx context.Context, group *models.AccountGroup) error
	Delete(ctx context.Context, id string) error
}

// AccountRepository defines operations for accounts.
type AccountRepository interface {
	Create(ctx context.Context, account *models.Account) error
	GetByID(ctx context.Context, id string) (*models.Account, error)
	GetByBookID(ctx context.Context, bookID string) ([]*models.Account, error)
	GetByCompanyID(ctx context.Context, companyID string) ([]*models.Account, error)
	GetByGroupID(ctx context.Context, groupID string) ([]*models.Account, error)
	GetByControlType(ctx context.Context, bookID string, controlType string) ([]*models.Account, error)
	Update(ctx context.Context, account *models.Account) error
	Delete(ctx context.Context, id string) error
}

// VoucherSequenceRepository defines operations for voucher sequences.
type VoucherSequenceRepository interface {
	Create(ctx context.Context, seq *models.VoucherSequence) error
	GetByID(ctx context.Context, id string) (*models.VoucherSequence, error)
	GetByBookIDAndVoucherType(ctx context.Context, bookID string, voucherType string) (*models.VoucherSequence, error)
	GetNextSequence(ctx context.Context, id string) (int64, error)
	Update(ctx context.Context, seq *models.VoucherSequence) error
	Delete(ctx context.Context, id string) error
}

// JournalBatchRepository defines operations for journal batches.
type JournalBatchRepository interface {
	Create(ctx context.Context, batch *models.JournalBatch) error
	GetByID(ctx context.Context, id string) (*models.JournalBatch, error)
	GetByBookID(ctx context.Context, bookID string) ([]*models.JournalBatch, error)
	GetByStatus(ctx context.Context, status string) ([]*models.JournalBatch, error)
	Update(ctx context.Context, batch *models.JournalBatch) error
	Delete(ctx context.Context, id string) error
}

// JournalRepository defines operations for journals.
type JournalRepository interface {
	Create(ctx context.Context, journal *models.Journal) error
	GetByID(ctx context.Context, id string) (*models.Journal, error)
	GetByBookID(ctx context.Context, bookID string) ([]*models.Journal, error)
	GetByPeriodID(ctx context.Context, periodID string) ([]*models.Journal, error)
	GetBySourceDocument(ctx context.Context, sourceModule, sourceDocType string, sourceDocID string) (*models.Journal, error)
	GetByIdempotencyKey(ctx context.Context, bookID, key string) (*models.Journal, error)
	Update(ctx context.Context, journal *models.Journal) error
	Delete(ctx context.Context, id string) error
}

// JournalLineRepository defines operations for journal lines.
type JournalLineRepository interface {
	Create(ctx context.Context, line *models.JournalLine) error
	GetByID(ctx context.Context, id string) (*models.JournalLine, error)
	GetByJournalID(ctx context.Context, journalID string) ([]*models.JournalLine, error)
	Delete(ctx context.Context, id string) error
}

// PostingRuleVersionRepository defines operations for posting rule versions.
type PostingRuleVersionRepository interface {
	Create(ctx context.Context, rule *models.PostingRuleVersion) error
	GetByID(ctx context.Context, id string) (*models.PostingRuleVersion, error)
	GetActiveBySourceDocType(ctx context.Context, sourceModule, sourceDocType string) (*models.PostingRuleVersion, error)
	Update(ctx context.Context, rule *models.PostingRuleVersion) error
	Delete(ctx context.Context, id string) error
}

// PostingRequestRepository defines operations for posting requests.
type PostingRequestRepository interface {
	Create(ctx context.Context, req *models.PostingRequest) error
	GetByID(ctx context.Context, id string) (*models.PostingRequest, error)
	GetByIdempotencyKey(ctx context.Context, bookID, key string) (*models.PostingRequest, error)
	GetBySourceDocument(ctx context.Context, sourceModule, sourceDocType string, sourceDocID string) (*models.PostingRequest, error)
	GetByStatus(ctx context.Context, status string) ([]*models.PostingRequest, error)
	Update(ctx context.Context, req *models.PostingRequest) error
	Delete(ctx context.Context, id string) error
}

// PostingRequestSnapshotRepository defines operations for posting request snapshots.
type PostingRequestSnapshotRepository interface {
	Create(ctx context.Context, snap *models.PostingRequestSnapshot) error
	GetByID(ctx context.Context, id string) (*models.PostingRequestSnapshot, error)
	GetByPostingRequestID(ctx context.Context, postingRequestID string) ([]*models.PostingRequestSnapshot, error)
}

// JournalSourceLinkRepository defines operations for journal source links.
type JournalSourceLinkRepository interface {
	Create(ctx context.Context, link *models.JournalSourceLink) error
	GetByID(ctx context.Context, id string) (*models.JournalSourceLink, error)
	GetByPostingRequestID(ctx context.Context, postingRequestID string) ([]*models.JournalSourceLink, error)
	GetByJournalID(ctx context.Context, journalID string) ([]*models.JournalSourceLink, error)
}

// OpenItemRepository defines operations for open items.
type OpenItemRepository interface {
	Create(ctx context.Context, item *models.OpenItem) error
	GetByID(ctx context.Context, id string) (*models.OpenItem, error)
	GetByPartyID(ctx context.Context, partyID string) ([]*models.OpenItem, error)
	GetByBookID(ctx context.Context, bookID string) ([]*models.OpenItem, error)
	GetByStatus(ctx context.Context, bookID string, side string, status string) ([]*models.OpenItem, error)
	Update(ctx context.Context, item *models.OpenItem) error
	Delete(ctx context.Context, id string) error
}

// OpenItemAllocationRepository defines operations for open item allocations.
type OpenItemAllocationRepository interface {
	Create(ctx context.Context, alloc *models.OpenItemAllocation) error
	GetByID(ctx context.Context, id string) (*models.OpenItemAllocation, error)
	GetByFromOpenItemID(ctx context.Context, fromOpenItemID string) ([]*models.OpenItemAllocation, error)
	GetByToOpenItemID(ctx context.Context, toOpenItemID string) ([]*models.OpenItemAllocation, error)
}

// OpenItemAdjustmentRepository defines operations for open item adjustments.
type OpenItemAdjustmentRepository interface {
	Create(ctx context.Context, adj *models.OpenItemAdjustment) error
	GetByID(ctx context.Context, id string) (*models.OpenItemAdjustment, error)
	GetByOpenItemID(ctx context.Context, openItemID string) ([]*models.OpenItemAdjustment, error)
}
