package model

import "time"

type User struct {
	AdminBaseModel
	Name         string `gorm:"not null; type:varchar(255)"`
	Username     string `gorm:"not null; type:varchar(255)"`
	Email        string `gorm:"not null; type:varchar(255); uniqueIndex"`
	Password     string `gorm:"not null; type:varchar(255)"`
	Phone        string `gorm:"type:varchar(30)"`
	RoleId       string `gorm:"not null"`
	Role         Role
	IsActive     bool       `gorm:""`
	LastLoggedIn *time.Time `gorm:""`
}
