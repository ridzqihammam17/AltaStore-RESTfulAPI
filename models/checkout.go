package models

import (
	"time"

	"gorm.io/gorm"
)

type Checkout struct {
	ID        int `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time
}

type Checkout_Response struct {
	ID      int
	Product []Checkout_Response
}

type Checkout_Input struct {
	Courier   string `json:"courier" form:"courier"`
	ProductID []int  `json:"product_id" form:"product_id"`
}

type GormCheckoutModel struct {
	db *gorm.DB
}

func NewCheckoutModel(db *gorm.DB) *GormProductModel {
	return &GormProductModel{db: db}
}

func (m *GormProductModel) AddCheckoutID() (int, error) {
	var checkout Checkout
	if err := m.db.Save(&checkout).Error; err != nil {
		return 0, err
	}
	return checkout.ID, nil
}

func (m *GormProductModel) UpdateCheckoutIdInCartItem(checkoutID, cartID, productID int) (CartDetails, error) {
	var cartItem CartDetails

	if err := m.db.Where("cart_id = ? and product_id = ? and checkout_id IS NULL", cartID, productID).First(&cartItem).Error; err != nil {
		return cartItem, err
	}

	var CheckoutID = checkoutID

	if err := m.db.Model(&cartItem).Update("checkout_id", CheckoutID).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}
