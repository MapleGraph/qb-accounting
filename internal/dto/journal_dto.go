package dto

// JournalLineRequest is a single journal line on create.
type JournalLineRequest struct {
	AccountID        string  `json:"account_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440070"`
	BranchID         *string `json:"branch_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440080"`
	PartyID          *string `json:"party_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440081"`
	EmployeeID       *string `json:"employee_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440082"`
	CostCenterID     *string `json:"cost_center_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440083"`
	IncomeHeadID     *string `json:"income_head_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440084"`
	TaxCodeID        *string `json:"tax_code_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440085"`
	Description      *string `json:"description" validate:"omitempty,max=500" example:"Rent expense"`
	DebitAmountTxn   float64 `json:"debit_amount_txn" validate:"gte=0" example:"10000.50"`
	CreditAmountTxn  float64 `json:"credit_amount_txn" validate:"gte=0" example:"0"`
	DebitAmountBase  float64 `json:"debit_amount_base" validate:"gte=0" example:"10000.50"`
	CreditAmountBase float64 `json:"credit_amount_base" validate:"gte=0" example:"0"`
	ExchangeRate     float64 `json:"exchange_rate" validate:"gte=0" example:"1"`
}

// CreateJournalRequest is the request body for creating a journal (draft).
type CreateJournalRequest struct {
	BookID             string               `json:"book_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440010"`
	FiscalYearID       string               `json:"fiscal_year_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440020"`
	PeriodID           string               `json:"period_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440030"`
	CompanyID          string               `json:"company_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID           *string              `json:"branch_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440080"`
	JournalKind        string               `json:"journal_kind" validate:"required,oneof=MANUAL OPENING ADJUSTMENT REVERSAL SYSTEM CLOSING" example:"MANUAL"`
	SourceModule       *string              `json:"source_module" validate:"omitempty,max=50" example:"BILLING"`
	SourceDocumentType *string              `json:"source_document_type" validate:"omitempty,max=80" example:"INVOICE"`
	SourceDocumentID   *string              `json:"source_document_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440100"`
	SourceEventID      *string              `json:"source_event_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440101"`
	IdempotencyKey     *string              `json:"idempotency_key" validate:"omitempty,max=200" example:"inv-2026-001-post"`
	JournalDate        string               `json:"journal_date" validate:"required" example:"2026-04-15T00:00:00Z"`
	PostingDate        string               `json:"posting_date" validate:"required" example:"2026-04-15T00:00:00Z"`
	CurrencyCode       string               `json:"currency_code" validate:"required,len=3" example:"INR"`
	ExchangeRate       float64              `json:"exchange_rate" validate:"required,gte=0" example:"1"`
	ReferenceNo        *string              `json:"reference_no" validate:"omitempty,max=80" example:"REF-7788"`
	Narration          *string              `json:"narration" validate:"omitempty,max=2000" example:"Monthly accruals"`
	Lines              []JournalLineRequest `json:"lines" validate:"required,min=1,dive"`
}

// PostJournalRequest posts a draft journal.
type PostJournalRequest struct {
	JournalID   string  `json:"journal_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440110"`
	PostingDate *string `json:"posting_date" validate:"omitempty" example:"2026-04-16T00:00:00Z"`
}

// ReverseJournalRequest reverses a posted journal.
type ReverseJournalRequest struct {
	JournalID    string  `json:"journal_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440110"`
	ReversalDate *string `json:"reversal_date" validate:"omitempty" example:"2026-04-20T00:00:00Z"`
	Narration    *string `json:"narration" validate:"omitempty,max=2000" example:"Reversal of incorrect accrual"`
}

// JournalLineResponse is a journal line in API responses.
type JournalLineResponse struct {
	ID               string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440120"`
	JournalID        string  `json:"journal_id" example:"550e8400-e29b-41d4-a716-446655440110"`
	LineNo           int     `json:"line_no" example:"1"`
	AccountID        string  `json:"account_id" example:"550e8400-e29b-41d4-a716-446655440070"`
	CompanyID        string  `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID         *string `json:"branch_id" example:"550e8400-e29b-41d4-a716-446655440080"`
	PartyID          *string `json:"party_id" example:"550e8400-e29b-41d4-a716-446655440081"`
	EmployeeID       *string `json:"employee_id" example:"550e8400-e29b-41d4-a716-446655440082"`
	CostCenterID     *string `json:"cost_center_id" example:"550e8400-e29b-41d4-a716-446655440083"`
	IncomeHeadID     *string `json:"income_head_id" example:"550e8400-e29b-41d4-a716-446655440084"`
	TaxCodeID        *string `json:"tax_code_id" example:"550e8400-e29b-41d4-a716-446655440085"`
	Description      *string `json:"description" example:"Rent expense"`
	DebitAmountTxn   float64 `json:"debit_amount_txn" example:"10000.50"`
	CreditAmountTxn  float64 `json:"credit_amount_txn" example:"0"`
	DebitAmountBase  float64 `json:"debit_amount_base" example:"10000.50"`
	CreditAmountBase float64 `json:"credit_amount_base" example:"0"`
	ExchangeRate     float64 `json:"exchange_rate" example:"1"`
}

// JournalResponse is the API response for a journal including lines.
type JournalResponse struct {
	ID                  string                `json:"id" example:"550e8400-e29b-41d4-a716-446655440110"`
	BookID              string                `json:"book_id" example:"550e8400-e29b-41d4-a716-446655440010"`
	BatchID             *string               `json:"batch_id" example:"550e8400-e29b-41d4-a716-446655440115"`
	FiscalYearID        string                `json:"fiscal_year_id" example:"550e8400-e29b-41d4-a716-446655440020"`
	PeriodID            string                `json:"period_id" example:"550e8400-e29b-41d4-a716-446655440030"`
	CompanyID           string                `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID            *string               `json:"branch_id" example:"550e8400-e29b-41d4-a716-446655440080"`
	JournalNo           string                `json:"journal_no" example:"JV-000042"`
	JournalKind         string                `json:"journal_kind" example:"MANUAL"`
	Status              string                `json:"status" example:"POSTED"`
	SourceModule        *string               `json:"source_module" example:"BILLING"`
	SourceDocumentType  *string               `json:"source_document_type" example:"INVOICE"`
	SourceDocumentID    *string               `json:"source_document_id" example:"550e8400-e29b-41d4-a716-446655440100"`
	SourceEventID       *string               `json:"source_event_id" example:"550e8400-e29b-41d4-a716-446655440101"`
	IdempotencyKey      *string               `json:"idempotency_key" example:"inv-2026-001-post"`
	JournalDate         string                `json:"journal_date" example:"2026-04-15T00:00:00Z"`
	PostingDate         string                `json:"posting_date" example:"2026-04-15T00:00:00Z"`
	CurrencyCode        string                `json:"currency_code" example:"INR"`
	ExchangeRate        float64               `json:"exchange_rate" example:"1"`
	ReferenceNo         *string               `json:"reference_no" example:"REF-7788"`
	ExternalReferenceNo *string               `json:"external_reference_no" example:"EXT-999"`
	Narration           *string               `json:"narration" example:"Monthly accruals"`
	ReversalOfJournalID *string               `json:"reversal_of_journal_id" example:"550e8400-e29b-41d4-a716-446655440111"`
	ReversedByJournalID *string               `json:"reversed_by_journal_id" example:"550e8400-e29b-41d4-a716-446655440112"`
	PostedAt            *string               `json:"posted_at" example:"2026-04-15T10:30:00Z"`
	PostedBy            *string               `json:"posted_by" example:"550e8400-e29b-41d4-a716-446655440099"`
	ReversedAt          *string               `json:"reversed_at" example:"2026-04-20T11:00:00Z"`
	ReversedBy          *string               `json:"reversed_by" example:"550e8400-e29b-41d4-a716-446655440099"`
	Lines               []JournalLineResponse `json:"lines"`
}
