package dto

type CreateBookRequest struct {
	CompanyID             string  `json:"company_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440001"`
	Code                  string  `json:"code" validate:"required,max=30" example:"PRIMARY"`
	Name                  string  `json:"name" validate:"required,max=120" example:"Primary Book"`
	BookType              string  `json:"book_type" validate:"required,oneof=PRIMARY STATUTORY MANAGEMENT" example:"PRIMARY"`
	BaseCurrencyCode      string  `json:"base_currency_code" validate:"required,len=3" example:"INR"`
	ReportingCurrencyCode *string `json:"reporting_currency_code" validate:"omitempty,len=3" example:"USD"`
}

type UpdateBookRequest struct {
	Name                  *string `json:"name" validate:"omitempty,max=120" example:"Primary Book Updated"`
	ReportingCurrencyCode *string `json:"reporting_currency_code" validate:"omitempty,len=3" example:"USD"`
	Status                *string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE" example:"ACTIVE"`
}

type BookResponse struct {
	ID                    string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	CompanyID             string  `json:"company_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	Code                  string  `json:"code" example:"PRIMARY"`
	Name                  string  `json:"name" example:"Primary Book"`
	BookType              string  `json:"book_type" example:"PRIMARY"`
	BaseCurrencyCode      string  `json:"base_currency_code" example:"INR"`
	ReportingCurrencyCode *string `json:"reporting_currency_code" example:"USD"`
	Status                string  `json:"status" example:"ACTIVE"`
}
