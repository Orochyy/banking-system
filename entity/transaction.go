package entity

type Transaction struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Amount   uint64  `gorm:"not null" json:"amount"`
	Currency string  `gorm:"not null" json:"currency"`
	Account  Account `gorm:"foreignkey:AccountID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"account"`
}
