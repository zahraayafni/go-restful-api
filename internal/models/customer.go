package models

import (
	"time"

	"github.com/google/uuid"
)

// Customer model
type Customer struct {
	ID        uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Email     string    `json:"email" db:"email" validate:"required"`
	Msisdn    string    `json:"msisdn" db:"msisdn" validate:"required"`
	Address   string    `json:"address" db:"addres" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
