package models

import "time"

type JournalLinkRole string

const (
	JournalLinkRolePrimary    JournalLinkRole = "PRIMARY"
	JournalLinkRoleReversal   JournalLinkRole = "REVERSAL"
	JournalLinkRoleAdjustment JournalLinkRole = "ADJUSTMENT"
)

type JournalSourceLink struct {
	ID               string          `json:"id" db:"id" db_pk:"true"`
	PostingRequestID string          `json:"posting_request_id" db:"posting_request_id"`
	JournalID        string          `json:"journal_id" db:"journal_id"`
	LinkRole         JournalLinkRole `json:"link_role" db:"link_role"`
	IsReversal       bool            `json:"is_reversal" db:"is_reversal"`
	CreatedAt        time.Time       `json:"created_at" db:"created_at" db_auto:"true"`
}
