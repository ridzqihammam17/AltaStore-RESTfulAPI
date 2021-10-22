package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Price string `json:"price" form:"price"`
	Stock string `json:"stock" form:"stock"`

	//1 to many with category
	// Categories []Category `gorm:"foreignKey:CategoryID"`
}

type GormProductModel struct {
	db *gorm.DB
}

func NewProductModel(db *gorm.DB) *GormProductModel {
	return &GormProductModel{db: db}
}

// Interface Product
type ProductModel interface {
	GetAll() ([]Product, error)
	Get(productId int) (Product, error)
	Insert(Product) (Product, error)
	Edit(product Product, productId int) (Product, error)
	Delete(productId int) (Product, error)
}

func (m *GormProductModel) GetAll() ([]Product, error) {
	var product []Product
	if err := m.db.Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (m *GormProductModel) Get(productId int) (Product, error) {
	var product Product
	if err := m.db.Find(&product, productId).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (m *GormProductModel) Insert(product Product) (Product, error) {
	if err := m.db.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (m *GormProductModel) Edit(newProduct Product, productId int) (Product, error) {
	var product Product
	if err := m.db.Find(&product, productId).Error; err != nil {
		return product, err
	}

	product.Name = newProduct.Name
	product.Price = newProduct.Price
	product.Stock = newProduct.Stock

	if err := m.db.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (m *GormProductModel) Delete(productId int) (Product, error) {
	var product Product
	if err := m.db.Find(&product, productId).Error; err != nil {
		return product, err
	}

	if err := m.db.Delete(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
