package dto

type TransactionCreateDTO struct {
	AccountID uint64 `json:"account_id,omitempty" form:"account_id" binding:"account_id,omitempty"`
	Amount    uint64 `json:"amount" form:"amount" binding:"required"`
	Currency  string `json:"currency" form:"currency" binding:"required"`
}
