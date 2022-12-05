package delivery

import (
	"Kanbanboard/app/delivery/middleware"
	"Kanbanboard/app/delivery/params"
	"Kanbanboard/app/delivery/responses"
	"Kanbanboard/app/helper"
	"Kanbanboard/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	uc domain.TaskUseCase
}

func NewTaskController(r *gin.Engine, uc domain.TaskUseCase) {
	taskCtl := TaskController{
		uc: uc,
	}

	taskRouter := r.Group("/tasks")
	taskRouter.Use(middleware.Authorization([]string{"member", "admin"}))
	taskRouter.POST("/", taskCtl.CreateTask)
}

func (t *TaskController) CreateTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		responses.UnauthorizedRequest(c, err.Error())
		return
	}

	var req params.TaskCreate
	err = c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	err = helper.ValidateStruct(req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	resp, err := t.uc.CreateTask(req, userID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "task created successfully", resp)
}

func getUserID(c *gin.Context) (int, error) {
	id, exists := c.Get("userID")
	if !exists {
		return 0, fmt.Errorf("user id not found")
	}
	userID := int(id.(float64))
	return userID, nil
}
