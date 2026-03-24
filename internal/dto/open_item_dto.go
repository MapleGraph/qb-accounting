package dto

// CreateOpenItemRequest is the request body for creating an open item.
type CreateOpenItemRequest struct {
	BookID             string  `json:"book_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID          string  `json:"company_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID           *string `json:"branch_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440080"`
	ItemSide           string  `json:"item_side" validate:"required,oneof=RECEIVABLE PAYABLE" example:"RECEIVABLE"`
	PartyID            string  `json:"party_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440170"`
	PartyType          string  `json:"party_type" validate:"required,oneof=CUSTOMER VENDOR PARTY EMPLOYEE OTHER" example:"CUSTOMER"`
	ControlAccountID   string  `json:"control_account_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440070"`
	SourceModule       string  `json:"source_module" validate:"required,max=50" example:"BILLING"`
	SourceDocumentType string  `json:"source_document_type" validate:"required,max=80" example:"INVOICE"`
	SourceDocumentID   string  `json:"source_document_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440100"`
	JournalID          string  `json:"journal_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440110"`
	DocumentNo         string  `json:"document_no" validate:"required,max=80" example:"INV-1024"`
	DocumentDate       string  `json:"document_date" validate:"required" example:"2026-04-10T00:00:00Z"`
	DueDate            *string `json:"due_date" validate:"omitempty" example:"2026-05-10T00:00:00Z"`
	CurrencyCode       string  `json:"currency_code" validate:"required,len=3" example:"INR"`
	ExchangeRate       float64 `json:"exchange_rate" validate:"required,gte=0" example:"1"`
	OriginalAmountTxn  float64 `json:"original_amount_txn" validate:"required" example:"50000"`
	OriginalAmountBase float64 `json:"original_amount_base" validate:"required" example:"50000"`
}

// UpdateOpenItemRequest is the request body for updating an open item.
type UpdateOpenItemRequest struct {
	DueDate    *string `json:"due_date" validate:"omitempty" example:"2026-06-10T00:00:00Z"`
	ItemStatus *string `json:"item_status" validate:"omitempty,oneof=OPEN PARTIALLY_ALLOCATED SETTLED WRITTEN_OFF CANCELLED" example:"PARTIALLY_ALLOCATED"`
	Remarks    *string `json:"remarks" validate:"omitempty,max=2000" example:"Customer promised partial payment"`
}

// OpenItemResponse is the API response for an open item.
type OpenItemResponse struct {
	ID                 string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440180"`
	BookID             string  `json:"book_id" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID          string  `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID           *string `json:"branch_id" example:"550e8400-e29b-41d4-a716-446655440080"`
	ItemSide           string  `json:"item_side" example:"RECEIVABLE"`
	ItemStatus         string  `json:"item_status" example:"OPEN"`
	PartyID            string  `json:"party_id" example:"550e8400-e29b-41d4-a716-446655440170"`
	PartyType          string  `json:"party_type" example:"CUSTOMER"`
	ControlAccountID   string  `json:"control_account_id" example:"550e8400-e29b-41d4-a716-446655440070"`
	SourceModule       string  `json:"source_module" example:"BILLING"`
	SourceDocumentType string  `json:"source_document_type" example:"INVOICE"`
	SourceDocumentID   string  `json:"source_document_id" example:"550e8400-e29b-41d4-a716-446655440100"`
	SourceLineRef      *string `json:"source_line_ref" example:"LINE-1"`
	JournalID          string  `json:"journal_id" example:"550e8400-e29b-41d4-a716-446655440110"`
	JournalLineID      *string `json:"journal_line_id" example:"550e8400-e29b-41d4-a716-446655440120"`
	DocumentNo         string  `json:"document_no" example:"INV-1024"`
	DocumentDate       string  `json:"document_date" example:"2026-04-10T00:00:00Z"`
	DueDate            *string `json:"due_date" example:"2026-05-10T00:00:00Z"`
	CurrencyCode       string  `json:"currency_code" example:"INR"`
	ExchangeRate       float64 `json:"exchange_rate" example:"1"`
	OriginalAmountTxn  float64 `json:"original_amount_txn" example:"50000"`
	OriginalAmountBase float64 `json:"original_amount_base" example:"50000"`
	OpenAmountTxn      float64 `json:"open_amount_txn" example:"30000"`
	OpenAmountBase     float64 `json:"open_amount_base" example:"30000"`
	SettledAmountTxn   float64 `json:"settled_amount_txn" example:"20000"`
	SettledAmountBase  float64 `json:"settled_amount_base" example:"20000"`
	WriteoffAmountTxn  float64 `json:"writeoff_amount_txn" example:"0"`
	WriteoffAmountBase float64 `json:"writeoff_amount_base" example:"0"`
	Remarks            *string `json:"remarks" example:"On track"`
}

