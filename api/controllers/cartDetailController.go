package controllers

import (
	"net/http"
	"strconv"
	"time"

	"altastore/models"

	echo "github.com/labstack/echo/v4"
)

type CartDetailsController struct {
	cartModel models.CartModel
	cartDetailsModel models.CartDetailsModel
	productModel models.ProductModel
}

func NewCartDetailController(cartModel models.CartModel, cartDetailsModel models.CartDetailsModel, productModel models.ProductModel) *CartDetailsController {
	return &CartDetailsController{
		cartModel,
		cartDetailsModel,
		productModel,
	}
}

type CartDetails struct {
	ProductsID int `json:"products_id" form:"products_id"`
	CartsID    int `json:"carts_id" form:"carts_id"`
	Quantity   int `json:"quantity" form:"quantity"`
	Price      int `json:"price" form:"price"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
func (controller *CartDetailsController) AddToCartController(c echo.Context) error {


	//check id cart is exist
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Cart Id is Invalid",
		})
	}
	checkCartId, err := controller.cartModel.CheckCartId(cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Can't find cart",
			"checkCartId": checkCartId,
		})
	}

	// record user's input
	var cartDetails models.CartDetails
	c.Bind(&cartDetails) //entry key: product id, qty

	//check product id on table product
	productId := cartDetails.ProductsID //get product_id
	checkProductId, err := controller.productModel.CheckProductId(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        "Can't find product",
			"checkProductId": checkProductId,
		})
	}

	//get price
	getProduct, _ := controller.productModel.Get(productId)
	productPrice, _ := strconv.Atoi(getProduct.Price)
	//set data cart details
	cartDetails = CartDetails{
		ProductsID: productId,
		CartsID:    cartId,
		Quantity:   cartDetails.Quantity,
		Price:      productPrice,
	}

	//create cart detail
	newCartDetail, _ := controller.cartDetailsModel.AddToCart(cartDetails)
	//update total quantity and total price on table carts
	newTotalQty, newTotalPrice := controller.cartModel.UpdateTotalCart(cartId, productPrice, 1)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"cartDetails":    newCartDetail,
		"Total Quantity": newTotalQty,
		"Total Price":    newTotalPrice,
		"status":         "Successfully added product to cart",
	})
}

func (controller *CartDetailsController) DeleteProductFromCartController(c echo.Context) error {
	//convert cart id
	cartId, err := strconv.Atoi(c.Param("carts_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Cart id is invalid",
		})
	}

	//check is cart id exist on table cart
	checkCartId, err := controller.cartModel.CheckCartId(cartId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Cart isn't found",
			"checkCartId": checkCartId,
		})
	}

	//convert product id
	productId, err := strconv.Atoi(c.Param("products_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Product id is invalid",
		})
	}

	//check is product id exist on table product
	checkProductId, err := controller.productModel.CheckProductId(productId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Product isn't found",
			"checkCartId": checkProductId,
		})
	}

	//check is product id and cart id exist on table cart_detail
	var cartDetails []CartDetails
	checkProductAndCartId, err := controller.cartDetailsModel.CheckProductAndCartId(productId, cartId, cartDetails)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Cant find product id and cart id",
			"checkCartId": checkProductAndCartId,
		})
	}

	//---------delete product------//
	countProduct, _ := controller.cartDetailsModel.CountProductOnCart(cartId) //count product
	var deleteProduct interface{}
	newTotalQty, newTotalPrice, _ := controller.cartDetailsModel.CountProductandPriceOnCart(cartId)

	if countProduct > 1 { //if product on cart > 1, delete product on cart detail + update total on cart
		deleteProduct, _ = controller.cartDetailsModel.DeleteProductFromCart(cartId, productId)
		controller.cartModel.UpdateTotalCart(cartId, newTotalPrice, countProduct-1)
	} else if countProduct == 1 { //if product only 1, delete product on cart detail + delete cart + output total = 0
		deleteProduct, _ = controller.cartDetailsModel.DeleteProductFromCart(cartId, productId)
		controller.cartModel.DeleteCart(cartId)
		newTotalPrice = 0
		newTotalQty = 0
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Deleted Product": deleteProduct,
		"Total Quantity":  newTotalQty,
		"Total Price":     newTotalPrice,
		"status":          "Successfully deleted product on table cart_details",
	})
}
