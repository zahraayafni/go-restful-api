package usecase

import (
	"github.com/zahraayafni/go-restful-api/config"
	"github.com/zahraayafni/go-restful-api/internal/customer"
)

// Customer UseCase
type CustomerUC struct {
	cfg               *config.Config
	customerRepo      customer.CustomerRepository
	customerRedisRepo customer.CustomerRedisRepository
}
