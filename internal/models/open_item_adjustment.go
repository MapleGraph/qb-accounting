package models

import "time"

type OpenItemAdjustmentType string

const (
	AdjustmentTypeWriteOff         OpenItemAdjustmentType = "WRITE_OFF"
	AdjustmentTypeRoundOff         OpenItemAdjustmentType = "ROUND_OFF"
	AdjustmentTypeManual           OpenItemAdjustmentType = "MANUAL_ADJUSTMENT"
	AdjustmentTypeFXRevaluation    OpenItemAdjustmentType = "FX_REVALUATION"
)

type OpenItemAdjustment struct {
	ID                  string                 `json:"id" db:"id" db_pk:"true"`
	OpenItemID          string                 `json:"open_item_id" db:"open_item_id"`
	AdjustmentType      OpenItemAdjustmentType `json:"adjustment_type" db:"adjustment_type"`
	AdjustmentDate      time.Time              `json:"adjustment_date" db:"adjustment_date"`
	AdjustmentJournalID *string                `json:"adjustment_journal_id" db:"adjustment_journal_id"`
	AmountTxn           float64                `json:"amount_txn" db:"amount_txn"`
	AmountBase          float64                `json:"amount_base" db:"amount_base"`
	ReasonCode          *string                `json:"reason_code" db:"reason_code"`
	Notes               *string                `json:"notes" db:"notes"`
	CreatedBy           *string                `json:"created_by" db:"created_by"`
	CreatedAt           time.Time              `json:"created_at" db:"created_at" db_auto:"true"`
}
