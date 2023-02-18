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

// @Summary Create Task
// @Description Create Task by Data Provided
// @Tags Tasks
// @Accept json
// @Produce json
// @Param data body params.TaskCreate true "Create Task"
// @Success 200 {object} responses.Response{data=domain.CreateTaskResponse}
// @Router /tasks [post]
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


// @Summary Get All Task
// @Description Get All Task
// @Tags Tasks
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response{data=[]domain.GetAllTasksResponse}
// @Router /tasks [get]
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


// @Summary Put Task
// @Description Put Task by Data Provided
// @Tags Tasks
// @Accept json
// @Produce json
// @Param data body params.TaskPutByID true "Put Task"
// @Param id path int true "Task ID"
// @Success 200 {object} responses.Response{data=domain.TaskResponse}
// @Router /tasks/{id} [put]
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


// @Summary Patch Task Status
// @Description Patch Task Status by Data Provided
// @Tags Tasks
// @Accept json
// @Produce json
// @Param data body params.TaskUpdateStatus true "Patch Task Status"
// @Param id path int true "Task ID"
// @Success 200 {object} responses.Response{data=domain.TaskResponse}
// @Router /tasks/update-status/{id} [patch]
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

// @Summary Patch Task Category
// @Description Patch Task Category by Data Provided
// @Tags Tasks
// @Accept json
// @Produce json
// @Param data body params.TaskUpdateCategory true "Patch Task Category"
// @Param id path int true "Task ID"
// @Success 200 {object} responses.Response{data=domain.TaskResponse}
// @Router /tasks/update-category/{id} [patch]
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

// @Summary Delete Task
// @Description Delete Task by Data Provided
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path int true "Delete Task"
// @Success 200 {object} responses.Response{data=string}
// @Router /tasks/{id} [delete]
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
