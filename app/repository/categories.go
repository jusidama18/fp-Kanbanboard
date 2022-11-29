package repository

import (
	"Kanbanboard/app/delivery/params"
	"Kanbanboard/domain"
	"fmt"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (c *CategoryRepository) StoreCategory(req params.CategoryCreate) (*domain.Category, error) {

	cat := domain.Category{
		Type: req.Type,
	}

	err := c.db.Create(&cat).Preload("Tasks").Find(&cat).Error
	if err != nil {
		return nil, err
	}

	return &cat, nil
}

func (c *CategoryRepository) FindAllCategories() ([]domain.Category, error) {

	var categories []domain.Category
	err := c.db.Preload("Tasks").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *CategoryRepository) DeleteCategoryByID(id int) error {
	res := c.db.Delete(&domain.Category{}, id)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected < 1 {
		return fmt.Errorf("Category with id %d not found.", id)
	}

	return nil
}

func (c *CategoryRepository) FindCategoryByID(id int) (*domain.Category, error) {
	var category domain.Category

	res := c.db.Find(&category, id)
	if res.Error != nil {
		return nil, res.Error
	} else if res.RowsAffected < 1 {
		return nil, fmt.Errorf("Category with id %d not found", id)
	}

	return &category, nil
}

func (c *CategoryRepository) UpdateCategoryByID(id int, req params.CategoryUpdate) (*domain.Category, error) {
	currentCategory, err := c.FindCategoryByID(id)
	if err != nil {
		return nil, err
	}

	newCategory := &domain.Category{
		Type: req.Type,
	}

	err = c.db.Model(&currentCategory).Updates(&newCategory).Find(&newCategory).Error
	if err != nil {
		return nil, err
	}

	return newCategory, nil
}
