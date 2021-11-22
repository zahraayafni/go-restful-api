package middleware

import (
	"github.com/zahraayafni/go-restful-api/config"
)

// Middleware manager
type MiddlewareManager struct {
	cfg     *config.Config
	origins []string
}

// Middleware manager constructor
func NewMiddlewareManager(cfg *config.Config, origins []string) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, origins: origins}
}
