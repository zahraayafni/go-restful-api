//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package customer

import (
	"context"

	"github.com/google/uuid"

	"github.com/zahraayafni/go-restful-api/internal/models"
)

// Customer Repository
type CustomerRepository interface {
	InsertCustomerDB(ctx context.Context, customer *models.Customer) (uuid.UUID, error)
	UpdateCustomerDB(ctx context.Context, customer *models.Customer) error
	GetCustomerByIDsDB(ctx context.Context, customerIDs []uuid.UUID) ([]models.Customer, error)
	DeleteCustomerByIDDB(ctx context.Context, customerID uuid.UUID) error
}
