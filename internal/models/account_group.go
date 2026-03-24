package models

type AccountNature string

const (
	AccountNatureAsset     AccountNature = "ASSET"
	AccountNatureLiability AccountNature = "LIABILITY"
	AccountNatureEquity    AccountNature = "EQUITY"
	AccountNatureIncome    AccountNature = "INCOME"
	AccountNatureExpense   AccountNature = "EXPENSE"
)

type AccountGroup struct {
	ID            string        `json:"id" db:"id" db_pk:"true"`
	BookID        string        `json:"book_id" db:"book_id"`
	Code          string        `json:"code" db:"code"`
	Name          string        `json:"name" db:"name"`
	ParentGroupID *string       `json:"parent_group_id" db:"parent_group_id"`
	AccountNature AccountNature `json:"account_nature" db:"account_nature"`
	SortOrder     int           `json:"sort_order" db:"sort_order"`
	IsSystem      bool          `json:"is_system" db:"is_system"`
	BaseModel
}
