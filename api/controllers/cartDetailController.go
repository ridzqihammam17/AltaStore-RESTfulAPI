package controllers

import (
	"net/http"
	"strconv"

	"altastore/models"

	echo "github.com/labstack/echo/v4"
)

type CartDetailsController struct {
	cartDetailModel models.CartDetailsModel
}

func NewCartDetailController(cartDetailsModel models.CartDetails) *CartDetailsController {
	return &CartDetailsController{
		cartDetailsModel,
	}
}
func (controller *CartDetailsController) AddToCartController(c echo.Context) error {
	var cart models.Carts

	//check id cart is exist
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Cart Id is Invalid",
		})
	}
	checkCartId, err := controller.cartDetailModel.CheckCartId(cartId, cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Can't find cart",
			"checkCartId": checkCartId,
		})
	}

	// record user's input
	var cartDetails models.CartDetails
	c.Bind(&cartDetails) //entry key: productid, qty

	//check product id on table product
	productId := cartDetails.ProductsID //get product_id
	var product models.Products
	checkProductId, err := controller.cartDetailModel
	.CheckProductId(productId, product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        "Can't find product",
			"checkProductId": checkProductId,
		})
	}

	//get price
	getProduct, _ := controller.cartDetailModel.GetProduct(productId)

	//set data cart details
	cartDetails = models.CartDetails{
		ProductsID: productId,
		CartsID:    cartId,
		Quantity:   cartDetails.Quantity,
		Price:      getProduct.Price,
	}

	//create cart detail
	newCartDetail, _ := controller.cartDetailModel.AddToCart(cartDetails)

	//update total quantity and total price on table carts
	newTotalQty, newTotalPrice := controller.cartDetailModel.UpdateTotalCart(cartId)

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
	var cart models.Carts
	checkCartId, err := controller.cartModel.CheckCartId(cartId, cart)
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
	var product models.Products
	checkProductId, err := controller.productModel.CheckProductId(productId, product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Product isn't found",
			"checkCartId": checkProductId,
		})
	}

	//check is product id and cart id exist on table cart_detail
	var cartDetails models.CartDetails
	checkProductAndCartId, err := controller.cartDetailModel.CheckProductAndCartId(productId, cartId, cartDetails)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Cant find product id and cart id",
			"checkCartId": checkProductAndCartId,
		})
	}

	//---------delete product------//
	countProduct, _ := controller.cartDetailModel.CountProductOnCart(cartId) //count product
	var deleteProduct interface{}
	var newTotalQty, newTotalPrice int

	if countProduct > 1 { //if product on cart > 1, delete product on cart detail + update total on cart
		deleteProduct, _ = controller.cartDetailModel.DeleteProductFromCart(cartId, productId)
		newTotalQty, newTotalPrice = controller.cartModel.UpdateTotalCart(cartId)
	} else if countProduct == 1 { //if product only 1, delete product on cart detail + delete cart + output total = 0
		deleteProduct, _ = controller.cartDetailModel.DeleteProductFromCart(cartId, productId)
		controller.cartDetailModel.DeleteCart(cartId)
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
