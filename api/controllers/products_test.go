package controllers

import (
	"altastore/config"
	"altastore/models"
	"altastore/util"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	// create database connection
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)

	// cleaning data before testing
	db.Migrator().DropTable(&models.Product{})
	db.AutoMigrate(&models.Product{})

	// preparate dummy data
	var newProduct models.Product
	newProduct.Name = "Product A"
	newProduct.Price = "10000"
	newProduct.Stock = "100"

	// user dummy data with model
	customerModel := models.NewProductModel(db)
	_, err := customerModel.Insert(newProduct)
	if err != nil {
		fmt.Println(err)
	}

}

func TestGetAllProductController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	productModel := models.NewProductModel(db)
	productController := NewProductController(productModel)

	// reqBody, _ := json.Marshal(models.Product{Email: "test_a@alterra.id", Password: "password123"})

	// // setting controller
	e := echo.New()
	// loginReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	// loginRes := httptest.NewRecorder()
	// loginReq.Header.Set("Content-Type", "application/json")
	// loginContext := e.NewContext(loginReq, loginRes)
	// loginContext.SetPath("/customers/login")

	// if err := LoginProductController(loginContext); err != nil {
	// 	t.Errorf("should not get error, get error: %s", err)
	// 	return
	// }

	// var c models.Product
	// resBody := loginRes.Body.String()
	// json.Unmarshal([]byte(resBody), &c)

	// // testing stuff
	// t.Run("POST /customers/login", func(t *testing.T) {
	// 	assert.Equal(t, 200, loginRes.Code)
	// 	assert.NotEqual(t, "", c.Token)
	// })

	// token := c.Token

	// setting controller
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/api/products")

	productController.GetAllProductController(context)
	// var m *models.GormProductModel
	var productList []models.Product
	json.Unmarshal(res.Body.Bytes(), &productList)

	t.Run("GET /api/products", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, 1, len(productList))
		assert.Equal(t, "Product A", productList[0].Name)
		assert.Equal(t, "10000", productList[0].Price)
		assert.Equal(t, "100", productList[0].Stock)
	})
}

func TestGetProductController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	productModel := models.NewProductModel(db)
	productController := NewProductController(productModel)

	// setting controller
	e := echo.New()
	// reqBody, _ := json.Marshal(LoginProductRequest{Email: "test_a@alterra.id", Password: "password123"})
	// loginReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	// loginRes := httptest.NewRecorder()
	// loginReq.Header.Set("Content-Type", "application/json")
	// loginContext := e.NewContext(loginReq, loginRes)
	// loginContext.SetPath("/products/login")

	// if err := productController.LoginProductController(loginContext); err != nil {
	// 	t.Errorf("should not get error, get error: %s", err)
	// 	return
	// }

	// var c models.Product
	// json.Unmarshal(loginRes.Body.Bytes(), &c)

	// // testing stuff
	// t.Run("POST /products/login", func(t *testing.T) {
	// 	assert.Equal(t, 200, loginRes.Code)
	// 	assert.NotEqual(t, "", c.Token)
	// })

	// token := c.Token

	// setting controller
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/api/products/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	// id, _ := strconv.Atoi(context.Param("id"))
	productController.GetProductController(context)

	var productList []models.Product

	json.Unmarshal(res.Body.Bytes(), &productList)

	t.Run("GET /api/products/:id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code) // response.Data.
		// assert.Equal(t, 1, int(c.ID))
		// assert.Equal(t, "Name Test A", c.Name)
		// assert.Equal(t, "test_a@alterra.id", c.Email)
	})
}

func TestPostProductController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(userModel)
	productModel := models.NewProductModel(db)
	productController := NewProductController(productModel)

	// setting controller
	e := echo.New()
	reqBody, _ := json.Marshal(models.Customer{Email: "ilham@gmail.com@alterra.id", Password: "pass123"})
	loginReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	loginRes := httptest.NewRecorder()
	loginReq.Header.Set("Content-Type", "application/json")
	loginContext := e.NewContext(loginReq, loginRes)
	loginContext.SetPath("/api/login")

	if err := customerController.LoginCustomerController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	var c models.Customer
	json.Unmarshal(loginRes.Body.Bytes(), &c)

	// testing stuff
	t.Run("POST /products/login", func(t *testing.T) {
		assert.Equal(t, 200, loginRes.Code)
		assert.NotEqual(t, "", c.Token)
	})

	token := c.Token

	// input controller
	reqBodyPost, _ := json.Marshal(map[string]string{
		"name":  "Product B",
		"price": "5000",
		"stock": "5",
	})

	// setting controller
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBodyPost))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products")

	productController.PostProductController(context)

	// build struct response
	// type Response struct {
	// 	Code    int    `json:"code"`
	// 	Message string `json:"message"`
	// }
	// var response Response
	// resBody := res.Body.String()
	// json.Unmarshal([]byte(resBody), &response)
	var p models.Product
	json.Unmarshal(loginRes.Body.Bytes(), &c)

	// testing stuff
	// t.Run("POST /products/login", func(t *testing.T) {
	// 	assert.Equal(t, 200, loginRes.Code)
	// 	assert.NotEqual(t, "", c.Token)
	// })

	// testing stuff
	t.Run("POST /api/products", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		// assert.Equal(t, "Successful Operation", response.Message)
		assert.Equal(t, "Product B", p.Name)
		assert.Equal(t, "5000", p.Price)
		assert.Equal(t, "5", p.Stock)
	})
}

