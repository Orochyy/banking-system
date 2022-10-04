package dto

type TransactionCreateDTO struct {
	Amount   uint64 `json:"amount" form:"amount" binding:"required"`
	Currency string `json:"currency" form:"currency" binding:"required"`
}
