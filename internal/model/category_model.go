package model

import "github.com/google/uuid"

type CategoryResponse struct {
	ID       uuid.UUID  `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	ParentId *uuid.UUID `json:"parent_id,omitempty"`
}
