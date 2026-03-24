package models

type BookType string

const (
	BookTypePrimary    BookType = "PRIMARY"
	BookTypeStatutory  BookType = "STATUTORY"
	BookTypeManagement BookType = "MANAGEMENT"
)

type BookStatus string

const (
	BookStatusActive   BookStatus = "ACTIVE"
	BookStatusInactive BookStatus = "INACTIVE"
)

type Book struct {
	ID                    string     `json:"id" db:"id" db_pk:"true"`
	CompanyID             string     `json:"company_id" db:"company_id"`
	Code                  string     `json:"code" db:"code"`
	Name                  string     `json:"name" db:"name"`
	BookType              BookType   `json:"book_type" db:"book_type"`
	BaseCurrencyCode      string     `json:"base_currency_code" db:"base_currency_code"`
	ReportingCurrencyCode *string    `json:"reporting_currency_code" db:"reporting_currency_code"`
	Status                BookStatus `json:"status" db:"status"`
	BaseModel
}
