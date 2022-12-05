package repository

import (
	"Kanbanboard/app/delivery/params"
	"Kanbanboard/domain"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (t *taskRepository) CreateTask(req params.TaskCreate, userID int) (*domain.Task, error) {

	task := domain.Task{
		Title:       req.Title,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		UserID:      userID,
	}

	err := t.db.Create(&task).Find(&task).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.ForeignKeyViolation:
				err = fmt.Errorf("Category with id %d not found!", req.CategoryID)
			}
		}

		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetAllTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	err := t.db.Preload("User").Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
