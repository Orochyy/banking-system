package entity

type Transaction struct {
	ID               uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Amount           uint64  `gorm:"not null" json:"amount"`
	Currency         string  `gorm:"not null" json:"currency"`
	Type             string  `gorm:"not null" json:"type"`
	AccountSender    Account `gorm:"foreignkey:AccountSenderID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"account_sender"`
	AccountRecipient Account `gorm:"foreignkey:AccountRecipientID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"account_recipient"`
}
