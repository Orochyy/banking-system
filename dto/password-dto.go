package dto

type ManagerUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `gorm:"not null" json:"name"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

type ManagerCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}
