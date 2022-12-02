package usecase

import (
	"Kanbanboard/app/delivery/params"
	"Kanbanboard/domain"
)

type CategoryUsecase struct {
	repo domain.CategoryRepository
}

func NewCategoryUsecase(repo domain.CategoryRepository) domain.CategoryUsecase {
	return &CategoryUsecase{
		repo: repo,
	}
}

func (c *CategoryUsecase) CreateCategory(req params.CategoryCreate) (*domain.CategoryCreateResponse, error) {
	cat, err := c.repo.StoreCategory(req)
	if err != nil {
		return nil, err
	}

	resp := &domain.CategoryCreateResponse{
		ID:        cat.ID,
		Type:      cat.Type,
		CreatedAt: cat.CreatedAt,
	}

	return resp, nil
}

func (c *CategoryUsecase) FindAllCategories() ([]domain.Category, error) {
	resp, err := c.repo.FindAllCategories()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CategoryUsecase) DeleteCategoryByID(id int) error {
	err := c.repo.DeleteCategoryByID(id)
	return err
}

func (c *CategoryUsecase) UpdateCategoryByID(id int, req params.CategoryUpdate) (*domain.CategoryUpdateResponse, error) {
	updatedCategory, err := c.repo.UpdateCategoryByID(id, req)
	if err != nil {
		return nil, err
	}

	resp := &domain.CategoryUpdateResponse{
		ID:        updatedCategory.ID,
		Type:      updatedCategory.Type,
		UpdatedAt: updatedCategory.UpdatedAt,
	}

	return resp, nil
}
