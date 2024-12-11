package crud

import "github.com/google/uuid"

type UpdateFieldBookRequest struct {
	ID    uuid.UUID `json:"id" validate:"required,notempty"`
	Value string    `json:"value" validate:"required,notempty"`
}
