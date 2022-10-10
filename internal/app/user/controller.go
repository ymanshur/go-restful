package user

import (
	"go-restful/internal/dto"
	"go-restful/internal/repository"
	res "go-restful/pkg/util/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	Controller interface {
		Get(ctx echo.Context) error
		Update(ctx echo.Context) error
		Delete(ctx echo.Context) error
		GetAll(ctx echo.Context) error
	}
	controller struct {
		repo repository.User
	}
)

func NewController(r repository.User) *controller {
	return &controller{
		repo: r,
	}
}

func (c *controller) Get(ctx echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := ctx.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(ctx)
	}

	if err := ctx.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.ValidationError, err).Send(ctx)
	}

	user, err := c.repo.FindById(payload.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(ctx)
		}
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(ctx)
	}

	return res.CustomSuccessBuilder(
		http.StatusOK, user, "success get a user",
	).Send(ctx)
}

func (c *controller) Update(ctx echo.Context) error {
	payload := new(dto.UpdateUserRequest)
	if err := ctx.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(ctx)
	}

	if err := ctx.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.ValidationError, err).Send(ctx)
	}

	user, err := c.repo.FindById(payload.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(ctx)
		}
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(ctx)
	}

	updatedUser, err := c.repo.UpdateById(user, payload)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(ctx)
	}

	return res.CustomSuccessBuilder(
		http.StatusOK, updatedUser, "success update a user",
	).Send(ctx)
}

func (c *controller) Delete(ctx echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := ctx.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(ctx)
	}

	if err := ctx.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.ValidationError, err).Send(ctx)
	}

	isExist, err := c.repo.ExistById(payload.ID)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(ctx)
	}
	if !isExist {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(ctx)
	}

	if err := c.repo.DeleteById(payload.ID); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(ctx)
	}

	return res.CustomSuccessBuilder(
		http.StatusOK, nil, "success delete a user",
	).Send(ctx)
}

func (c *controller) GetAll(ctx echo.Context) error {
	users, err := c.repo.FindAll()
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err).Send(ctx)
	}

	return res.CustomSuccessBuilder(
		http.StatusOK, users, "success get all users",
	).Send(ctx)
}
