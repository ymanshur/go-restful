package dto

import "go-restful/internal/model"

type (
	CreateUserRequest struct {
		model.User
	}
	UpdateUserRequest struct {
		ByIDRequest
		Name  *string `json:"name" validate:"omitempty"`
		Email *string `json:"email" validate:"omitempty,email"`
	}
	UserResponse struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	UserWithTokenResponse struct {
		UserResponse
		Token string `json:"token"`
	}
)
