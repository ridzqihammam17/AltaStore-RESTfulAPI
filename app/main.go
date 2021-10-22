package main

import (
	"fmt"

	"altastore/api/controllers"
	"altastore/api/router"
	"altastore/config"
	"altastore/models"
	"altastore/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDatabaseConnection(config)

	//initiate model
	customerModel := models.NewCustomerModel(db)
	categoryModel := models.NewCategoryModel(db)

	//initiate controller
	newCustomerController := controllers.NewCustomerController(customerModel)
	newCategoryController := controllers.NewCategoryController(categoryModel)

	//create echo http
	e := echo.New()

	//register API path and controller
	router.Route(e, newCustomerController, newCategoryController)

	// run server
	address := fmt.Sprintf("localhost:%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
