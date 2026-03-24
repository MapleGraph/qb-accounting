package dto

type CreateAccountGroupRequest struct {
	BookID        string  `json:"book_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440010"`
	Code          string  `json:"code" validate:"required,max=30" example:"CA"`
	Name          string  `json:"name" validate:"required,max=200" example:"Current Assets"`
	ParentGroupID *string `json:"parent_group_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440040"`
	AccountNature string  `json:"account_nature" validate:"required,oneof=ASSET LIABILITY EQUITY INCOME EXPENSE" example:"ASSET"`
	SortOrder     int     `json:"sort_order" validate:"required,min=0" example:"10"`
	IsSystem      bool    `json:"is_system" example:"false"`
}

type UpdateAccountGroupRequest struct {
	Name          *string `json:"name" validate:"omitempty,max=200" example:"Current Assets (Operating)"`
	ParentGroupID *string `json:"parent_group_id" validate:"omitempty,uuid" example:"550e8400-e29b-41d4-a716-446655440040"`
	SortOrder     *int    `json:"sort_order" validate:"omitempty,min=0" example:"15"`
}

type AccountGroupResponse struct {
	ID            string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440050"`
	BookID        string  `json:"book_id" example:"550e8400-e29b-41d4-a716-446655440010"`
	Code          string  `json:"code" example:"CA"`
	Name          string  `json:"name" example:"Current Assets"`
	ParentGroupID *string `json:"parent_group_id" example:"550e8400-e29b-41d4-a716-446655440040"`
	AccountNature string  `json:"account_nature" example:"ASSET"`
	SortOrder     int     `json:"sort_order" example:"10"`
	IsSystem      bool    `json:"is_system" example:"false"`
}
