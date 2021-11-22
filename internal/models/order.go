package models

import (
	"time"

	"github.com/google/uuid"
)

// Order model
type Order struct {
	ID             uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	CustomerID     uuid.UUID `json:"customer_id" db:"customer_id" validate:"omitempty,uuid"`
	ServicesID     uuid.UUID `json:"services_id" db:"services_id" validate:"omitempty,uuid"`
	TechnicianID   uuid.UUID `json:"technician_id" db:"technician_id" validate:"omitempty,uuid"`
	Brand          string    `json:"brand" db:"brand" validate:"required"`
	TechnicianName string    `json:"technician_name" db:"technician_name" validate:"required"`
	DamageType     string    `json:"damage_type" db:"damage_type" validate:"required"`
	Description    string    `json:"description" db:"description" validate:"required"`
	Status         int       `json:"status" db:"status" validate:"required"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
