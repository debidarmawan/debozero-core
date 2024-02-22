package usecase

import (
	"github.com/debidarmawan/debozero-core/dto"
	"github.com/debidarmawan/debozero-core/global"
	"github.com/debidarmawan/debozero-core/helper"
	"github.com/debidarmawan/debozero-core/model"
	"github.com/debidarmawan/debozero-core/repository"
)

type UserUseCase interface {
	Register(request dto.UserRegisterRequest) (*model.User, global.ErrorResponse)
}

type userUseCase struct {
	txManager helper.TxManager
	userRepo  repository.UserRepository
}

func NewUserUseCase(
	txManager helper.TxManager,
	userRepo repository.UserRepository,
) UserUseCase {
	return &userUseCase{
		txManager: txManager,
		userRepo:  userRepo,
	}
}

func (uu *userUseCase) Register(request dto.UserRegisterRequest) (*model.User, global.ErrorResponse) {
	hashedPassword, err := uu.validateUserRegistrationData(request)
	if err != nil {
		return nil, err
	}

	tx := uu.txManager.New()
	defer tx.CheckPanic()

	user := &model.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: *hashedPassword,
		Phone:    request.Phone,
		IsActive: true,
		RoleId:   "9a428ae5-fade-4605-ac3d-41c324455861",
	}

	user, err = uu.userRepo.CreateUser(tx, user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, global.InternalServerError(err)
	}

	return user, nil
}

func (uu *userUseCase) validateUserRegistrationData(request dto.UserRegisterRequest) (*string, global.ErrorResponse) {
	errMessages := make([]dto.UserValidationFieldErr, 0)

	user, _ := uu.userRepo.GetUserByEmail(request.Email)
	if user != nil {
		errMessage := dto.UserValidationFieldErr{
			Field:        "email",
			ErrorMessage: "Email is invalid or already registered",
		}

		errMessages = append(errMessages, errMessage)
	}

	user, _ = uu.userRepo.GetUserByUsername(request.Username)
	if user != nil {
		errMessage := dto.UserValidationFieldErr{
			Field:        "username",
			ErrorMessage: "Username is already registered",
		}

		errMessages = append(errMessages, errMessage)
	}

	if request.Password != request.ConfirmPassword {
		errMessage := dto.UserValidationFieldErr{
			Field:        "password",
			ErrorMessage: "Password confirmation is missmatch",
		}

		errMessages = append(errMessages, errMessage)
	}

	hashedPassword := helper.HashPassword(request.Password)

	if len(errMessages) > 0 {
		return nil, nil
	}

	return &hashedPassword, nil
}
