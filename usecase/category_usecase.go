package usecase

import (
    "mini-shop/domain"
)

type categoryUsecase struct {
    categoryRepo domain.CategoryRepository
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepository) domain.CategoryUsecase {
    return &categoryUsecase{categoryRepo: categoryRepo}
}

func (u *categoryUsecase) GetAllCategories() ([]domain.Category, error) {
    return u.categoryRepo.FindAll()
}

func (u *categoryUsecase) GetCategoryByID(id uint) (domain.Category, error) {
    return u.categoryRepo.FindByID(id)
}

func (u *categoryUsecase) CreateCategory(category domain.Category) (domain.Category, error) {
    return u.categoryRepo.Create(category)
}

func (u *categoryUsecase) UpdateCategory(category domain.Category) (domain.Category, error) {
    return u.categoryRepo.Update(category)
}

func (u *categoryUsecase) DeleteCategory(id uint) error {
    return u.categoryRepo.Delete(id)
}