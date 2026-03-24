package models

import "time"

type AllocationStatus string

const (
	AllocationStatusApplied  AllocationStatus = "APPLIED"
	AllocationStatusReversed AllocationStatus = "REVERSED"
)

type OpenItemAllocation struct {
	ID                     string           `json:"id" db:"id" db_pk:"true"`
	BookID                 string           `json:"book_id" db:"book_id"`
	CompanyID              string           `json:"company_id" db:"company_id"`
	AllocationStatus       AllocationStatus `json:"allocation_status" db:"allocation_status"`
	AllocationDate         time.Time        `json:"allocation_date" db:"allocation_date"`
	FromOpenItemID         string           `json:"from_open_item_id" db:"from_open_item_id"`
	ToOpenItemID           string           `json:"to_open_item_id" db:"to_open_item_id"`
	AllocationCurrencyCode string           `json:"allocation_currency_code" db:"allocation_currency_code"`
	AllocationAmountTxn    float64          `json:"allocation_amount_txn" db:"allocation_amount_txn"`
	AllocationAmountBase   float64          `json:"allocation_amount_base" db:"allocation_amount_base"`
	AllocationJournalID    *string          `json:"allocation_journal_id" db:"allocation_journal_id"`
	ReferenceNo            *string          `json:"reference_no" db:"reference_no"`
	ReversalOfAllocationID *string          `json:"reversal_of_allocation_id" db:"reversal_of_allocation_id"`
	Notes                  *string          `json:"notes" db:"notes"`
	CreatedBy              *string          `json:"created_by" db:"created_by"`
	CreatedAt              time.Time        `json:"created_at" db:"created_at" db_auto:"true"`
}
