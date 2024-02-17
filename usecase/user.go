package usecase

import (
	"github.com/debidarmawan/debozero-core/helper"
	"github.com/debidarmawan/debozero-core/repository"
)

type UserUseCase interface{}

type userUseCase struct {
	txManager helper.TxManager
	userRepo  repository.UserRepo
}

func NewUserUseCase(
	txManager helper.TxManager,
	userRepo repository.UserRepo,
) UserUseCase {
	return &userUseCase{
		txManager: txManager,
		userRepo:  userRepo,
	}
}
