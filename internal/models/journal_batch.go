package models

import "time"

type JournalBatchStatus string

const (
	JournalBatchStatusOpen      JournalBatchStatus = "OPEN"
	JournalBatchStatusPosted    JournalBatchStatus = "POSTED"
	JournalBatchStatusFailed    JournalBatchStatus = "FAILED"
	JournalBatchStatusCancelled JournalBatchStatus = "CANCELLED"
)

type JournalBatch struct {
	ID           string             `json:"id" db:"id" db_pk:"true"`
	BookID       string             `json:"book_id" db:"book_id"`
	CompanyID    string             `json:"company_id" db:"company_id"`
	BranchID     *string            `json:"branch_id" db:"branch_id"`
	BatchNo      string             `json:"batch_no" db:"batch_no"`
	SourceModule *string            `json:"source_module" db:"source_module"`
	BatchType    string             `json:"batch_type" db:"batch_type"`
	PostingDate  time.Time          `json:"posting_date" db:"posting_date"`
	Status       JournalBatchStatus `json:"status" db:"status"`
	Narration    *string            `json:"narration" db:"narration"`
	BaseModel
}
