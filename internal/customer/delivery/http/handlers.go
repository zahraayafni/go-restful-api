package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/zahraayafni/go-restful-api/config"
	"github.com/zahraayafni/go-restful-api/internal/customer"
	"github.com/zahraayafni/go-restful-api/pkg/httpErrors"
)

// Customer handlers
type customerHandlers struct {
	cfg        *config.Config
	customerUC customer.CustomerUseCase
}

// Init customer handlers constructor
func InitCustomerHandlers(cfg *config.Config, customerUC customer.CustomerUseCase) customer.CustomerHandlers {
	return &customerHandlers{cfg: cfg, customerUC: customerUC}
}
