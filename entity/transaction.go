package entity

type Transaction struct {
	ID               uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Amount           uint64 `gorm:"not null" json:"amount"`
	Currency         string `gorm:"not null" json:"currency"`
	AccountSender    uint64 `gorm:"foreignkey:AccountSenderID" json:"account_sender"`
	AccountRecipient uint64 `gorm:"foreignkey:AccountRecipientID" json:"account_recipient"`
	Type             string `gorm:"not null" json:"type"`
	UserID           uint64 `gorm:"not null" json:"-"`
	User             User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
