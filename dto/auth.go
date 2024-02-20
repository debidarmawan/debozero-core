package dto

import "time"

type Login struct {
	Email    string `json:"email" validate:"required" example:"user@email.com"`
	Password string `json:"password" validate:"required" example:"password"`
}

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	RefreshToken string    `json:"refresh_token"`
	UserEmail    string    `json:"email"`
	Username     string    `json:"username"`
	Name         string    `json:"name"`
	Role         string    `json:"role"`
}

type TokenInfo struct {
	AccessToken  string    `json:"access_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	RefreshToken string    `json:"refresh_token"`
}
