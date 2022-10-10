package dto

type (
	ByIDRequest struct {
		ID uint `param:"id" validate:"required"`
	}
)
