package dto

type CreateVoucherSequenceRequest struct {
	BookID      string  `json:"book_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID   string  `json:"company_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID    *string `json:"branch_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440080"`
	VoucherType string  `json:"voucher_type" validate:"required,oneof=JV RV PV SV" example:"JV"`
	Prefix      *string `json:"prefix" validate:"omitempty,max=20" example:"JV-"`
	Suffix      *string `json:"suffix" validate:"omitempty,max=20" example:"-FY26"`
	Padding     int16   `json:"padding" validate:"required,min=1,max=12" example:"6"`
	NextNumber  int64   `json:"next_number" validate:"required,min=1" example:"1"`
	ResetPolicy string  `json:"reset_policy" validate:"required,oneof=FY NEVER" example:"FY"`
	IsActive    bool    `json:"is_active" example:"true"`
}

type UpdateVoucherSequenceRequest struct {
	BranchID    *string `json:"branch_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440080"`
	Prefix      *string `json:"prefix" validate:"omitempty,max=20" example:"JV-"`
	Suffix      *string `json:"suffix" validate:"omitempty,max=20" example:"-FY26"`
	Padding     *int16  `json:"padding" validate:"omitempty,min=1,max=12" example:"6"`
	NextNumber  *int64  `json:"next_number" validate:"omitempty,min=1" example:"42"`
	ResetPolicy *string `json:"reset_policy" validate:"omitempty,oneof=FY NEVER" example:"NEVER"`
	IsActive    *bool   `json:"is_active" example:"true"`
}

type VoucherSequenceResponse struct {
	ID          string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440090"`
	BookID      string  `json:"book_id" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID   string  `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	BranchID    *string `json:"branch_id" example:"550e8400-e29b-41d4-a716-446655440080"`
	VoucherType string  `json:"voucher_type" example:"JV"`
	Prefix      *string `json:"prefix" example:"JV-"`
	Suffix      *string `json:"suffix" example:"-FY26"`
	Padding     int16   `json:"padding" example:"6"`
	NextNumber  int64   `json:"next_number" example:"1"`
	ResetPolicy string  `json:"reset_policy" example:"FY"`
	IsActive    bool    `json:"is_active" example:"true"`
}
