package controllers

import (
	"net/http"

	"altastore/models"

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

func (controller *Controller) GetAllCustomerController(c echo.Context) error {
	customer, err := controller.customerModel.GetAll()
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Request")
	}

	return c.JSON(http.StatusOK, customer)
}

func (controller *Controller) PostCustomerController(c echo.Context) error {
	// bind request value
	var customerRequest models.Customer
	if err := c.Bind(&customerRequest); err != nil {
		return c.String(http.StatusBadRequest, "Invalid Request")
	}

	customer := models.Customer{
		Name:     customerRequest.Name,
		Address:  customerRequest.Address,
		Gender:   customerRequest.Gender,
		Email:    customerRequest.Email,
		Password: customerRequest.Password,
	}
	_, err := controller.customerModel.Insert(customer)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Error")
	}

	return c.String(http.StatusOK, "Success Add Account")
}