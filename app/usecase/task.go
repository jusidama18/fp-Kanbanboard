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

func (t *taskService) GetAllTasks() ([]domain.GetAllTasksResponse, error) {
	tasks, err := t.repo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	var respTasks []domain.GetAllTasksResponse
	respTasks = parseGetAllTasks(tasks)

	return respTasks, nil
}

func parseGetAllTasks(tasks []domain.Task) []domain.GetAllTasksResponse {
	var respTasks []domain.GetAllTasksResponse
	for i := 0; i < len(tasks); i++ {
		respTask := domain.GetAllTasksResponse{}

		respTask.ID = tasks[i].ID
		respTask.Title = tasks[i].Title
		respTask.Description = tasks[i].Description
		respTask.Status = tasks[i].Status
		respTask.UserID = tasks[i].UserID
		respTask.CreatedAt = tasks[i].CreatedAt
		respTask.CategoryID = tasks[i].CategoryID
		respTask.User.ID = int(tasks[i].User.ID)
		respTask.User.Email = tasks[i].User.Email
		respTask.User.FullName = tasks[i].User.FullName
		respTasks = append(respTasks, respTask)
	}

	return respTasks
}
