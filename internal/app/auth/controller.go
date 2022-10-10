package auth

import (
	"errors"
	"go-restful/internal/app/dto"
	"go-restful/internal/pkg/util"
	"go-restful/internal/repository"
	res "go-restful/pkg/util/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Controller interface {
		SignUp(ctx echo.Context) error
		SignIn(ctx echo.Context) error

		Route(g *echo.Group)
	}
	controller struct {
		repo repository.User
	}
)

func NewController(r repository.User) Controller {
	return &controller{
		repo: r,
	}
}

func (c *controller) SignUp(ctx echo.Context) error {
	payload := new(dto.CreateUserRequest)
	if err := ctx.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(ctx)
	}

	if err := ctx.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err).Send(ctx)
	}

	if err := c.repo.Save(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return res.CustomSuccessBuilder(
		http.StatusOK, nil, "thanks for registering",
	).Send(ctx)
}

func (c *controller) SignIn(ctx echo.Context) error {
	payload := new(dto.AuthSignInRequest)
	if err := ctx.Bind(&payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(ctx)
	}

	if err := ctx.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
	}

	// Find user by email
	user, err := c.repo.FindByEmail(&payload.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(ctx)
		}
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(ctx)
	}

	// Matching password
	if payload.Password != user.Password {
		return res.ErrorBuilder(&res.ErrorConstant.EmailOrPasswordIncorrect, err)
	}

	// Create JWT token
	token, err := util.CreateJwt(user)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(ctx)
	}

	return res.CustomSuccessBuilder(
		http.StatusOK,
		dto.UserWithTokenResponse{
			UserResponse: dto.UserResponse{
				Name:  user.Name,
				Email: user.Email,
			},
			Token: token,
		},
		"success logged in",
	).Send(ctx)
}
