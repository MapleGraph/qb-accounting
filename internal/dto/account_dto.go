package dto

type CreateAccountRequest struct {
	BookID             string  `json:"book_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID          string  `json:"company_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440001"`
	Code               string  `json:"code" validate:"required,max=50" example:"110010"`
	Name               string  `json:"name" validate:"required,max=200" example:"Cash on Hand"`
	DisplayName        *string `json:"display_name" validate:"omitempty,max=200" example:"Cash"`
	GroupID            string  `json:"group_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440050"`
	ParentAccountID    *string `json:"parent_account_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440060"`
	AccountNature      string  `json:"account_nature" validate:"required,oneof=ASSET LIABILITY EQUITY INCOME EXPENSE" example:"ASSET"`
	Usage              string  `json:"usage" validate:"required,oneof=HEADER POSTABLE" example:"POSTABLE"`
	NormalBalance      string  `json:"normal_balance" validate:"required,oneof=D C" example:"D"`
	ControlType        *string `json:"control_type" validate:"omitempty,max=50" example:"CASH"`
	AllowManualPosting bool    `json:"allow_manual_posting" example:"true"`
	RequireParty       bool    `json:"require_party" example:"false"`
	RequireBranch      bool    `json:"require_branch" example:"false"`
	RequireCostCenter  bool    `json:"require_cost_center" example:"false"`
	RequireEmployee    bool    `json:"require_employee" example:"false"`
	RequireTaxBreakup  bool    `json:"require_tax_breakup" example:"false"`
	IsSystem           bool    `json:"is_system" example:"false"`
	IsActive           bool    `json:"is_active" example:"true"`
}

type UpdateAccountRequest struct {
	Name               *string `json:"name" validate:"omitempty,max=200" example:"Cash on Hand — Main"`
	DisplayName        *string `json:"display_name" validate:"omitempty,max=200" example:"Petty Cash"`
	ParentAccountID    *string `json:"parent_account_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440060"`
	ControlType        *string `json:"control_type" validate:"omitempty,max=50" example:"CASH"`
	AllowManualPosting *bool   `json:"allow_manual_posting" example:"true"`
	RequireParty       *bool   `json:"require_party" example:"false"`
	RequireBranch      *bool   `json:"require_branch" example:"false"`
	RequireCostCenter  *bool   `json:"require_cost_center" example:"false"`
	RequireEmployee    *bool   `json:"require_employee" example:"false"`
	RequireTaxBreakup  *bool   `json:"require_tax_breakup" example:"false"`
	IsActive           *bool   `json:"is_active" example:"true"`
}

type AccountResponse struct {
	ID                 string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440070"`
	BookID             string  `json:"book_id" example:"550e8400-e29b-41d4-a716-446655440010"`
	CompanyID          string  `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	Code               string  `json:"code" example:"110010"`
	Name               string  `json:"name" example:"Cash on Hand"`
	DisplayName        *string `json:"display_name" example:"Cash"`
	GroupID            string  `json:"group_id" example:"550e8400-e29b-41d4-a716-446655440050"`
	ParentAccountID    *string `json:"parent_account_id" example:"550e8400-e29b-41d4-a716-446655440060"`
	AccountNature      string  `json:"account_nature" example:"ASSET"`
	Usage              string  `json:"usage" example:"POSTABLE"`
	NormalBalance      string  `json:"normal_balance" example:"D"`
	ControlType        *string `json:"control_type" example:"CASH"`
	AllowManualPosting bool    `json:"allow_manual_posting" example:"true"`
	RequireParty       bool    `json:"require_party" example:"false"`
	RequireBranch      bool    `json:"require_branch" example:"false"`
	RequireCostCenter  bool    `json:"require_cost_center" example:"false"`
	RequireEmployee    bool    `json:"require_employee" example:"false"`
	RequireTaxBreakup  bool    `json:"require_tax_breakup" example:"false"`
	IsSystem           bool    `json:"is_system" example:"false"`
	IsActive           bool    `json:"is_active" example:"true"`
}
