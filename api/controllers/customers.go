package controllers

import (
	"altastore/models"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	customerModel models.CustomerModel
}

func NewController(customerModel models.CustomerModel) *Controller {
	return &Controller{
		customerModel,
	}
}

func (controller *Controller) RegisterCustomerController(c echo.Context) error {
	var customerRequest models.Customer
	if err := c.Bind(&customerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	customer := models.Customer{
		Name:     customerRequest.Name,
		Address:  customerRequest.Address,
		Gender:   customerRequest.Gender,
		Email:    customerRequest.Email,
		Password: customerRequest.Password,
	}
	_, err := controller.customerModel.Register(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"code":    500,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Register",
	})
}

func (controller *Controller) LoginCustomerController(c echo.Context) error {
	var customerRequest models.Customer

	if err := c.Bind(&customerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	customer, err := controller.customerModel.Login(customerRequest.Email, customerRequest.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Login",
		"token":   customer.Token,
	})
}

func (controller *Controller) GetAllCustomerController(c echo.Context) error {
	customer, err := controller.customerModel.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Get All Customer",
		"data":    customer,
	})
}
