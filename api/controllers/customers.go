package controllers

import (
	"altastore/models"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

	// Encrypt Password on Register
	hashedPassword, er := bcrypt.GenerateFromPassword([]byte(customerRequest.Password), bcrypt.MinCost)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"code":    500,
			"message": "Internal Server Error",
		})
	}

	customer := models.Customer{
		Name:     customerRequest.Name,
		Address:  customerRequest.Address,
		Gender:   customerRequest.Gender,
		Email:    customerRequest.Email,
		Password: string(hashedPassword),
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
