package remote

import (
	"context"
	"encoding/json"
	"fmt"

	pbSetup "qb-accounting/internal/proto/setup"

	qbgrpc "github.com/MapleGraph/qb-core/v2/pkg/grpc"
)

const ServiceNameSetup = "setup"

// SetupService defines the interface for setup service operations
type SetupService interface {
	GetOrganizationConfig(ctx context.Context, organizationID string) (*OrganizationConfig, error)
}

// OrganizationConfig represents the organization configuration from setup service
type OrganizationConfig struct {
	Organization *Organization
	Companies    []*Company
	Locations    []*Location
	Features     map[string]*FeatureAccess
}

// Organization represents organization details from setup service
type Organization struct {
	OrganizationID   string
	OrganizationName string
	Timezone         *string
}

// Company represents company details from setup service
type Company struct {
	CompanyID      string
	OrganizationID string
	CompanyName    string
	GSTNumber      *string
}

// Location represents location details from setup service
type Location struct {
	LocationID     string
	OrganizationID string
	CompanyID      string
	LocationName   string
	Timezone       *string
}

// FeatureAccess represents a feature access configuration
type FeatureAccess struct {
	AccessID       string
	OrganizationID string
	FeatureID      string
	SubFeatureID   string
	Settings       json.RawMessage
	IsActive       int32
}

type setupRepository struct {
	handler qbgrpc.ClientHandler
}

// NewSetupRepository creates a new setup repository using a qb-core gRPC handler.
func NewSetupRepository(handler qbgrpc.ClientHandler) SetupService {
	return &setupRepository{handler: handler}
}

func (r *setupRepository) client(ctx context.Context) (pbSetup.SetupServiceClient, error) {
	if r.handler == nil {
		return nil, fmt.Errorf("setup gRPC handler is not available")
	}
	conn, err := r.handler.GetConnectionWithRetry(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get setup gRPC connection: %w", err)
	}
	return pbSetup.NewSetupServiceClient(conn), nil
}

// GetOrganizationConfig fetches organization configuration from the setup service
func (r *setupRepository) GetOrganizationConfig(ctx context.Context, organizationID string) (*OrganizationConfig, error) {
	client, err := r.client(ctx)
	if err != nil {
		return nil, err
	}

	req := &pbSetup.OrganizationConfigRequest{
		OrganizationId: organizationID,
	}

	resp, err := client.OrganizationConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call OrganizationConfig: %w", err)
	}

	// Convert response to our domain model
	config := &OrganizationConfig{
		Companies: make([]*Company, 0, len(resp.Company)),
		Locations: make([]*Location, 0, len(resp.Locations)),
		Features:  make(map[string]*FeatureAccess),
	}

	if resp.Organization != nil {
		org := &Organization{
			OrganizationID:   resp.Organization.OrganizationId,
			OrganizationName: resp.Organization.OrganizationName,
		}
		if resp.Organization.Timezone != nil {
			tz := resp.Organization.Timezone.Value
			org.Timezone = &tz
		}
		config.Organization = org
	}

	for _, c := range resp.Company {
		company := &Company{
			CompanyID:      c.CompanyId,
			OrganizationID: c.OrganizationId,
			CompanyName:    c.CompanyName,
		}
		if c.ComplianceInfo != nil {
			if gstRaw, ok := c.ComplianceInfo.Fields["gst_number"]; ok {
				if gst := gstRaw.GetStringValue(); gst != "" {
					company.GSTNumber = &gst
				}
			}
		}
		config.Companies = append(config.Companies, company)
	}

	for _, l := range resp.Locations {
		location := &Location{
			LocationID:     l.LocationId,
			OrganizationID: l.OrganizationId,
			CompanyID:      l.CompanyId,
			LocationName:   l.LocationName,
		}
		if l.Timezone != nil {
			tz := l.Timezone.Value
			location.Timezone = &tz
		}
		config.Locations = append(config.Locations, location)
	}

	for _, fa := range resp.FeatureAccess {
		var settings json.RawMessage
		if fa.Scopes != nil {
			settingsBytes, err := fa.Scopes.MarshalJSON()
			if err != nil {
				return nil, fmt.Errorf("failed to marshal settings for feature %s: %w", fa.FeatureId, err)
			}
			settings = json.RawMessage(settingsBytes)
		} else {
			settings = json.RawMessage("{}")
		}

		config.Features[fa.FeatureId] = &FeatureAccess{
			AccessID:       fa.AccessId,
			OrganizationID: fa.OrganizationId,
			FeatureID:      fa.FeatureId,
			SubFeatureID:   fa.SubFeatureId,
			Settings:       settings,
			IsActive:       fa.IsActive,
		}
	}

	return config, nil
}
