package remote

import (
	"context"
	"fmt"

	pbCatalogue "qb-accounting/internal/proto/catalogue"

	qbgrpc "github.com/MapleGraph/qb-core/v2/pkg/grpc"
)

// ServiceNameCatalogue is the service name for the catalogue service
const ServiceNameCatalogue = "catalogue"

// CatalogueService defines the interface for catalogue service operations
type CatalogueService interface {
	GetProductsByIDs(ctx context.Context, organizationID string, productIDs []string) ([]*Product, error)
	GetProductGroupsByIDs(ctx context.Context, organizationID string, productGroupIDs []string) ([]*ProductGroup, error)
	GetMeasurementUnits(ctx context.Context, organizationID string) ([]*MeasurementUnit, error)
}

// Product represents a product from the catalogue service
type Product struct {
	ProductID                 string
	OrganizationID            string
	ProductName               string
	ProductType               string
	ShortDescription          *string
	LongDescription           *string
	Images                    []string
	ColourCode                *string
	MeasurementUnitID         *string
	OperationalUnitID         *string
	Classification            *string
	IsSellableOnPos           bool
	AllowVariablePricing      bool
	PriceIncludesTaxes        bool
	DefaultSellingPrice       float64
	EnableProductGroup        bool
	PricingAttribute          *string
	IsInventoryManaged        bool
	Barcode                   *string
	Sku                       *string
	DimensionUnit             *string
	PackagingInfo             *string
	InventoryReductionMethod  *string
	RecipeID                  *string
	DefaultStockAlertLevel    *int32
	DefaultSafeStockLevel     *int32
	CountryOfOrigin           *string
	HsnSacCode                *string
	DisableSalesOnNoStock     bool
	DisableSalesOnStockExpiry bool
	ReturnExchange            *string
	Attributes                *string
	Relationships             *string
	ItemTypeAttributes        *string
	CreatedBy                 *string
	UpdatedBy                 *string
	CreatedAtLocal            string
	CreatedAtUtc              string
	UpdatedAtLocal            *string
	UpdatedAtUtc              *string
	DemoMode                  bool
	IsActive                  int32
}

type catalogueRepository struct {
	handler qbgrpc.ClientHandler
}

// NewCatalogueRepository creates a new catalogue repository using a qb-core gRPC handler.
func NewCatalogueRepository(handler qbgrpc.ClientHandler) CatalogueService {
	return &catalogueRepository{handler: handler}
}

func (r *catalogueRepository) client(ctx context.Context) (pbCatalogue.ProductServiceClient, error) {
	if r.handler == nil {
		return nil, fmt.Errorf("catalogue gRPC handler is not available")
	}
	conn, err := r.handler.GetConnectionWithRetry(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get catalogue gRPC connection: %w", err)
	}
	return pbCatalogue.NewProductServiceClient(conn), nil
}

// GetProductsByIDs fetches products by their IDs from the catalogue service
func (r *catalogueRepository) GetProductsByIDs(ctx context.Context, organizationID string, productIDs []string) ([]*Product, error) {
	client, err := r.client(ctx)
	if err != nil {
		return nil, err
	}

	if len(productIDs) == 0 {
		return []*Product{}, nil
	}

	req := &pbCatalogue.GetProductsByIDsRequest{
		OrganizationId: organizationID,
		ProductIds:     productIDs,
	}

	resp, err := client.GetProductsByIDs(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetProductsByIDs: %w", err)
	}

	// Map proto products to local Product structs
	products := make([]*Product, 0, len(resp.Products))
	for _, p := range resp.Products {
		products = append(products, mapProtoToProduct(p))
	}

	return products, nil
}

// GetProductGroupsByIDs fetches product groups by ID.
func (r *catalogueRepository) GetProductGroupsByIDs(ctx context.Context, organizationID string, productGroupIDs []string) ([]*ProductGroup, error) {
	client, err := r.client(ctx)
	if err != nil {
		return nil, err
	}

	if len(productGroupIDs) == 0 {
		return []*ProductGroup{}, nil
	}

	req := &pbCatalogue.GetProductGroupsByIDsRequest{
		OrganizationId:  organizationID,
		ProductGroupIds: productGroupIDs,
	}

	resp, err := client.GetProductGroupsByIDs(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetProductGroupsByIDs: %w", err)
	}

	groupResults := make([]*ProductGroup, 0, len(resp.ProductGroups))
	for _, pg := range resp.ProductGroups {
		groupResults = append(groupResults, mapProtoToProductGroup(pg))
	}

	return groupResults, nil
}

