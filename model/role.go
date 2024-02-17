package model

type Role struct {
	BaseModel
	Name string `gorm:"not null;type:varchar(32);unique"`
}
