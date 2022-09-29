package entity

type Account struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Currency string `gorm:"not null" json:"currency"`
	Amount   uint64 `gorm:"not null" json:"amount"`
	UserID   uint64 `gorm:"not null" json:"-"`
	User     User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
