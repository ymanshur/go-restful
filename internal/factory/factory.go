package factory

import (
	"go-restful/database"
	"go-restful/internal/model"
	"go-restful/internal/repository"
	"go-restful/pkg/constant"
)

type Factory struct {
	UserRepository *repository.User
}

func NewFactory() *Factory {
	// Define databases
	userDB := database.New(database.Config{
		User:     constant.Env.Get("DB_USER", ""),
		Password: constant.Env.Get("DB_PASS", ""),
		Host:     constant.Env.Get("DB_HOST", ""),
		Port:     constant.Env.Get("DB_PORT", ""),
		Name:     constant.Env.Get("DB_NAME", ""),
	})
	database.Load(userDB, &model.User{})

	return &Factory{
		repository.NewUserRepository(userDB),
	}
}
