package http

import (
	"go-restful/internal/app/auth"
	"go-restful/internal/app/user"
	"go-restful/internal/factory"

	"github.com/labstack/echo/v4"
)

func New(g *echo.Group, f *factory.Factory) {
	user.NewController(f.UserRepository).Route(g.Group("/users"))
	auth.NewController(f.UserRepository).Route(g.Group("/auth"))
}
