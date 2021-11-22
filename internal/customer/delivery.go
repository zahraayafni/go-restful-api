package customer

import "github.com/labstack/echo/v4"

// Customer HTTP Handlers interface
type CustomerHandlers interface {
	InsertCustomerHandler() echo.HandlerFunc
	UpdateCustomerHandler() echo.HandlerFunc
	GetCustomerByIDsHandler() echo.HandlerFunc
	DeleteCustomerByIDsHandler() echo.HandlerFunc
}
