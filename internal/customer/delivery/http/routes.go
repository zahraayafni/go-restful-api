package http

import (
	"github.com/labstack/echo/v4"
	"github.com/zahraayafni/go-restful-api/internal/customer"
	"github.com/zahraayafni/go-restful-api/internal/middleware"
)

// Map customer routes
func MapCustomerRoutes(customerGroup *echo.Group, h customer.CustomerHandlers, mw *middleware.MiddlewareManager) {
	customerGroup.POST("/customer/insert", h.InsertCustomerHandler())
	customerGroup.POST("/customer/update", h.UpdateCustomerHandler())
	customerGroup.POST("/customer/delete", h.DeleteCustomerByIDsHandler())
	customerGroup.POST("/customer/get", h.GetCustomerByIDsHandler())
}
