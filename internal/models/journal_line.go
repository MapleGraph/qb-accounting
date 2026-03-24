package models

import (
	"time"

	"qb-accounting/internal/utils"
)

type JournalLine struct {
	ID              string      `json:"id" db:"id" db_pk:"true"`
	JournalID       string      `json:"journal_id" db:"journal_id"`
	LineNo          int         `json:"line_no" db:"line_no"`
	AccountID       string      `json:"account_id" db:"account_id"`
	CompanyID       string      `json:"company_id" db:"company_id"`
	BranchID        *string     `json:"branch_id" db:"branch_id"`
	PartyID         *string     `json:"party_id" db:"party_id"`
	EmployeeID      *string     `json:"employee_id" db:"employee_id"`
	CostCenterID    *string     `json:"cost_center_id" db:"cost_center_id"`
	IncomeHeadID    *string     `json:"income_head_id" db:"income_head_id"`
	TaxCodeID       *string     `json:"tax_code_id" db:"tax_code_id"`
	ProjectID       *string     `json:"project_id" db:"project_id"`
	ItemID          *string     `json:"item_id" db:"item_id"`
	Description     *string     `json:"description" db:"description"`
	DebitAmountTxn  float64     `json:"debit_amount_txn" db:"debit_amount_txn"`
	CreditAmountTxn float64     `json:"credit_amount_txn" db:"credit_amount_txn"`
	DebitAmountBase float64     `json:"debit_amount_base" db:"debit_amount_base"`
	CreditAmountBase float64    `json:"credit_amount_base" db:"credit_amount_base"`
	ExchangeRate    float64     `json:"exchange_rate" db:"exchange_rate"`
	LineMetadata    utils.JSONB `json:"line_metadata" db:"line_metadata"`
	CreatedBy       *string     `json:"created_by" db:"created_by"`
	CreatedAt       time.Time   `json:"created_at" db:"created_at" db_auto:"true"`
}
