package controllers

import (
	"net/http"

	"altastore/models"

	echo "github.com/labstack/echo/v4"
)
type CartController struct {
	cartModel models.Carts
}

func NewCartController(cartModel models.Carts) *CartController {
	return &CartController{
		cartModel,
	}
}

func (controller *CartController) CreateCartController(c echo.Context) error {

	// ------------ cart -------------//
	//rec user input
	var cart models.Carts
	c.Bind(&cart) //input: payment method id

	// get id user login
	customerId := middlewares.ExtractTokenUserId(c)

	//check product id on table product
	paymentId := cart.PaymentMethodsID
	var payment models.PaymentMethods
	checkPayment, err := controller.cartModel.CheckPayment(paymentId, payment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        "Cant find payment method",
			"checkProductId": checkPayment,
		})
	}

	//set data cart and create new cart
	cart = models.Carts{
		StatusTransactions: "ordered",
		TotalQuantity:      0,
		TotalPrice:         0,
		CustomersID:        customerId,
		PaymentMethodsID:   paymentId,
	}
	newCart, _ := controller.cartModel.CreateCart(cart)

	//------------ cart detail -------------//
	// convert product id
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Product id is invalid",
		})
	}

	// check product id on table product
	var product models.Products
	checkProductId, err := controller.cartModel.CheckProductId(productId, product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        "Product isn't found with product id #" + strconv.Itoa(productId),
			"checkProductId": checkProductId,
		})
	}

	//get price
	getProduct, _ := controller.cartModel.GetProduct(productId)

	//convert qty
	qty, _ := strconv.Atoi(c.Param("qty"))

	//set data cart details
	cartDetails := models.CartDetails{
		ProductsID: productId,
		CartsID:    newCart.ID,
		Quantity:   qty,
		Price:      getProduct.Price,
	}

	//create cart detail
	newCartDetail, _ := controller.cartModel.AddToCart(cartDetails)

	//update total quantity and total price on table carts
	UpdateTotalCart(newCart.ID)

	//get cart updated (total qty&total price)
	updatedCart, _ := controller.cartModel.GetCart(newCart.ID)

	//custom data cart for body response
	outputCart := map[string]interface{}{
		"ID":                  updatedCart.ID,
		"customers_id":        updatedCart.CustomersID,
		"payment_methods_id":  updatedCart.PaymentMethodsID,
		"status_transactions": updatedCart.StatusTransactions,
		"total_quantity":      updatedCart.TotalQuantity,
		"total_price":         updatedCart.TotalPrice,
		"CreatedAt":           updatedCart.CreatedAt,
		"UpdatedAt":           updatedCart.UpdatedAt,
		"DeletedAt":           updatedCart.DeletedAt,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"cart":        outputCart,
		"cartDetails": newCartDetail,
		"status":      "Create cart success",
	})
}

//func for update total quantity and total price on table carts
func (controller *CartController) UpdateTotalCart(cartId int) (int, int) {
	newTotalPrice, _ := controller.cartModel.GetTotalPrice(cartId)
	newTotalQty, _ := database.GetTotalQty(cartId)
	newCart, _ := database.UpdateTotalCart(cartId, newTotalPrice, newTotalQty)

	return newCart.TotalQuantity, newCart.TotalPrice
}

func (controller *CartController) GetCartController(c echo.Context) error {
	//convert cart_id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id cart",
		})
	}
	//is cart id exist
	var cart models.Carts
	checkCartId, err := controller.cartModel.CheckCartId(id, cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":        "Cant find cart id",
			"checkProductId": checkCartId,
		})
	}

	listCart, _ := controller.cartModel.GetCartById(id)        //get cart by id
	products, _ := controller.cartModel.GetListProductCart(id) //get all products based on cart id

	//custom data cart for body response
	outputCart := map[string]interface{}{
		"ID":                  listCart.ID,
		"customers_id":        listCart.CustomersID,
		"payment_methods_id":  listCart.PaymentMethodsID,
		"status_transactions": listCart.StatusTransactions,
		"total_quantity":      listCart.TotalQuantity,
		"total_price":         listCart.TotalPrice,
		"CreatedAt":           listCart.CreatedAt,
		"UpdatedAt":           listCart.UpdatedAt,
		"DeletedAt":           listCart.DeletedAt,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"cart":     outputCart,
		"products": products,
		"status":   "Success get all products by cart id",
	})
}

func (controller *CartController) DeleteCartController(c echo.Context) error {
	//convert cart id
	cartId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid cart id",
		})
	}

	//check is cart id exist on table cart
	var cart models.Carts
	checkCartId, err := controller.cartModel.CheckCartId(cartId, cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Can't find cart",
			"checkCartId": checkCartId,
		})
	}

	//delete cart and all products included on it
	deletedCart, _ := controller.cartModel.DeleteCart(cartId)

	//custom output data cart for body response
	outputCart := map[string]interface{}{
		"ID":                  deletedCart.ID,
		"customers_id":        deletedCart.CustomersID,
		"payment_methods_id":  deletedCart.PaymentMethodsID,
		"status_transactions": deletedCart.StatusTransactions,
		"total_quantity":      deletedCart.TotalQuantity,
		"total_price":         deletedCart.TotalPrice,
		"CreatedAt":           deletedCart.CreatedAt,
		"UpdatedAt":           deletedCart.UpdatedAt,
		"DeletedAt":           deletedCart.DeletedAt,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":       "Delete cart success",
		"Deleted Cart": outputCart,
	})
}
