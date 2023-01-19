package dto

type AccountUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Currency string `json:"currency" form:"currency" binding:"required"`
	Amount   uint64 `json:"amount" form:"amount" binding:"required"`
	UserID   uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	Hex      string `json:"hex,omitempty"  form:"hex,omitempty"`
}

type AccountCreateDTO struct {
	Currency string `json:"currency" form:"currency" binding:"required"`
	Amount   uint64 `json:"amount" form:"amount" binding:"required"`
	UserID   uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	Hex      string `json:"hex,omitempty"  form:"hex,omitempty"`
}

type AccountUpdateAmountDTO struct {
	//ID       uint64 `json:"id" form:"id" binding:"required"`
	//Currency string `gorm:"not null" json:"currency"`
	Amount uint64 `json:"amount" form:"amount" binding:"required"`
	UserID uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	//Hex      string `gorm:"not null" json:"hex"`
}

type AccountHexDTO struct {
	Hex string `json:"hex" form:"hex" binding:"required"`
}
