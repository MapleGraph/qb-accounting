package models

import (
	"time"

	"qb-accounting/internal/utils"
)

type PostingRequestSnapshot struct {
	ID               string      `json:"id" db:"id" db_pk:"true"`
	PostingRequestID string      `json:"posting_request_id" db:"posting_request_id"`
	SnapshotType     string      `json:"snapshot_type" db:"snapshot_type"`
	SnapshotVersion  int         `json:"snapshot_version" db:"snapshot_version"`
	DocumentNumber   *string     `json:"document_number" db:"document_number"`
	DocumentDate     *time.Time  `json:"document_date" db:"document_date"`
	CounterpartyName *string     `json:"counterparty_name" db:"counterparty_name"`
	CurrencyCode     *string     `json:"currency_code" db:"currency_code"`
	GrossAmountTxn   *float64    `json:"gross_amount_txn" db:"gross_amount_txn"`
	NetAmountTxn     *float64    `json:"net_amount_txn" db:"net_amount_txn"`
	TaxAmountTxn     *float64    `json:"tax_amount_txn" db:"tax_amount_txn"`
	SnapshotPayload  utils.JSONB `json:"snapshot_payload" db:"snapshot_payload"`
	CapturedAt       time.Time   `json:"captured_at" db:"captured_at" db_auto:"true"`
}
