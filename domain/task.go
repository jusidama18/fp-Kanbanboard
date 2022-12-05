package domain

import (
	"Kanbanboard/app/delivery/params"
	"time"
)

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `gorm:"notNull" json:"title"`
	Description string    `gorm:"notNull" json:"description"`
	Status      bool      `gorm:"notNull" json:"status"`
	UserID      int       `json:"user_id"`
	User        User      `json:"-"`
	CategoryID  int       `json:"category_id"`
	Category    Category  `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type TaskRepository interface {
	CreateTask(params.TaskCreate, int) (*Task, error)
}

type CreateTaskResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskUseCase interface {
	CreateTask(params.TaskCreate, int) (*CreateTaskResponse, error)
}
