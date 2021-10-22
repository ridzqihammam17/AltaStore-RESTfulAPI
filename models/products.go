package models

import (
	"altastore/config"
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ID          int    `gorm:"primaryKey" json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Code        string `json:"code" form:"code"`
	Status      string `json:"status" form:"status"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`

	//many to many with carts
	Carts []*Carts `gorm:"many2many:cart_details"`

	//1 to many with product category
	ProductCategoriesID int `json:"product_categories_id" form:"product_categories_id"`
}


type GormProductsModel struct {
	db *gorm.DB
}

func (m *GormProductsModel) NewCartsModel(db *gorm.DB) *GormCartsModel {
	return &GormCartsModel{db: db}
}
type ProductModel interface {

}

//check is product exist on table product
func (m *GormProductsModel)  CheckProductId(productId int, product []Products) (interface{}, error) {
	var product []Products
	if err := m.db.Where("id=?", productId).First(&product).Error; err != nil {
		return nil, err
	}
	return product.ID, nil
}

// get product by id
func (m *GormProductsModel)  GetProduct(productId int) ([]Products, error) {
	var product []Products
	if err := m.db.Find(&product, "id=?", productId).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (m *GormProductsModel) GetProducts() (interface{}, error) {
	var products []Products
	if err := m.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (m *GormProductsModel) GetProductid(id int) (interface{}, error) {
	var product []Products
	var count int64
	if err1 := m.db.Model(&product).Where("id=?", id).Count(&count).Error; count == 0 {
		return nil, err1
	}
	if err := m.db.Find(&product, "id=?", id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (m *GormProductsModel) CreateProduct(products []Products) (interface{}, error) {
	if err := m.db.Save(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (m *GormProductsModel) GetProductByProductCategory(name string) (interface{}, error) {
	var productcategories []ProductCategories
	if err := m.db.Where("name=?", name).First(&productcategories).Error; err != nil {
		return nil, err
	}
	var products []Products
	if err := m.db.Find(&products, "product_categories_id=?", productcategories.ID).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (m *GormProductsModel) DeleteProductById(id int) (interface{}, error) {
	var products []Products
	if err := m.db.Where("id=?", id).Delete(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

//update user info from database
func (m *GormProductsModel) UpdateProduct(product []Products) (interface{}, error) {
	if tx := m.db.Save(&product).Error; tx != nil {
		return nil, tx
	}
	return product, nil
}

//get 1 specified user with User struct output
func (m *GormProductsModel) GetUpdateProduct(id int) []Products {
	var product []Products
	m.db.Find(&product, "id=?", id)
	return product
}
