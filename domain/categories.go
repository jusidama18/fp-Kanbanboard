package domain

import (
	"Kanbanboard/app/delivery/params"
	"time"
)

type Category struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Type      string    `json:"type" gorm:"notNull"`
	Tasks     []Task    `json:"tasks"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryCreateResponse struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryUpdateResponse struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryRepository interface {
	StoreCategory(params.CategoryCreate) (*Category, error)
	FindAllCategories() ([]Category, error)
	DeleteCategoryByID(id int) error
	FindCategoryByID(id int) (*Category, error)
	UpdateCategoryByID(id int, req params.CategoryUpdate) (*Category, error)
}

type CategoryUsecase interface {
	CreateCategory(params.CategoryCreate) (*CategoryCreateResponse, error)
	FindAllCategories() ([]Category, error)
	DeleteCategoryByID(id int) error
	UpdateCategoryByID(id int, req params.CategoryUpdate) (*CategoryUpdateResponse, error)
}
