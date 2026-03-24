package models

import "time"

type PeriodStatus string

const (
	PeriodStatusDraft      PeriodStatus = "DRAFT"
	PeriodStatusOpen       PeriodStatus = "OPEN"
	PeriodStatusSoftLocked PeriodStatus = "SOFT_LOCKED"
	PeriodStatusHardLocked PeriodStatus = "HARD_LOCKED"
	PeriodStatusClosed     PeriodStatus = "CLOSED"
)

type AccountingPeriod struct {
	ID                 string       `json:"id" db:"id" db_pk:"true"`
	BookID             string       `json:"book_id" db:"book_id"`
	FiscalYearID       string       `json:"fiscal_year_id" db:"fiscal_year_id"`
	CompanyID          string       `json:"company_id" db:"company_id"`
	PeriodNo           int16        `json:"period_no" db:"period_no"`
	PeriodName         string       `json:"period_name" db:"period_name"`
	StartDate          time.Time    `json:"start_date" db:"start_date"`
	EndDate            time.Time    `json:"end_date" db:"end_date"`
	Status             PeriodStatus `json:"status" db:"status"`
	IsAdjustmentPeriod bool         `json:"is_adjustment_period" db:"is_adjustment_period"`
	SoftLockedAt       *time.Time   `json:"soft_locked_at" db:"soft_locked_at"`
	SoftLockedBy       *string      `json:"soft_locked_by" db:"soft_locked_by"`
	HardLockedAt       *time.Time   `json:"hard_locked_at" db:"hard_locked_at"`
	HardLockedBy       *string      `json:"hard_locked_by" db:"hard_locked_by"`
	ClosedAt           *time.Time   `json:"closed_at" db:"closed_at"`
	ClosedBy           *string      `json:"closed_by" db:"closed_by"`
	LockReason         *string      `json:"lock_reason" db:"lock_reason"`
	BaseModel
}
