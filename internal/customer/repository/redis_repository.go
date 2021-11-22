package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/zahraayafni/go-restful-api/internal/customer"
	"github.com/zahraayafni/go-restful-api/internal/models"
)

// Customer redis repository
type customerRedisRepository struct {
	redisClient *redis.Client
}

// Customer redis repository constructor
func InitCustomerRedisRepository(redisClient *redis.Client) customer.CustomerRedisRepository {
	return &customerRedisRepository{redisClient: redisClient}
}

// Get customer by ids
func (r *customerRedisRepository) GetCustomerByIDsCache(ctx context.Context, customerIDs []uuid.UUID) (map[uuid.UUID]models.Customer, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "customerRedisRepository.GetCustomerByIDsCache")
	defer span.Finish()

	keys := make([]string, 0, len(customerIDs))
	for _, id := range customerIDs {
		keys = append(keys, getCustomerCacheKey(id))
	}

	customersRaw, err := r.redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, errors.Wrap(err, "customerRedisRepository.GetCustomerByIDsCache.redisClient.MGet")
	}
	customers := make(map[uuid.UUID]models.Customer, len(customersRaw))
	for _, c := range customersRaw {
		customer := c.(models.Customer)
		customers[customer.ID] = customer
	}

	return customers, nil
}

// Cache customers data
func (r *customerRedisRepository) SetCustomersCache(ctx context.Context, customers []models.Customer, ttl int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "customerRedisRepository.SetCustomersCache")
	defer span.Finish()

	cacheData := constructCustomersCacheData(customers)
	if err := r.redisClient.MSetNX(ctx, cacheData, time.Second*time.Duration(ttl)).Err(); err != nil {
		return errors.Wrap(err, "customerRedisRepository.SetCustomersCache.redisClient.MSetNX")
	}
	return nil
}

// Clear customer cache by ids
func (r *customerRedisRepository) ClearCustomersCache(ctx context.Context, customerIDs []uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "customerRedisRepository.ClearCustomersCache")
	defer span.Finish()

	keys := make([]string, 0, len(customerIDs))
	for _, id := range customerIDs {
		keys = append(keys, getCustomerCacheKey(id))
	}

	_, err := r.redisClient.Del(ctx, keys...).Result()
	if err != nil {
		return errors.Wrap(err, "customerRedisRepository.ClearCustomersCache.redisClient.Del")
	}

	return nil
}

func constructCustomersCacheData(customers []models.Customer) map[string]interface{} {
	customersCacheData := make(map[string]interface{}, len(customers))
	for _, c := range customers {
		customersCacheData[getCustomerCacheKey(c.ID)] = c
	}
	return customersCacheData
}

func getCustomerCacheKey(customerID uuid.UUID) string {
	return fmt.Sprintf(CustomerByIDRedisKey, customerID.String())
}
