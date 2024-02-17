package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        string `gorm:"type:varchar(36);primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	base.Id = uuid.NewString()
	return nil
}

type AdminBaseModel struct {
	BaseModel
	CreatedBy string `gorm:"default:null;type:varchar(36)"`
	UpdatedBy string `gorm:"default:null;type:varchar(36)"`
	DeletedBy string `gorm:"default:null;type:varchar(36)"`
}
