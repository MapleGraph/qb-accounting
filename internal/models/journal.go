package models

import (
	"time"

	"qb-accounting/internal/utils"
)

type JournalKind string

const (
	JournalKindManual     JournalKind = "MANUAL"
	JournalKindOpening    JournalKind = "OPENING"
	JournalKindAdjustment JournalKind = "ADJUSTMENT"
	JournalKindReversal   JournalKind = "REVERSAL"
	JournalKindSystem     JournalKind = "SYSTEM"
	JournalKindClosing    JournalKind = "CLOSING"
)

type JournalStatus string

const (
	JournalStatusDraft     JournalStatus = "DRAFT"
	JournalStatusPosted    JournalStatus = "POSTED"
	JournalStatusReversed  JournalStatus = "REVERSED"
	JournalStatusCancelled JournalStatus = "CANCELLED"
)

type Journal struct {
	ID                   string        `json:"id" db:"id" db_pk:"true"`
	BookID               string        `json:"book_id" db:"book_id"`
	BatchID              *string       `json:"batch_id" db:"batch_id"`
	FiscalYearID         string        `json:"fiscal_year_id" db:"fiscal_year_id"`
	PeriodID             string        `json:"period_id" db:"period_id"`
	CompanyID            string        `json:"company_id" db:"company_id"`
	BranchID             *string       `json:"branch_id" db:"branch_id"`
	JournalNo            string        `json:"journal_no" db:"journal_no"`
	JournalKind          JournalKind   `json:"journal_kind" db:"journal_kind"`
	Status               JournalStatus `json:"status" db:"status"`
	SourceModule         *string       `json:"source_module" db:"source_module"`
	SourceDocumentType   *string       `json:"source_document_type" db:"source_document_type"`
	SourceDocumentID     *string       `json:"source_document_id" db:"source_document_id"`
	SourceEventID        *string       `json:"source_event_id" db:"source_event_id"`
	IdempotencyKey       *string       `json:"idempotency_key" db:"idempotency_key"`
	JournalDate          time.Time     `json:"journal_date" db:"journal_date"`
	PostingDate          time.Time     `json:"posting_date" db:"posting_date"`
	CurrencyCode         string        `json:"currency_code" db:"currency_code"`
	ExchangeRate         float64       `json:"exchange_rate" db:"exchange_rate"`
	ReferenceNo          *string       `json:"reference_no" db:"reference_no"`
	ExternalReferenceNo  *string       `json:"external_reference_no" db:"external_reference_no"`
	Narration            *string       `json:"narration" db:"narration"`
	ReversalOfJournalID  *string       `json:"reversal_of_journal_id" db:"reversal_of_journal_id"`
	ReversedByJournalID  *string       `json:"reversed_by_journal_id" db:"reversed_by_journal_id"`
	Metadata             utils.JSONB   `json:"metadata" db:"metadata"`
	PostedAt             *time.Time    `json:"posted_at" db:"posted_at"`
	PostedBy             *string       `json:"posted_by" db:"posted_by"`
	ReversedAt           *time.Time    `json:"reversed_at" db:"reversed_at"`
	ReversedBy           *string       `json:"reversed_by" db:"reversed_by"`
	BaseModel
}
