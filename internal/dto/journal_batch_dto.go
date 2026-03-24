package dto

type CreateJournalBatchRequest struct {
	BookID       string  `json:"book_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID    string  `json:"company_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID     *string `json:"branch_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440080"`
	BatchNo      string  `json:"batch_no" validate:"required,max=50" example:"BATCH-20260415-001"`
	SourceModule *string `json:"source_module" validate:"omitempty,max=50" example:"PAYROLL"`
	BatchType    string  `json:"batch_type" validate:"required,max=50" example:"SALARY_POSTING"`
	PostingDate  string  `json:"posting_date" validate:"required" example:"2026-04-15T00:00:00Z"`
	Narration    *string `json:"narration" validate:"omitempty,max=2000" example:"April payroll run"`
}

type UpdateJournalBatchRequest struct {
	SourceModule *string `json:"source_module" validate:"omitempty,max=50" example:"PAYROLL"`
	BatchType    *string `json:"batch_type" validate:"omitempty,max=50" example:"SALARY_POSTING"`
	PostingDate  *string `json:"posting_date" validate:"omitempty" example:"2026-04-16T00:00:00Z"`
	Status       *string `json:"status" validate:"omitempty,oneof=OPEN POSTED FAILED CANCELLED" example:"POSTED"`
	Narration    *string `json:"narration" validate:"omitempty,max=2000" example:"April payroll (corrected)"`
}

type JournalBatchResponse struct {
	ID           string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440115"`
	BookID       string  `json:"book_id" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID    string  `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID     *string `json:"branch_id" example:"550e8400-e29b-41d4-a716-446655440080"`
	BatchNo      string  `json:"batch_no" example:"BATCH-20260415-001"`
	SourceModule *string `json:"source_module" example:"PAYROLL"`
	BatchType    string  `json:"batch_type" example:"SALARY_POSTING"`
	PostingDate  string  `json:"posting_date" example:"2026-04-15T00:00:00Z"`
	Status       string  `json:"status" example:"POSTED"`
	Narration    *string `json:"narration" example:"April payroll run"`
}
