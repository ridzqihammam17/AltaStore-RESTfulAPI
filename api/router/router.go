package router

import (
	"altastore/api/controllers"
	"altastore/constants"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, customerController *controllers.Controller) {
	e.POST("/customers/", customerController.PostCustomerController)
	e.POST("/customers", customerController.PostCustomerController)
	// e.POST("/customers/login/", customerController.LoginCustomerController)
	// e.POST("/customers/login", customerController.LoginCustomerController)

	// Auth JWT
	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// Get All User
	eAuth.GET("customers/", customerController.GetAllCustomerController)
	eAuth.GET("customers", customerController.GetAllCustomerController)
}
