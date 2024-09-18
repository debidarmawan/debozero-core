package dto

type UserRegisterRequest struct {
	Name            string `json:"name" validate:"required,min=3,max=60" example:"William"`
	Username        string `json:"username" validate:"required,min=3,max=60" example:"william"`
	Email           string `json:"email" validate:"required,email" example:"william@debozero.id"`
	Phone           string `json:"phone" validate:"required" example:"08218833123"`
	Password        string `json:"password" validate:"required,min=6,max=30" example:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,max=30" example:"password"`
}

type UserValidationFieldErr struct {
	Field        string `json:"field"`
	ErrorMessage string `json:"error_message"`
}
