package router

import (
	"altastore/api/controllers"
	"altastore/constants"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo,
	customerController *controllers.CustomerController,
	categoryController *controllers.CategoryController,
) {
	// ------------------------------------------------------------------
	// Login & Register
	// ------------------------------------------------------------------
	e.POST("/api/register", customerController.RegisterCustomerController)
	e.POST("/api/login", customerController.LoginCustomerController)

	// Auth JWT
	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// Customers
	eAuth.GET("/api/customers", customerController.GetAllCustomerController)

	// ------------------------------------------------------------------
	// CRUD Categories
	// ------------------------------------------------------------------
	eAuth.GET("/api/categories", categoryController.GetAllCategoryController)
	eAuth.GET("/api/categories/:id", categoryController.GetCategoryController)
	eAuth.POST("/api/categories", categoryController.AddCategoryController)
	eAuth.PUT("/api/categories/:id", categoryController.EditCategoryController)
	eAuth.DELETE("/api/categories/:id", categoryController.DeleteCategoryController)

}
