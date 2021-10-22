package products

import (
	"altastore/api/common"
	"altastore/models"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	productModel models.ProductModel
}

func NewController(productModel models.ProductModel) *Controller {
	return &Controller{
		productModel,
	}
}

func (controller *Controller) GetAllProductController(c echo.Context) error {
	product, err := controller.productModel.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, product)
}

func (controller *Controller) GetProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	product, err := controller.productModel.Get(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	response := GetProductResponse{
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) PostProductController(c echo.Context) error {
	// bind request value
	var productRequest PostProductRequest
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	product := models.Product{
		Name:  productRequest.Name,
		Price: productRequest.Price,
		Stock: productRequest.Stock,
	}
	_, err := controller.productModel.Insert(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (controller *Controller) UpdateProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	// bind request value
	var productRequest EditProductRequest
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	product := models.Product{
		Name:  productRequest.Name,
		Price: productRequest.Price,
		Stock: productRequest.Stock,
	}

	if _, err := controller.productModel.Edit(product, id); err != nil {
		return c.JSON(http.StatusNotFound, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (controller *Controller) DeleteProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	if _, err := controller.productModel.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}
