package models

import (
	"time"

	"qb-accounting/internal/utils"
)

type PostingRequestStatus string

const (
	PostingRequestStatusReceived  PostingRequestStatus = "RECEIVED"
	PostingRequestStatusValidated PostingRequestStatus = "VALIDATED"
	PostingRequestStatusPosted    PostingRequestStatus = "POSTED"
	PostingRequestStatusFailed    PostingRequestStatus = "FAILED"
	PostingRequestStatusReversed  PostingRequestStatus = "REVERSED"
	PostingRequestStatusIgnored   PostingRequestStatus = "IGNORED"
)

type PostingRequest struct {
	ID                   string               `json:"id" db:"id" db_pk:"true"`
	BookID               string               `json:"book_id" db:"book_id"`
	CompanyID            string               `json:"company_id" db:"company_id"`
	BranchID             *string              `json:"branch_id" db:"branch_id"`
	SourceModule         string               `json:"source_module" db:"source_module"`
	SourceDocumentType   string               `json:"source_document_type" db:"source_document_type"`
	SourceDocumentID     string               `json:"source_document_id" db:"source_document_id"`
	SourceEventID        string               `json:"source_event_id" db:"source_event_id"`
	IdempotencyKey       string               `json:"idempotency_key" db:"idempotency_key"`
	RequestHash          string               `json:"request_hash" db:"request_hash"`
	RequestStatus        PostingRequestStatus `json:"request_status" db:"request_status"`
	RequestedPostingDate time.Time            `json:"requested_posting_date" db:"requested_posting_date"`
	RequestedBy          *string              `json:"requested_by" db:"requested_by"`
	RuleVersionID        *string              `json:"rule_version_id" db:"rule_version_id"`
	CurrentJournalID     *string              `json:"current_journal_id" db:"current_journal_id"`
	ErrorCode            *string              `json:"error_code" db:"error_code"`
	ErrorMessage         *string              `json:"error_message" db:"error_message"`
	RetryCount           int                  `json:"retry_count" db:"retry_count"`
	FirstReceivedAt      time.Time            `json:"first_received_at" db:"first_received_at" db_auto:"true"`
	LastProcessedAt      *time.Time           `json:"last_processed_at" db:"last_processed_at"`
	LastFailedAt         *time.Time           `json:"last_failed_at" db:"last_failed_at"`
	RequestPayload       utils.JSONB          `json:"request_payload" db:"request_payload"`
	CreatedAt            time.Time            `json:"created_at" db:"created_at" db_auto:"true"`
	UpdatedAt            *time.Time           `json:"updated_at" db:"updated_at" db_auto:"true"`
}
