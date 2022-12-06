package delivery

import (
	"Kanbanboard/app/delivery/middleware"
	"Kanbanboard/app/delivery/params"
	"Kanbanboard/app/delivery/responses"
	"Kanbanboard/app/helper"
	"Kanbanboard/domain"
	"fmt"
	"net/http"
	"strconv"

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
	taskRouter.GET("/", taskCtl.GetAllTasks)
	taskRouter.PUT("/:id", taskCtl.PutTask)
	taskRouter.PATCH("/update-status/:id", taskCtl.PatchTaskStatus)
	taskRouter.PATCH("/update-category/:id", taskCtl.PatchTaskCategory)
	taskRouter.DELETE("/:id", taskCtl.DeleteTask)
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

func (t *TaskController) GetAllTasks(c *gin.Context) {
	resp, err := t.uc.GetAllTasks()
	if err != nil {
		responses.InternalServerError(c, err.Error())
	}

	responses.Success(c, http.StatusOK, "get all tasks successfully", resp)
}

func getUserID(c *gin.Context) (int, error) {
	id, exists := c.Get("userID")
	if !exists {
		return 0, fmt.Errorf("user id not found")
	}
	userID := int(id.(float64))
	return userID, nil
}

func (t *TaskController) PutTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}
	var req params.TaskPutByID
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

	resp, err := t.uc.PutTask(id, req)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}
	responses.Success(c, http.StatusOK, "task update successfully", resp)
}

func (t *TaskController) PatchTaskStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}
	var req params.TaskUpdateStatus
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

	resp, err := t.uc.PatchTaskStatus(id, req)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}
	responses.Success(c, http.StatusOK, "task update-status patched successfully", resp)
}

func (t *TaskController) PatchTaskCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}
	var req params.TaskUpdateCategory
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

	resp, err := t.uc.PatchTaskCategory(id, req)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}
	responses.Success(c, http.StatusOK, "task update-category patched successfully", resp)
}

func (t *TaskController) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	err2 := t.uc.DeleteTask(id)
	if err2 != nil {
		responses.InternalServerError(c, err.Error())
		return
	}
	responses.Success(c, http.StatusOK, "task deleted successfully", nil)
}
