package usecase

import "evermos/internal/domain"

type categoryUsecase struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUsecase{categoryRepo}
}

func (uc *categoryUsecase) AddCategory(category *domain.Category) error {
	return uc.categoryRepo.Create(category)
}

func (uc *categoryUsecase) UpdateCategory(category *domain.Category) error {
	return uc.categoryRepo.Update(category)
}

func (uc *categoryUsecase) DeleteCategory(id uint) error {
	return uc.categoryRepo.Delete(id)
}

func (uc *categoryUsecase) GetAllCategories() ([]domain.Category, error) {
	return uc.categoryRepo.GetAll()
}

func (uc *categoryUsecase) GetCategoryByID(id uint) (*domain.Category, error) {
	return uc.categoryRepo.GetByID(id)
}