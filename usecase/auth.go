package usecase

import (
	"github.com/debidarmawan/debozero-core/dto"
	"github.com/debidarmawan/debozero-core/global"
	"github.com/debidarmawan/debozero-core/helper"
	"github.com/debidarmawan/debozero-core/model"
	"github.com/debidarmawan/debozero-core/repository"
)

type AuthUseCase interface {
	Login(request dto.Login) (*dto.LoginResponse, global.ErrorResponse)
	Logout(request dto.Logout) global.ErrorResponse
}

type authUseCase struct {
	txManager     helper.TxManager
	userRepo      repository.UserRepo
	oauth2UseCase Oauth2UseCase
}

func NewAuthUseCase(
	txManager helper.TxManager,
	userRepo repository.UserRepo,
	oauth2UseCase Oauth2UseCase,
) AuthUseCase {
	return &authUseCase{
		txManager:     txManager,
		userRepo:      userRepo,
		oauth2UseCase: oauth2UseCase,
	}
}

func (au *authUseCase) Login(request dto.Login) (*dto.LoginResponse, global.ErrorResponse) {
	var (
		user *model.User
		err  global.ErrorResponse
	)

	user, err = au.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return nil, global.BadRequestError("Email or Password is invalid")
	}

	isValidPassword := helper.CheckPasswordHash(request.Password, user.Password)
	if !isValidPassword {
		return nil, global.BadRequestError("Email or Password is invalid")
	}

	tokenInfo, error := au.oauth2UseCase.GenerateToken(user.Id, "")
	if error != nil {
		return nil, global.InternalServerError(error)
	}

	loginResponse := dto.LoginResponse{
		AccessToken:  tokenInfo.AccessToken,
		ExpiresAt:    tokenInfo.ExpiresAt,
		RefreshToken: tokenInfo.RefreshToken,
		UserEmail:    user.Email,
		Username:     user.Username,
		Name:         user.Name,
		Role:         user.Role.Code,
	}

	return &loginResponse, nil
}

func (au *authUseCase) Logout(request dto.Logout) global.ErrorResponse {
	err := au.oauth2UseCase.RemoveToken(request.Request)
	if err != nil {
		return global.InternalServerError(err)
	}

	return nil
}
