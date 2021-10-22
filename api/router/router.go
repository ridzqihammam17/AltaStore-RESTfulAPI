package router

import (
	"altastore/api/controllers"
	"altastore/constants"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, customerController *controllers.Controller) {
	// ------------------------------------------------------------------
	// Login & Register
	// ------------------------------------------------------------------
	e.POST("/api/register", customerController.RegisterCustomerController)
	e.POST("/api/login", customerController.LoginCustomerController)

	// Auth JWT
	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eAuth.GET("/api/customers", customerController.GetAllCustomerController)
}