func TestUpdateProductController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(userModel)
	productModel := models.NewProductModel(db)
	productController := NewProductController(productModel)

	// setting controller
	e := echo.New()
	reqBody, _ := json.Marshal(models.Customer{Email: "ilham@gmail.com@alterra.id", Password: "pass123"})
	loginReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	loginRes := httptest.NewRecorder()
	loginReq.Header.Set("Content-Type", "application/json")
	loginContext := e.NewContext(loginReq, loginRes)
	loginContext.SetPath("/api/login")

	if err := customerController.LoginCustomerController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	var c models.Customer
	json.Unmarshal(loginRes.Body.Bytes(), &c)

	// testing stuff
	t.Run("POST /products/login", func(t *testing.T) {
		assert.Equal(t, 200, loginRes.Code)
		assert.NotEqual(t, "", c.Token)
	})

	token := c.Token

	// input controller
	reqBodyPost, _ := json.Marshal(map[string]string{
		"name":  "Product B",
		"price": "5000",
		"stock": "5",
	})

	// setting controller
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBodyPost))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products")

	context.SetPath("/products/:id")
	context.SetParamNames("id")
	context.SetParamValues("2")

	productController.UpdateProductController(context)

	var p models.Product
	json.Unmarshal(loginRes.Body.Bytes(), &p)

	// testing stuff
	t.Run("PUT /api/products", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		// assert.Equal(t, "Successful Operation", response.Message)
		assert.Equal(t, "Product B", p.Name)
		assert.Equal(t, "5000", p.Price)
		assert.Equal(t, "5", p.Stock)
	})

	// // build struct response
	// type Response struct {
	// 	Code    int    `json:"code"`
	// 	Message string `json:"message"`
	// }
	// var response Response
	// resBody := res.Body.String()
	// json.Unmarshal([]byte(resBody), &response)

	// // testing stuff
	// t.Run("PUT /products/:id", func(t *testing.T) {
	// 	assert.Equal(t, 200, res.Code)
	// 	assert.Equal(t, "Successful Operation", response.Message)
	// })
}

func TestDeleteProductController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewCustomerModel(db)
	customerController := NewCustomerController(userModel)
	productModel := models.NewProductModel(db)
	productController := NewProductController(productModel)

	// setting controller
	e := echo.New()
	reqBody, _ := json.Marshal(models.Customer{Email: "ilham@gmail.com@alterra.id", Password: "pass123"})
	loginReq := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	loginRes := httptest.NewRecorder()
	loginReq.Header.Set("Content-Type", "application/json")
	loginContext := e.NewContext(loginReq, loginRes)
	loginContext.SetPath("/api/login")

	if err := customerController.LoginCustomerController(loginContext); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}

	var c models.Customer
	json.Unmarshal(loginRes.Body.Bytes(), &c)

	// testing stuff
	t.Run("POST /products/login", func(t *testing.T) {
		assert.Equal(t, 200, loginRes.Code)
		assert.NotEqual(t, "", c.Token)
	})

	token := c.Token

	// input controller
	reqBodyPost, _ := json.Marshal(map[string]string{
		"name":  "Product B",
		"price": "5000",
		"stock": "5",
	})

	// setting controller
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBodyPost))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/products")

	context.SetPath("/products/:id")
	context.SetParamNames("id")
	context.SetParamValues("2")

	productController.UpdateProductController(context)

	var p models.Product
	json.Unmarshal(loginRes.Body.Bytes(), &p)

	// testing stuff
	t.Run("PUT /api/products", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		// assert.Equal(t, "Successful Operation", response.Message)
		assert.Equal(t, "Product B", p.Name)
		assert.Equal(t, "5000", p.Price)
		assert.Equal(t, "5", p.Stock)
	})

	context.SetPath("/products/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	productController.DeleteProductController(context)

	// build struct response
	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	// testing stuff
	t.Run("PUT /products/:id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		// assert.Equal(t, "Successful Operation", response.Message)
	})
}