// GetMeasurementUnits fetches measurement units from the catalogue service.
func (r *catalogueRepository) GetMeasurementUnits(ctx context.Context, organizationID string) ([]*MeasurementUnit, error) {
	client, err := r.client(ctx)
	if err != nil {
		return nil, err
	}

	req := &pbCatalogue.GetMeasurementUnitsRequest{
		OrganizationId: organizationID,
	}

	resp, err := client.GetMeasurementUnits(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to call GetMeasurementUnits: %w", err)
	}

	units := make([]*MeasurementUnit, 0, len(resp.MeasurementUnits))
	for _, mu := range resp.MeasurementUnits {
		units = append(units, mapProtoToMeasurementUnit(mu))
	}

	return units, nil
}

// ProductGroup represents catalogue product group metadata.
type ProductGroup struct {
	ProductGroupID           string
	OrganizationID           string
	ProductGroupName         string
	ProductID                string
	ShortDescription         *string
	LongDescription          *string
	PgStatus                 string
	Images                   []string
	PgBarcode                *string
	PgSku                    *string
	ManufacturingDate        *string
	ExpiryDate               *string
	ShelfLife                *int32
	CostPerUnit              *float64
	FulfillmentConfiguration *string
	Attributes               *string
	SellByDate               *string
	Mrp                      *float64
	ReturnConfiguration      *string
	ExchangeConfiguration    *string
	CreatedAt                *string
	UpdatedAt                *string
}

// MeasurementUnit mirrors catalogue measurement unit metadata.
type MeasurementUnit struct {
	UnitID           string
	OrganizationID   string
	ParentUnitID     *string
	Name             string
	PrintName        string
	Description      string
	IsUneditable     bool
	IsPrincipalUnit  bool
	ConversionFactor *string
	CreatedAtLocal   *string
	CreatedAtUtc     *string
	UpdatedAtLocal   *string
	UpdatedAtUtc     *string
	IsActive         int32
	DemoMode         bool
}

func mapProtoToProductGroup(pg *pbCatalogue.ProductGroup) *ProductGroup {
	if pg == nil {
		return nil
	}
	return &ProductGroup{
		ProductGroupID:           pg.ProductGroupId,
		OrganizationID:           pg.OrganizationId,
		ProductGroupName:         pg.ProductGroupName,
		ProductID:                pg.ProductId,
		ShortDescription:         pg.ShortDescription,
		LongDescription:          pg.LongDescription,
		PgStatus:                 pg.PgStatus,
		Images:                   pg.Images,
		PgBarcode:                pg.PgBarcode,
		PgSku:                    pg.PgSku,
		ManufacturingDate:        pg.ManufacturingDate,
		ExpiryDate:               pg.ExpiryDate,
		ShelfLife:                pg.ShelfLife,
		CostPerUnit:              pg.CostPerUnit,
		FulfillmentConfiguration: pg.FulfillmentConfiguration,
		Attributes:               pg.Attributes,
		SellByDate:               pg.SellByDate,
		Mrp:                      pg.Mrp,
		ReturnConfiguration:      pg.ReturnConfiguration,
		ExchangeConfiguration:    pg.ExchangeConfiguration,
		CreatedAt:                pg.CreatedAt,
		UpdatedAt:                pg.UpdatedAt,
	}
}

func mapProtoToMeasurementUnit(mu *pbCatalogue.MeasurementUnit) *MeasurementUnit {
	if mu == nil {
		return nil
	}
	return &MeasurementUnit{
		UnitID:           mu.UnitId,
		OrganizationID:   mu.OrganizationId,
		ParentUnitID:     mu.ParentUnitId,
		Name:             mu.Name,
		PrintName:        mu.PrintName,
		Description:      mu.Description,
		IsUneditable:     mu.IsUneditable,
		IsPrincipalUnit:  mu.IsPrincipalUnit,
		ConversionFactor: mu.ConversionFactor,
		CreatedAtLocal:   mu.CreatedAtLocal,
		CreatedAtUtc:     mu.CreatedAtUtc,
		UpdatedAtLocal:   mu.UpdatedAtLocal,
		UpdatedAtUtc:     mu.UpdatedAtUtc,
		IsActive:         mu.IsActive,
		DemoMode:         mu.DemoMode,
	}
}

