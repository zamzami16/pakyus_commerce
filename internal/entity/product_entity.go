package entity

import "github.com/google/uuid"

type Product struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
}