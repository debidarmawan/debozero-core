package dto

import (
	"net/http"
	"time"
)

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

type Logout struct {
	Request *http.Request
}

type Verify struct {
	Request *http.Request
	Path    string
	Method  string
}

type VerifyHeader struct {
	Path   string `reqHeader:"X-Path" validate:"required"`
	Method string `reqHeader:"X-Method" validate:"required"`
}

type VerifyResult struct {
	UserId string
	Scope  string
}

type VerifyResponse struct {
	UserId string `json:"user_id"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenSpec struct {
	RefreshToken string
}

type RefreshTokenResponse struct {
	AccessToken  string    `json:"access_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	RefreshToken string    `json:"refresh_token"`
}
