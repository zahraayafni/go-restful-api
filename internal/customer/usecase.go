package customer

import (
	"context"

	"github.com/google/uuid"

	"github.com/zahraayafni/go-restful-api/internal/models"
)

// Customer use case
type CustomerUseCase interface {
	InsertCustomer(ctx context.Context, customer *models.Customer) (uuid.UUID, error)
	UpdateCustomer(ctx context.Context, customer *models.Customer) error
	GetCustomerByIDs(ctx context.Context, customerIDs []uuid.UUID) ([]models.Customer, error)
	DeleteCustomer(ctx context.Context, customerID uuid.UUID) error
}
