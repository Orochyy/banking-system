package dto

type TransactionCreateDTO struct {
	Amount           uint64 `json:"amount" form:"amount" binding:"required"`
	AccountSender    uint64 `gorm:"foreignkey:AccountSenderID" json:"account_sender" binding:"required"`
	AccountRecipient uint64 `gorm:"foreignkey:AccountRecipientID" json:"account_recipient" binding:"required"`
	Currency         string `json:"currency" form:"currency"  binding:"required"`
	Type             string `json:"type" form:"type" binding:"required" validate:"oneof=deposit money,withdraw money"`
	UserID           uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
