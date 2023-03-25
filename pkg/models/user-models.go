package models

import "github.com/salmanfr/catatann-api/pkg/entities"

type GoogleOAuthSigninDto struct {
	ClientId string `json:"client_id"`
	Credential string `json:"credential"`
}

type SignupResponse struct {
	User entities.User `json:"user"`
}

type SigninResponse struct {
	User entities.User `json:"user"`
	AccessToken string `json:"access_token"`
	RefreshToken string
	ExtRefreshToken string
}

type ExtensionSigninResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignupDto struct {
	Email string `json:"email" validate:"required,email,max=255"`
	FullName string `json:"full_name" validate:"required,min=6,max=255"`
	Password string `json:"password" validate:"required,min=6,max=30"`
}

type SigninDto struct {
	Email string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=6,max=30"`
}

type ExtensionRefreshTokenDto struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

