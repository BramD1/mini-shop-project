package repository

import (
    "mini-shop/domain"
    "gorm.io/gorm"
)

type categoryRepository struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
    return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAll() ([]domain.Category, error) {
    var categories []domain.Category
    err := r.db.Find(&categories).Error
    return categories, err
}

func (r *categoryRepository) FindByID(id uint) (domain.Category, error) {
    var category domain.Category
    err := r.db.First(&category, id).Error
    return category, err
}

func (r *categoryRepository) Create(category domain.Category) (domain.Category, error) {
    err := r.db.Create(&category).Error
    return category, err
}

func (r *categoryRepository) Update(category domain.Category) (domain.Category, error) {
    err := r.db.Save(&category).Error
    return category, err
}

func (r *categoryRepository) Delete(id uint) error {
    result := r.db.Delete(&domain.Category{}, id)
    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
    return result.Error
}