package models

import (
	// "altastore/middlewares"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	// ID       int    `gorm:"primaryKey" json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`

	//1 to many with carts
	// Carts []Carts `gorm:"foreignKey:CustomersID"`
}

type GormCustomerModel struct {
	db *gorm.DB
}

func NewCustomerModel(db *gorm.DB) *GormCustomerModel {
	return &GormCustomerModel{db: db}
}

// Interface Customer
type CustomerModel interface {
	// Get(customerId int) (Customer, error)
	Register(Customer) (Customer, error)
	GetAll() ([]Customer, error)
	// Edit(csutomer Customer, customerId int) (Customer, error)
	// Delete(customerId int) (Customer, error)
	// Login(email, password string) (Customer, error)
}

func (m *GormCustomerModel) Register(customer Customer) (Customer, error) {
	if err := m.db.Save(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (m *GormCustomerModel) GetAll() ([]Customer, error) {
	var customer []Customer
	if err := m.db.Find(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}
