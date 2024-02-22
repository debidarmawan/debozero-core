package repository

import (
	"github.com/debidarmawan/debozero-core/global"
	"github.com/debidarmawan/debozero-core/model"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAllowedRoleKeyListByRoleId(roleId string) (*[]model.RoleKey, global.ErrorResponse)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (rr *roleRepository) GetAllowedRoleKeyListByRoleId(roleId string) (*[]model.RoleKey, global.ErrorResponse) {
	var roleKeys []model.RoleKey

	result := rr.db.
		Table("role_key").
		Joins("LEFT JOIN role_key_mapping ON role_key_mapping.role_key_id = role_key.id").
		Joins("LEFT JOIN role ON role_key_mapping.role_id = role.id").
		Where("role.id = ? AND role_key_mapping.is_allow = ?", roleId, true).
		Find(&roleKeys)

	if result.Error != nil {
		return nil, global.InternalServerError(result.Error)
	}

	return &roleKeys, nil
}
