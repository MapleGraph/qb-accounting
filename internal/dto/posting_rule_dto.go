package dto

import "encoding/json"

// CreatePostingRuleVersionRequest is the request body for creating a posting rule version.
type CreatePostingRuleVersionRequest struct {
	SourceModule       string          `json:"source_module" validate:"required,max=50" example:"BILLING"`
	SourceDocumentType string          `json:"source_document_type" validate:"required,max=80" example:"INVOICE"`
	VersionNo          int             `json:"version_no" validate:"required,min=1" example:"1"`
	Name               string          `json:"name" validate:"required,max=200" example:"Default invoice posting"`
	Status             string          `json:"status" validate:"omitempty,oneof=DRAFT ACTIVE RETIRED" example:"DRAFT"`
	EffectiveFrom      string          `json:"effective_from" validate:"required" example:"2026-04-01T00:00:00Z"`
	EffectiveTo        *string         `json:"effective_to" validate:"omitempty" example:"2027-03-31T23:59:59Z"`
	RulePayload        json.RawMessage `json:"rule_payload" validate:"required" swaggertype:"object"`
	Notes              *string         `json:"notes" validate:"omitempty,max=2000" example:"Initial version"`
}

// UpdatePostingRuleVersionRequest is the request body for updating a posting rule version.
type UpdatePostingRuleVersionRequest struct {
	Name        *string          `json:"name" validate:"omitempty,max=200" example:"Default invoice posting (v2)"`
	Status      *string          `json:"status" validate:"omitempty,oneof=DRAFT ACTIVE RETIRED" example:"ACTIVE"`
	EffectiveTo *string          `json:"effective_to" validate:"omitempty" example:"2027-03-31T23:59:59Z"`
	RulePayload *json.RawMessage `json:"rule_payload" swaggertype:"object"`
	Notes       *string          `json:"notes" validate:"omitempty,max=2000" example:"Added tax split"`
}

// PostingRuleVersionResponse is the API response for a posting rule version.
type PostingRuleVersionResponse struct {
	ID                 string          `json:"id" example:"550e8400-e29b-41d4-a716-446655440130"`
	SourceModule       string          `json:"source_module" example:"BILLING"`
	SourceDocumentType string          `json:"source_document_type" example:"INVOICE"`
	VersionNo          int             `json:"version_no" example:"1"`
	Name               string          `json:"name" example:"Default invoice posting"`
	Status             string          `json:"status" example:"ACTIVE"`
	EffectiveFrom      string          `json:"effective_from" example:"2026-04-01T00:00:00Z"`
	EffectiveTo        *string         `json:"effective_to" example:"2027-03-31T23:59:59Z"`
	RulePayload        json.RawMessage `json:"rule_payload" swaggertype:"object"`
	Notes              *string         `json:"notes" example:"Initial version"`
}
