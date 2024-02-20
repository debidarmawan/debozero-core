package model

type Role struct {
	BaseModel
	Code string `gorm:"not null;type:varchar(32);unique"`
	Name string `gorm:"not null;type:varchar(32);unique"`
}
