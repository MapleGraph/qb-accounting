package dto

type CreateFiscalYearRequest struct {
	BookID    string `json:"book_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID string `json:"company_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440001"`
	Code      string `json:"code" validate:"required,max=30" example:"FY2026"`
	Name      string `json:"name" validate:"required,max=200" example:"FY 2026-27"`
	StartDate string `json:"start_date" validate:"required" example:"2026-04-01T00:00:00Z"`
	EndDate   string `json:"end_date" validate:"required" example:"2027-03-31T23:59:59Z"`
	Status    string `json:"status" validate:"omitempty,oneof=DRAFT OPEN CLOSED ARCHIVED" example:"DRAFT"`
}

type UpdateFiscalYearRequest struct {
	Name    *string `json:"name" validate:"omitempty,max=200" example:"FY 2026-27 (revised)"`
	EndDate *string `json:"end_date" validate:"omitempty" example:"2027-03-31T23:59:59Z"`
	Status  *string `json:"status" validate:"omitempty,oneof=DRAFT OPEN CLOSED ARCHIVED" example:"OPEN"`
}

type FiscalYearResponse struct {
	ID            string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440020"`
	BookID        string  `json:"book_id" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID     string  `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	Code          string  `json:"code" example:"FY2026"`
	Name          string  `json:"name" example:"FY 2026-27"`
	StartDate     string  `json:"start_date" example:"2026-04-01T00:00:00Z"`
	EndDate       string  `json:"end_date" example:"2027-03-31T23:59:59Z"`
	Status        string  `json:"status" example:"OPEN"`
	CloseSequence int     `json:"close_sequence" example:"0"`
	ClosedAt      *string `json:"closed_at" example:"2027-04-01T10:00:00Z"`
	ClosedBy      *string `json:"closed_by" example:"550e8400-e29b-41d4-a716-446655440099"`
}
