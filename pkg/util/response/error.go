package response

import (
	"fmt"

	"net/http"

	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Meta  Meta   `json:"meta"`
	Error string `json:"error"`
}

type Error struct {
	Response     errorResponse `json:"response"`
	Code         int           `json:"code"`
	ErrorMessage error
}

const (
	E_DUPLICATE            = "duplicate"
	E_NOT_FOUND            = "not_found"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
	E_UNAUTHORIZED         = "unauthorized"
	E_BAD_REQUEST          = "bad_request"
	E_SERVER_ERROR         = "server_error"
)

type errorConstant struct {
	Duplicate                Error
	EmailOrPasswordIncorrect Error
	NotFound                 Error
	RouteNotFound            Error
	UnprocessableEntity      Error
	Unauthorized             Error
	BadRequest               Error
	ValidationError          Error
	InternalServerError      Error
}

var ErrorConstant errorConstant = errorConstant{
	Duplicate: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "created value already exists",
			},
			Error: E_DUPLICATE,
		},
		Code: http.StatusConflict,
	},
	EmailOrPasswordIncorrect: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "email or password is incorrect",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	NotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "data not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	RouteNotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "route not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	UnprocessableEntity: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "invalid parameters or payload",
			},
			Error: E_UNPROCESSABLE_ENTITY,
		},
		Code: http.StatusUnprocessableEntity,
	},
	Unauthorized: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "unauthorized, please login or use different role",
			},
			Error: E_UNAUTHORIZED,
		},
		Code: http.StatusUnauthorized,
	},
	BadRequest: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "bad Request",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	ValidationError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "invalid parameters or payload",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	InternalServerError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Something bad happened",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
}

func ErrorBuilder(res *Error, message error) *Error {
	res.ErrorMessage = message
	return res
}

func CustomErrorBuilder(code int, err string, message string) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: message,
			},
			Error: err,
		},
		Code: code,
	}
}

func ErrorResponse(err error) *Error {
	if res, ok := err.(*Error); ok {
		return res
	}

	return ErrorBuilder(&ErrorConstant.InternalServerError, err)
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d", e.Code)
}

func (e *Error) ParseToError() error {
	return e
}

func (e *Error) Send(ctx echo.Context) error {
	// var errorMessage string
	// if e.ErrorMessage != nil {
	// 	errorMessage = fmt.Sprintf("%+v", errors.WithStack(e.ErrorMessage))
	// }
	// logrus.Error(errorMessage)

	return ctx.JSON(e.Code, e.Response)
}
