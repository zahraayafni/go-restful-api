package models

import (
	"time"

	"github.com/google/uuid"
)

// Technician model
type Technician struct {
	ID              uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	Name            string    `json:"name" db:"name" validate:"required"`
	BrandSpecialist string    `json:"brand_specialist" db:"brand_specialist" validate:"required"`
	Platform        string    `json:"platform" db:"platform" validate:"required"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
