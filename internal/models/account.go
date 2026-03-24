package models

import "qb-accounting/internal/utils"

type AccountUsage string

const (
	AccountUsageHeader   AccountUsage = "HEADER"
	AccountUsagePostable AccountUsage = "POSTABLE"
)

type Account struct {
	ID                      string        `json:"id" db:"id" db_pk:"true"`
	BookID                  string        `json:"book_id" db:"book_id"`
	CompanyID               string        `json:"company_id" db:"company_id"`
	Code                    string        `json:"code" db:"code"`
	Name                    string        `json:"name" db:"name"`
	DisplayName             *string       `json:"display_name" db:"display_name"`
	GroupID                 string        `json:"group_id" db:"group_id"`
	ParentAccountID         *string       `json:"parent_account_id" db:"parent_account_id"`
	AccountNature           AccountNature `json:"account_nature" db:"account_nature"`
	Usage                   AccountUsage  `json:"usage" db:"usage"`
	NormalBalance           string        `json:"normal_balance" db:"normal_balance"`
	ControlType             *string       `json:"control_type" db:"control_type"`
	AllowManualPosting      bool          `json:"allow_manual_posting" db:"allow_manual_posting"`
	RequireParty            bool          `json:"require_party" db:"require_party"`
	RequireBranch           bool          `json:"require_branch" db:"require_branch"`
	RequireCostCenter       bool          `json:"require_cost_center" db:"require_cost_center"`
	RequireEmployee         bool          `json:"require_employee" db:"require_employee"`
	RequireTaxBreakup       bool          `json:"require_tax_breakup" db:"require_tax_breakup"`
	IsSystem                bool          `json:"is_system" db:"is_system"`
	IsActive                bool          `json:"is_active" db:"is_active"`
	ExternalDimensionPolicy utils.JSONB   `json:"external_dimension_policy" db:"external_dimension_policy"`
	BaseModel
}
