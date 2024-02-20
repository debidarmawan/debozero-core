package repository

import (
	"github.com/debidarmawan/debozero-core/global"
	"github.com/debidarmawan/debozero-core/helper"
	"github.com/debidarmawan/debozero-core/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetUserByEmail(email string) (*model.User, global.ErrorResponse)
	GetUserByUsername(username string) (*model.User, global.ErrorResponse)
	CreateUser(tx helper.Tx, user *model.User) (*model.User, global.ErrorResponse)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) GetUserByEmail(email string) (*model.User, global.ErrorResponse) {
	var user model.User

	err := ur.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, global.NotFoundError("User not found")
	}

	return &user, nil
}

func (ur *userRepo) GetUserByUsername(username string) (*model.User, global.ErrorResponse) {
	var user model.User

	err := ur.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, global.NotFoundError("User not found")
	}

	return &user, nil
}

func (ur *userRepo) CreateUser(tx helper.Tx, user *model.User) (*model.User, global.ErrorResponse) {
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
