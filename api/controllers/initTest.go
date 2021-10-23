package controllers

import (
	"altastore/config"
	"altastore/models"
	"altastore/util"
	"fmt"
)

func setup() {
	// create database connection
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)

	// cleaning data before testing
	db.Migrator().DropTable(&models.Customer{})
	db.AutoMigrate(&models.Customer{})
	db.Migrator().DropTable(&models.Category{})
	db.AutoMigrate(&models.Category{})
	db.Migrator().DropTable(&models.Product{})
	db.AutoMigrate(&models.Product{})
	// preparate dummy data
	var newCustomer models.Customer
	newCustomer.Name = "Ilham"
	newCustomer.Email = "ilham@gmail.com"
	newCustomer.Password = "pass123"

	var newCategory models.Category
	newCategory.Name = "Category A"

	var newProduct models.Product
	newProduct.Name = "Product A"
	newProduct.Price = 10000
	newProduct.Stock = 100
	newProduct.CategoryID = 1
	// user dummy data with model
	customerModel := models.NewCustomerModel(db)
	_, err := customerModel.Register(newCustomer)
	if err != nil {
		fmt.Println(err)
	}
	productModel := models.NewProductModel(db)
	_, err = productModel.Insert(newProduct)
	if err != nil {
		fmt.Println(err)
	}
	categoryModel := models.NewCategoryModel(db)
	_, err = categoryModel.Add(newCategory)
	if err != nil {
		fmt.Println(err)
	}
}
