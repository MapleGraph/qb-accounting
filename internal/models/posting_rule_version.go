package models

import (
	"time"

	"qb-accounting/internal/utils"
)

type PostingRuleStatus string

const (
	PostingRuleStatusDraft   PostingRuleStatus = "DRAFT"
	PostingRuleStatusActive  PostingRuleStatus = "ACTIVE"
	PostingRuleStatusRetired PostingRuleStatus = "RETIRED"
)

type PostingRuleVersion struct {
	ID                 string            `json:"id" db:"id" db_pk:"true"`
	SourceModule       string            `json:"source_module" db:"source_module"`
	SourceDocumentType string            `json:"source_document_type" db:"source_document_type"`
	VersionNo          int               `json:"version_no" db:"version_no"`
	Name               string            `json:"name" db:"name"`
	Status             PostingRuleStatus `json:"status" db:"status"`
	EffectiveFrom      time.Time         `json:"effective_from" db:"effective_from"`
	EffectiveTo        *time.Time        `json:"effective_to" db:"effective_to"`
	RulePayload        utils.JSONB       `json:"rule_payload" db:"rule_payload"`
	Notes              *string           `json:"notes" db:"notes"`
	BaseModel
}
