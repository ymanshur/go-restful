package dto

import "go-restful/internal/model"

type (
	CreateUserRequest struct {
		model.User
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
