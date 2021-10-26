package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"altastore/config"
	"altastore/models"
	"altastore/util"

	echo "github.com/labstack/echo/v4"
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
	db.Migrator().DropTable(&models.Carts{})
	db.AutoMigrate(&models.Carts{})

	// preparate dummy data
	var newCarts models.Carts
	var newCartDetails models.CartDetails
	newCarts.Products =

	// user dummy data with model
	cartsModel := models.NewCartsModel(db)
	_, err := models.CartDetailsModel().Insert(newCartDetails)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetAllUserController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	cartModel := models.NewCartModel(db)
	userController := NewController(cartModel)

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/users")

	userController.GetAllUserController(context)

	// build struct response
	type Response []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var response Response
	resBody := res.Body.String()

	json.Unmarshal([]byte(resBody), &response)

	t.Run("GET /users", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, 1, len(response))
		assert.Equal(t, "Name Test B", response[0].Name)
		assert.Equal(t, "test@alterra.id", response[0].Email)
	})
}

func TestGetUserController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewUserModel(db)
	userController := NewController(userModel)

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	userController.GetUserController(context)

	// Unmarshal respose string to struct
	type Response struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var response Response
	resBody := res.Body.String()

	json.Unmarshal([]byte(resBody), &response)

	t.Run("GET /users/:id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code) // response.Data.
		assert.Equal(t, "Name Test B", response.Name)
		assert.Equal(t, "test@alterra.id", response.Email)
	})
}

func TestPostUserController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewUserModel(db)
	userController := NewController(userModel)

	// input controller
	reqBody, _ := json.Marshal(map[string]string{
		"name":     "Name Test",
		"email":    "test@alterra.id",
		"password": "test123",
	})

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	context := e.NewContext(req, res)
	context.SetPath("/users")

	userController.PostUserController(context)

	// build struct response
	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	// testing stuff
	t.Run("POST /users", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "Successful Operation", response.Message)
	})
}

func TestEditUserController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewUserModel(db)
	userController := NewController(userModel)

	// input controller
	reqBody, _ := json.Marshal(map[string]string{
		"name":     "Name Test New",
		"email":    "test@alterra.id",
		"password": "test123",
	})

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	context := e.NewContext(req, res)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	userController.EditUserController(context)

	// build struct response
	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	// testing stuff
	t.Run("PUT /users/:id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "Successful Operation", response.Message)
	})
}

func TestDeleteUserController(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewUserModel(db)
	userController := NewController(userModel)

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	context := e.NewContext(req, res)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	userController.DeleteUserController(context)

	// build struct response
	type Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	// testing stuff
	t.Run("PUT /users/:id", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "Successful Operation", response.Message)
	})
}
