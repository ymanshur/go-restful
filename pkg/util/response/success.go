package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type successConstant struct {
	OK Success
}

var SuccessConstant successConstant = successConstant{
	OK: Success{
		Response: successResponse{
			Meta: Meta{
				Success: true,
				Message: "request successfully proceed",
			},
			Data: nil,
		},
		Code: http.StatusOK,
	},
}

type successResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Success struct {
	Response successResponse `json:"response"`
	Code     int             `json:"code"`
}

func SuccessBuilder(res *Success, data interface{}) *Success {
	res.Response.Data = data
	return res
}

func CustomSuccessBuilder(code int, data interface{}, message string) *Success {
	return &Success{
		Response: successResponse{
			Meta: Meta{
				Success: true,
				Message: message,
			},
			Data: data,
		},
		Code: code,
	}
}

func SuccessResponse(data interface{}) *Success {
	return SuccessBuilder(&SuccessConstant.OK, data)
}

func (s *Success) Send(ctx echo.Context) error {
	return ctx.JSON(s.Code, s.Response)
}
