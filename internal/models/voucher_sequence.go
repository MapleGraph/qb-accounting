package models

type VoucherSequence struct {
	ID          string  `json:"id" db:"id" db_pk:"true"`
	BookID      string  `json:"book_id" db:"book_id"`
	CompanyID   string  `json:"company_id" db:"company_id"`
	BranchID    *string `json:"branch_id" db:"branch_id"`
	VoucherType string  `json:"voucher_type" db:"voucher_type"`
	Prefix      *string `json:"prefix" db:"prefix"`
	Suffix      *string `json:"suffix" db:"suffix"`
	Padding     int16   `json:"padding" db:"padding"`
	NextNumber  int64   `json:"next_number" db:"next_number"`
	ResetPolicy string  `json:"reset_policy" db:"reset_policy"`
	IsActive    bool    `json:"is_active" db:"is_active"`
	BaseModel
}
