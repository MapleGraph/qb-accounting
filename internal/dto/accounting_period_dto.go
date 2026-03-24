package dto

type CreateAccountingPeriodRequest struct {
	BookID             string `json:"book_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440010"`
	FiscalYearID       string `json:"fiscal_year_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440020"`
	CompanyID          string `json:"company_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440001"`
	PeriodNo           int16  `json:"period_no" validate:"required,min=1" example:"1"`
	PeriodName         string `json:"period_name" validate:"required,max=120" example:"Apr 2026"`
	StartDate          string `json:"start_date" validate:"required" example:"2026-04-01T00:00:00Z"`
	EndDate            string `json:"end_date" validate:"required" example:"2026-04-30T23:59:59Z"`
	IsAdjustmentPeriod bool   `json:"is_adjustment_period" example:"false"`
	Status             string `json:"status" validate:"omitempty,oneof=DRAFT OPEN SOFT_LOCKED HARD_LOCKED CLOSED" example:"DRAFT"`
}

type UpdateAccountingPeriodRequest struct {
	PeriodName      *string `json:"period_name" validate:"omitempty,max=120" example:"April 2026"`
	Status          *string `json:"status" validate:"omitempty,oneof=DRAFT OPEN SOFT_LOCKED HARD_LOCKED CLOSED" example:"OPEN"`
	LockReason      *string `json:"lock_reason" validate:"omitempty,max=500" example:"Month-end close"`
	ApplySoftLock   *bool   `json:"apply_soft_lock" example:"true"`
	ApplyHardLock   *bool   `json:"apply_hard_lock" example:"false"`
	ReleaseSoftLock *bool   `json:"release_soft_lock" example:"false"`
	ReleaseHardLock *bool   `json:"release_hard_lock" example:"false"`
	SoftLockedBy    *string `json:"soft_locked_by" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440099"`
	HardLockedBy    *string `json:"hard_locked_by" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440099"`
	ClosedBy        *string `json:"closed_by" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440099"`
}

type AccountingPeriodResponse struct {
	ID                 string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440030"`
	BookID             string  `json:"book_id" example:"550e8400-e29b-41d4-a716-446655440010"`
	FiscalYearID       string  `json:"fiscal_year_id" example:"550e8400-e29b-41d4-a716-446655440020"`
	CompanyID          string  `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	PeriodNo           int16   `json:"period_no" example:"1"`
	PeriodName         string  `json:"period_name" example:"Apr 2026"`
	StartDate          string  `json:"start_date" example:"2026-04-01T00:00:00Z"`
	EndDate            string  `json:"end_date" example:"2026-04-30T23:59:59Z"`
	Status             string  `json:"status" example:"OPEN"`
	IsAdjustmentPeriod bool    `json:"is_adjustment_period" example:"false"`
	SoftLockedAt       *string `json:"soft_locked_at" example:"2026-05-01T09:00:00Z"`
	SoftLockedBy       *string `json:"soft_locked_by" example:"550e8400-e29b-41d4-a716-446655440099"`
	HardLockedAt       *string `json:"hard_locked_at" example:"2026-05-02T09:00:00Z"`
	HardLockedBy       *string `json:"hard_locked_by" example:"550e8400-e29b-41d4-a716-446655440099"`
	ClosedAt           *string `json:"closed_at" example:"2026-05-03T18:00:00Z"`
	ClosedBy           *string `json:"closed_by" example:"550e8400-e29b-41d4-a716-446655440099"`
	LockReason         *string `json:"lock_reason" example:"Approved close"`
}
