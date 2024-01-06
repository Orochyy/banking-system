package entity

type Password struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Password string `gorm:"->;<-;not null" json:"password"`
}

type Manager struct {
	ID   uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"not null" json:"name"`
}
