package models

import "time"

type PartyType string

const (
	PartyTypeCustomer PartyType = "CUSTOMER"
	PartyTypeVendor   PartyType = "VENDOR"
	PartyTypeParty    PartyType = "PARTY"
	PartyTypeEmployee PartyType = "EMPLOYEE"
	PartyTypeOther    PartyType = "OTHER"
)

type OpenItemSide string

const (
	OpenItemSideReceivable OpenItemSide = "RECEIVABLE"
	OpenItemSidePayable    OpenItemSide = "PAYABLE"
)

type OpenItemStatus string

const (
	OpenItemStatusOpen               OpenItemStatus = "OPEN"
	OpenItemStatusPartiallyAllocated OpenItemStatus = "PARTIALLY_ALLOCATED"
	OpenItemStatusSettled            OpenItemStatus = "SETTLED"
	OpenItemStatusWrittenOff         OpenItemStatus = "WRITTEN_OFF"
	OpenItemStatusCancelled          OpenItemStatus = "CANCELLED"
)

type OpenItem struct {
	ID                 string         `json:"id" db:"id" db_pk:"true"`
	BookID             string         `json:"book_id" db:"book_id"`
	CompanyID          string         `json:"company_id" db:"company_id"`
	BranchID           *string        `json:"branch_id" db:"branch_id"`
	ItemSide           OpenItemSide   `json:"item_side" db:"item_side"`
	ItemStatus         OpenItemStatus `json:"item_status" db:"item_status"`
	PartyID            string         `json:"party_id" db:"party_id"`
	PartyType          PartyType      `json:"party_type" db:"party_type"`
	ControlAccountID   string         `json:"control_account_id" db:"control_account_id"`
	SourceModule       string         `json:"source_module" db:"source_module"`
	SourceDocumentType string         `json:"source_document_type" db:"source_document_type"`
	SourceDocumentID   string         `json:"source_document_id" db:"source_document_id"`
	SourceLineRef      *string        `json:"source_line_ref" db:"source_line_ref"`
	JournalID          string         `json:"journal_id" db:"journal_id"`
	JournalLineID      *string        `json:"journal_line_id" db:"journal_line_id"`
	DocumentNo         string         `json:"document_no" db:"document_no"`
	DocumentDate       time.Time      `json:"document_date" db:"document_date"`
	DueDate            *time.Time     `json:"due_date" db:"due_date"`
	CurrencyCode       string         `json:"currency_code" db:"currency_code"`
	ExchangeRate       float64        `json:"exchange_rate" db:"exchange_rate"`
	OriginalAmountTxn  float64        `json:"original_amount_txn" db:"original_amount_txn"`
	OriginalAmountBase float64        `json:"original_amount_base" db:"original_amount_base"`
	OpenAmountTxn      float64        `json:"open_amount_txn" db:"open_amount_txn"`
	OpenAmountBase     float64        `json:"open_amount_base" db:"open_amount_base"`
	SettledAmountTxn   float64        `json:"settled_amount_txn" db:"settled_amount_txn"`
	SettledAmountBase  float64        `json:"settled_amount_base" db:"settled_amount_base"`
	WriteoffAmountTxn  float64        `json:"writeoff_amount_txn" db:"writeoff_amount_txn"`
	WriteoffAmountBase float64        `json:"writeoff_amount_base" db:"writeoff_amount_base"`
	Remarks            *string        `json:"remarks" db:"remarks"`
	CreatedAt          time.Time      `json:"created_at" db:"created_at" db_auto:"true"`
	UpdatedAt          *time.Time     `json:"updated_at" db:"updated_at" db_auto:"true"`
}
