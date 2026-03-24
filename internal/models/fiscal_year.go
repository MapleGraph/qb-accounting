package models

import "time"

type FiscalYearStatus string

const (
	FiscalYearStatusDraft    FiscalYearStatus = "DRAFT"
	FiscalYearStatusOpen     FiscalYearStatus = "OPEN"
	FiscalYearStatusClosed   FiscalYearStatus = "CLOSED"
	FiscalYearStatusArchived FiscalYearStatus = "ARCHIVED"
)

type FiscalYear struct {
	ID            string           `json:"id" db:"id" db_pk:"true"`
	BookID        string           `json:"book_id" db:"book_id"`
	CompanyID     string           `json:"company_id" db:"company_id"`
	Code          string           `json:"code" db:"code"`
	Name          string           `json:"name" db:"name"`
	StartDate     time.Time        `json:"start_date" db:"start_date"`
	EndDate       time.Time        `json:"end_date" db:"end_date"`
	Status        FiscalYearStatus `json:"status" db:"status"`
	CloseSequence int              `json:"close_sequence" db:"close_sequence"`
	ClosedAt      *time.Time       `json:"closed_at" db:"closed_at"`
	ClosedBy      *string          `json:"closed_by" db:"closed_by"`
	BaseModel
}
