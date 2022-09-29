package dto

type AccountUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Currency string `json:"currency" form:"currency" binding:"required"`
	Amount   uint64 `json:"amount" form:"amount" binding:"required"`
	UserID   uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type AccountCreateDTO struct {
	Currency string `json:"currency" form:"currency" binding:"required"`
	Amount   uint64 `json:"amount" form:"amount" binding:"required"`
	UserID   uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
