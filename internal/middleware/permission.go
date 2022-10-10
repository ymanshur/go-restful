package middleware

import (
	"fmt"
	inutil "go-restful/internal/pkg/util"
	"go-restful/pkg/constant"
	res "go-restful/pkg/util/response"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &inutil.JwtClaims{},
		SigningKey: []byte(constant.Env.Get("JWT_SECRET", "")),
	})
)

func IsAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId, err := strconv.Atoi(ctx.Param("id"))
		if err == nil {
			user := ctx.Get("user").(*jwt.Token)
			claims := user.Claims.(*inutil.JwtClaims)
			if claims.UserId != uint(userId) {
				return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, fmt.Errorf("unauthorized user id")).Send(ctx)
			}
		}

		return next(ctx)
	}
}
