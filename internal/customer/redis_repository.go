//go:generate mockgen -source redis_repository.go -destination mock/redis_repository_mock.go -package mock
package customer

import (
	"context"

	"github.com/google/uuid"
	"github.com/zahraayafni/go-restful-api/internal/models"
)

// Customer redis repository
type CustomerRedisRepository interface {
	GetCustomerByIDsCache(ctx context.Context, customerIDs []uuid.UUID) (map[uuid.UUID]models.Customer, error)
	SetCustomersCache(ctx context.Context, customers []models.Customer, ttl int) error
	ClearCustomersCache(ctx context.Context, customerIDs []uuid.UUID) error
}
