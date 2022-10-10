package dto

import "go-restful/internal/model"

type (
	AuthSignUpRequest struct {
		model.User
	}

	AuthSignInRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	AuthSignInResponse struct {
		UserResponse
		Token string `json:"token"`
	}
)
