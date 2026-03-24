package dto

import "encoding/json"

// CreatePostingRequest is the request body for enqueueing a posting run.
type CreatePostingRequest struct {
	BookID               string          `json:"book_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID            string          `json:"company_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID             *string         `json:"branch_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440080"`
	SourceModule         string          `json:"source_module" validate:"required,max=50" example:"BILLING"`
	SourceDocumentType   string          `json:"source_document_type" validate:"required,max=80" example:"INVOICE"`
	SourceDocumentID     string          `json:"source_document_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440100"`
	SourceEventID        string          `json:"source_event_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440101"`
	IdempotencyKey       string          `json:"idempotency_key" validate:"required,max=200" example:"billing:inv:550e8400-e29b-41d4-a716-446655440100:v1"`
	RequestedPostingDate string          `json:"requested_posting_date" validate:"required" example:"2026-04-15T00:00:00Z"`
	RequestPayload       json.RawMessage `json:"request_payload" validate:"required" swaggertype:"object"`
}

// UpdatePostingRequest updates lifecycle fields on a posting request.
type UpdatePostingRequest struct {
	RequestStatus        *string          `json:"request_status" validate:"omitempty,oneof=RECEIVED VALIDATED POSTED FAILED REVERSED IGNORED"`
	RuleVersionID        *string          `json:"rule_version_id" validate:"omitempty,uuid"`
	CurrentJournalID     *string          `json:"current_journal_id" validate:"omitempty,uuid"`
	ErrorCode            *string          `json:"error_code" validate:"omitempty,max=80"`
	ErrorMessage         *string          `json:"error_message" validate:"omitempty,max=2000"`
	RetryCount           *int             `json:"retry_count" validate:"omitempty,min=0"`
	RequestedPostingDate *string          `json:"requested_posting_date" validate:"omitempty"`
	RequestPayload       *json.RawMessage `json:"request_payload" swaggertype:"object"`
}


// PostingRequestResponse is the API response for a posting request.
type PostingRequestResponse struct {
	ID                   string          `json:"id" example:"550e8400-e29b-41d4-a716-446655440140"`
	BookID               string          `json:"book_id" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID            string          `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID             *string         `json:"branch_id" example:"550e8400-e29b-41d4-a716-446655440080"`
	SourceModule         string          `json:"source_module" example:"BILLING"`
	SourceDocumentType   string          `json:"source_document_type" example:"INVOICE"`
	SourceDocumentID     string          `json:"source_document_id" example:"550e8400-e29b-41d4-a716-446655440100"`
	SourceEventID        string          `json:"source_event_id" example:"550e8400-e29b-41d4-a716-446655440101"`
	IdempotencyKey       string          `json:"idempotency_key" example:"billing:inv:550e8400-e29b-41d4-a716-446655440100:v1"`
	Status               string          `json:"status" example:"POSTED"`
	RequestedPostingDate string          `json:"requested_posting_date" example:"2026-04-15T00:00:00Z"`
	RuleVersionID        *string         `json:"rule_version_id" example:"550e8400-e29b-41d4-a716-446655440130"`
	JournalID            *string         `json:"journal_id" example:"550e8400-e29b-41d4-a716-446655440110"`
	ErrorCode            *string         `json:"error_code" example:"RULE_NOT_FOUND"`
	ErrorMessage         *string         `json:"error_message" example:"No active rule for document type"`
	RetryCount           int             `json:"retry_count" example:"0"`
	FirstReceivedAt      string          `json:"first_received_at" example:"2026-04-15T09:00:00Z"`
	LastProcessedAt      *string         `json:"last_processed_at" example:"2026-04-15T09:00:05Z"`
	LastFailedAt         *string         `json:"last_failed_at" example:"2026-04-15T09:00:05Z"`
	RequestPayload       json.RawMessage `json:"request_payload" swaggertype:"object"`
}

// PostingRequestSnapshotResponse is a denormalized snapshot tied to a posting request.
type PostingRequestSnapshotResponse struct {
	ID               string          `json:"id" example:"550e8400-e29b-41d4-a716-446655440150"`
	PostingRequestID string          `json:"posting_request_id" example:"550e8400-e29b-41d4-a716-446655440140"`
	SnapshotType     string          `json:"snapshot_type" example:"SOURCE_DOCUMENT"`
	SnapshotVersion  int             `json:"snapshot_version" example:"1"`
	DocumentNumber   *string         `json:"document_number" example:"INV-1024"`
	DocumentDate     *string         `json:"document_date" example:"2026-04-14T00:00:00Z"`
	CounterpartyName *string         `json:"counterparty_name" example:"Acme Retail"`
	CurrencyCode     *string         `json:"currency_code" example:"INR"`
	GrossAmountTxn   *float64        `json:"gross_amount_txn" example:"1000"`
	NetAmountTxn     *float64        `json:"net_amount_txn" example:"820"`
	TaxAmountTxn     *float64        `json:"tax_amount_txn" example:"180"`
	SnapshotPayload  json.RawMessage `json:"snapshot_payload" swaggertype:"object"`
	CapturedAt       string          `json:"captured_at" example:"2026-04-15T09:00:01Z"`
}

// JournalSourceLinkResponse links a posting request to a journal.
type JournalSourceLinkResponse struct {
	ID               string `json:"id" example:"550e8400-e29b-41d4-a716-446655440160"`
	PostingRequestID string `json:"posting_request_id" example:"550e8400-e29b-41d4-a716-446655440140"`
	JournalID        string `json:"journal_id" example:"550e8400-e29b-41d4-a716-446655440110"`
	LinkRole         string `json:"link_role" example:"PRIMARY"`
	IsReversal       bool   `json:"is_reversal" example:"false"`
	CreatedAt        string `json:"created_at" example:"2026-04-15T09:00:05Z"`
}
