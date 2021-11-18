package payloads

type (
	CreateAuthorPayload struct {
		Name string `json:"name" validate:"required"`
	}
)
