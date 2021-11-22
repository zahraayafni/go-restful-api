package models

import (
	"time"

	"github.com/google/uuid"
)

// Services model
type Services struct {
	ID          uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	DamageType  string    `json:"damage_type" db:"damage_type" validate:"required"`
	FixDuration int       `json:"fix_duration" db:"fix_duration" validate:"required"`
	Fee         int       `json:"fee" db:"fee" validate:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