// AllocationRequest allocates between two open items.
type AllocationRequest struct {
	FromOpenItemID         string  `json:"from_open_item_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440180"`
	ToOpenItemID           string  `json:"to_open_item_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440181"`
	AllocationDate         string  `json:"allocation_date" validate:"required" example:"2026-04-20T00:00:00Z"`
	AllocationCurrencyCode string  `json:"allocation_currency_code" validate:"required,len=3" example:"INR"`
	AllocationAmountTxn    float64 `json:"allocation_amount_txn" validate:"required" example:"10000"`
	AllocationAmountBase   float64 `json:"allocation_amount_base" validate:"required" example:"10000"`
}

// AllocationResponse is the API response for an allocation.
type AllocationResponse struct {
	ID                     string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440190"`
	BookID                 string  `json:"book_id" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID              string  `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	AllocationStatus       string  `json:"allocation_status" example:"APPLIED"`
	AllocationDate         string  `json:"allocation_date" example:"2026-04-20T00:00:00Z"`
	FromOpenItemID         string  `json:"from_open_item_id" example:"550e8400-e29b-41d4-a716-446655440180"`
	ToOpenItemID           string  `json:"to_open_item_id" example:"550e8400-e29b-41d4-a716-446655440181"`
	AllocationCurrencyCode string  `json:"allocation_currency_code" example:"INR"`
	AllocationAmountTxn    float64 `json:"allocation_amount_txn" example:"10000"`
	AllocationAmountBase   float64 `json:"allocation_amount_base" example:"10000"`
	AllocationJournalID    *string `json:"allocation_journal_id" example:"550e8400-e29b-41d4-a716-446655440110"`
	ReferenceNo            *string `json:"reference_no" example:"ALLOC-0099"`
	ReversalOfAllocationID *string `json:"reversal_of_allocation_id" example:"550e8400-e29b-41d4-a716-446655440195"`
	Notes                  *string `json:"notes" example:"Knock-off against advance"`
}

// AdjustmentRequest applies an adjustment to an open item.
type AdjustmentRequest struct {
	OpenItemID     string  `json:"open_item_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440180"`
	AdjustmentType string  `json:"adjustment_type" validate:"required,oneof=WRITE_OFF ROUND_OFF MANUAL_ADJUSTMENT FX_REVALUATION" example:"WRITE_OFF"`
	AdjustmentDate string  `json:"adjustment_date" validate:"required" example:"2026-04-25T00:00:00Z"`
	AmountTxn      float64 `json:"amount_txn" validate:"required" example:"500"`
	AmountBase     float64 `json:"amount_base" validate:"required" example:"500"`
	ReasonCode     *string `json:"reason_code" validate:"omitempty,max=50" example:"BAD_DEBT"`
	Notes          *string `json:"notes" validate:"omitempty,max=2000" example:"Approved by CFO"`
}

// AdjustmentResponse is the API response for an open item adjustment.
type AdjustmentResponse struct {
	ID                  string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440200"`
	OpenItemID          string  `json:"open_item_id" example:"550e8400-e29b-41d4-a716-446655440180"`
	AdjustmentType      string  `json:"adjustment_type" example:"WRITE_OFF"`
	AdjustmentDate      string  `json:"adjustment_date" example:"2026-04-25T00:00:00Z"`
	AdjustmentJournalID *string `json:"adjustment_journal_id" example:"550e8400-e29b-41d4-a716-446655440110"`
	AmountTxn           float64 `json:"amount_txn" example:"500"`
	AmountBase          float64 `json:"amount_base" example:"500"`
	ReasonCode          *string `json:"reason_code" example:"BAD_DEBT"`
	Notes               *string `json:"notes" example:"Approved by CFO"`
}
