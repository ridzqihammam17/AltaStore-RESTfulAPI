package router

import (
	"altastore/api/controllers/customers"
	"altastore/api/controllers/products"
	"altastore/constants"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, customerController *customers.Controller, productController *products.Controller) {
	// ------------------------------------------------------------------
	// Login & Register
	// ------------------------------------------------------------------
	e.POST("/api/register", customerController.RegisterCustomerController)
	e.POST("/api/login", customerController.LoginCustomerController)

	// Auth JWT
	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eAuth.GET("/api/customers", customerController.GetAllCustomerController)

	// Product
	e.GET("/api/product/", productController.GetAllProductController)
	e.GET("/api/product", productController.GetAllProductController)
	eAuth.POST("/api/product/", productController.PostProductController)
	eAuth.POST("/api/product", productController.PostProductController)
	e.GET("/api/product/:id/", productController.GetProductController)
	e.GET("/api/product/:id", productController.GetProductController)
	eAuth.PUT("/api/product/:id/", productController.UpdateProductController)
	eAuth.PUT("/api/product/:id", productController.UpdateProductController)
	eAuth.DELETE("/api/product/:id/", productController.DeleteProductController)
	eAuth.DELETE("/api/product/:id", productController.DeleteProductController)
	// e.GET("/api/product/category/", productController.GetProductCategoryController)
	// e.GET("/api/product/category", productController.GetProductCategoryController)

	// Category
	// eAuth.GET("/api/category/", categoryController.GetAllCategoryController)
	// eAuth.GET("/api/category", categoryController.GetAllCategoryController)
	// eAuth.POST("/api/category/", categoryController.PostCategoryController)
	// eAuth.POST("/api/category", categoryController.PostCategoryController)
	// eAuth.GET("/api/category/:id/", categoryController.GetCategoryController)
	// eAuth.GET("/api/category/:id", categoryController.GetCategoryController)
	// eAuth.PUT("/api/category/:id/", categoryController.UpdateCategoryController)
	// eAuth.PUT("/api/category/:id", categoryController.UpdateCategoryController)
	// eAuth.DELETE("/api/category/:id/", categoryController.DeleteCategoryController)
	// eAuth.DELETE("/api/category/:id", categoryController.DeleteCategoryController)

	// Checkout
	// eAuth.POST("/api/checkout/", checkoutController.PostCheckoutController)
	// eAuth.POST("/api/checkout", checkoutController.PostCheckoutController)

	// Transaction
	// eAuth.GET("/api/transaction/", transactionController.GetAllTransactionController)
	// eAuth.GET("/api/transaction", transactionController.GetAllTransactionController)
	// eAuth.GET("/api/transaction/:id", transactionController.GetAllTransactionController)
	// eAuth.GET("/api/transaction/:id", transactionController.GetAllTransactionController)

	// Cart
	// eAuth.GET("/api/cart/:id/", cartController.GetAllCategoryController)
	// eAuth.GET("/api/cart", cartController.GetAllCategoryController)
	// eAuth.POST("/api/cart/", cartController.PostCategoryController)
	// eAuth.POST("/api/cart", cartController.PostCategoryController)
	// eAuth.GET("/api/cart/:id/", cartController.GetCategoryController)
	// eAuth.GET("/api/cart/:id", cartController.GetCategoryController)
	// eAuth.PUT("/api/cart/:id/", cartController.UpdateCategoryController)
	// eAuth.PUT("/api/cart/:id", cartController.UpdateCategoryController)
	// eAuth.DELETE("/api/cart/:id/", cartController.DeleteCategoryController)
	// eAuth.DELETE("/api/cart/:id", cartController.DeleteCategoryController)
}
