package usecase

import (
	"Kanbanboard/app/delivery/params"
	"Kanbanboard/domain"
)

type taskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) domain.TaskUseCase {
	return &taskService{
		repo: repo,
	}
}

func (t *taskService) CreateTask(req params.TaskCreate, userID int) (*domain.CreateTaskResponse, error) {

	task, err := t.repo.CreateTask(req, userID)
	if err != nil {
		return nil, err
	}

	respTask := &domain.CreateTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		UserID:      task.UserID,
		CategoryID:  task.CategoryID,
		CreatedAt:   task.CreatedAt,
	}

	return respTask, nil
}
