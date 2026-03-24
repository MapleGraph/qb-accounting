package models

import "time"

type BaseModel struct {
	CreatedBy *string    `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at" db_auto:"true"`
	UpdatedBy *string    `json:"updated_by" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at" db_auto:"true"`
}
