package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	// ID       int    `gorm:"primaryKey" json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	//1 to many with products
	// Products Product `gorm:"foreignKey:ProductsID"`
}

type GormCategoryModel struct {
	db *gorm.DB
}

func NewCategoryModel(db *gorm.DB) *GormCategoryModel {
	return &GormCategoryModel{db: db}
}

// Interface Category
type CategoryModel interface {
	GetAll() ([]Category, error)
	Get(categoryId int) (Category, error)
	Add(Category) (Category, error)
	Edit(category Category, categoryId int) (Category, error)
	Delete(categoryId int) (Category, error)
}

func (m *GormCategoryModel) GetAll() ([]Category, error) {
	var category []Category
	if err := m.db.Find(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (m *GormCategoryModel) Get(categoryId int) (Category, error) {
	var category Category
	if err := m.db.Find(&category, categoryId).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (m *GormCategoryModel) Add(category Category) (Category, error) {
	if err := m.db.Save(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (m *GormCategoryModel) Edit(newCategory Category, categoryId int) (Category, error) {
	var category Category
	if err := m.db.Find(&category, "id=?", categoryId).Error; err != nil {
		return category, err
	}

	category.Name = newCategory.Name

	if err := m.db.Save(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (m *GormCategoryModel) Delete(categoryId int) (Category, error) {
	var category Category
	if err := m.db.Find(&category, "id=?", categoryId).Error; err != nil {
		return category, err
	}
	if err := m.db.Delete(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}
