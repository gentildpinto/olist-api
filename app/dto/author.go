package dto

type (
	Author struct {
		Name string `json:"name" validate:"required"`
	}
)
