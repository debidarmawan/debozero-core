package repository

import (
	"debozero-core/global"
	"debozero-core/helper"
	"debozero-core/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*model.User, global.ErrorResponse)
	GetUserByUsername(username string) (*model.User, global.ErrorResponse)
	CreateUser(tx helper.Tx, user *model.User) (*model.User, global.ErrorResponse)
	GetUserById(id string) (*model.User, global.ErrorResponse)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetUserByEmail(email string) (*model.User, global.ErrorResponse) {
	var user model.User

	err := ur.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, global.NotFoundError("User not found")
	}

	return &user, nil
}

func (ur *userRepository) GetUserByUsername(username string) (*model.User, global.ErrorResponse) {
	var user model.User

	err := ur.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, global.NotFoundError("User not found")
	}

	return &user, nil
}

func (ur *userRepository) CreateUser(tx helper.Tx, user *model.User) (*model.User, global.ErrorResponse) {
	db := ur.db

	if tx != nil {
		db = tx.Get()
	}

	err := db.Create(user).Error
	if err != nil {
		return nil, global.InternalServerError(err)
	}

	return user, nil
}

func (ur *userRepository) GetUserById(id string) (*model.User, global.ErrorResponse) {
	var user model.User

	err := ur.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, global.NotFoundError()
	}

	return &user, nil
}
