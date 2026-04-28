package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model           // ← ID, CreatedAt, UpdatedAt, DeletedAt all included here
	NamaCategory         string `json:"nama_category"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryRepository interface {
	FindAll() ([]Category, error)
	FindByID(id uint) (Category, error)
	Create(category Category) (Category, error)
	Update(category Category) (Category, error)
	Delete(id uint) error
}

type CategoryUsecase interface {
	GetAllCategories() ([]Category, error)
	GetCategoryByID(id uint) (Category, error)
	CreateCategory(category Category) (Category, error)
	UpdateCategory(category Category) (Category, error)
	DeleteCategory(id uint) error
}