// mapProtoToProduct maps a proto Product to a local Product struct
func mapProtoToProduct(p *pbCatalogue.Product) *Product {
	product := &Product{
		ProductID:                 p.ItemId,
		OrganizationID:            p.OrganizationId,
		ProductName:               p.ProductName,
		ProductType:               p.ProductType,
		Images:                    p.Images,
		IsSellableOnPos:           p.IsSellableOnPos,
		AllowVariablePricing:      p.AllowVariablePricing,
		PriceIncludesTaxes:        p.PriceIncludesTaxes,
		DefaultSellingPrice:       p.DefaultSellingPrice,
		EnableProductGroup:        p.EnableProductGroup,
		IsInventoryManaged:        p.IsInventoryManaged,
		DisableSalesOnNoStock:     p.DisableSalesOnNoStock,
		DisableSalesOnStockExpiry: p.DisableSalesOnStockExpiry,
		DemoMode:                  p.DemoMode,
		IsActive:                  p.IsActive,
		CreatedAtLocal:            p.CreatedAtLocal,
		CreatedAtUtc:              p.CreatedAtUtc,
	}

	// Map optional string fields
	if p.ShortDescription != nil {
		product.ShortDescription = p.ShortDescription
	}
	if p.LongDescription != nil {
		product.LongDescription = p.LongDescription
	}
	if p.ColourCode != nil {
		product.ColourCode = p.ColourCode
	}
	if p.MeasurementUnitId != nil {
		product.MeasurementUnitID = p.MeasurementUnitId
	}
	if p.OperationalUnitId != nil {
		product.OperationalUnitID = p.OperationalUnitId
	}
	if p.Classification != nil {
		product.Classification = p.Classification
	}
	if p.PricingAttribute != nil {
		product.PricingAttribute = p.PricingAttribute
	}
	if p.Barcode != nil {
		product.Barcode = p.Barcode
	}
	if p.Sku != nil {
		product.Sku = p.Sku
	}
	if p.DimensionUnit != nil {
		product.DimensionUnit = p.DimensionUnit
	}
	if p.PackagingInfo != nil {
		product.PackagingInfo = p.PackagingInfo
	}
	if p.InventoryReductionMethod != nil {
		product.InventoryReductionMethod = p.InventoryReductionMethod
	}
	if p.RecipeId != nil {
		product.RecipeID = p.RecipeId
	}
	if p.CountryOfOrigin != nil {
		product.CountryOfOrigin = p.CountryOfOrigin
	}
	if p.HsnSacCode != nil {
		product.HsnSacCode = p.HsnSacCode
	}
	if p.ReturnExchange != nil {
		product.ReturnExchange = p.ReturnExchange
	}
	if p.Attributes != nil {
		product.Attributes = p.Attributes
	}
	if p.Relationships != nil {
		product.Relationships = p.Relationships
	}
	if p.ItemTypeAttributes != nil {
		product.ItemTypeAttributes = p.ItemTypeAttributes
	}
	if p.CreatedBy != nil {
		product.CreatedBy = p.CreatedBy
	}
	if p.UpdatedBy != nil {
		product.UpdatedBy = p.UpdatedBy
	}
	if p.UpdatedAtLocal != nil {
		product.UpdatedAtLocal = p.UpdatedAtLocal
	}
	if p.UpdatedAtUtc != nil {
		product.UpdatedAtUtc = p.UpdatedAtUtc
	}

	// Map optional int32 fields
	if p.DefaultStockAlertLevel != nil {
		product.DefaultStockAlertLevel = p.DefaultStockAlertLevel
	}
	if p.DefaultSafeStockLevel != nil {
		product.DefaultSafeStockLevel = p.DefaultSafeStockLevel
	}

	return product
}
