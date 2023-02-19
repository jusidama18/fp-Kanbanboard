package delivery

import (
	"Kanbanboard/app/delivery/middleware"
	"Kanbanboard/app/delivery/params"
	"Kanbanboard/app/delivery/responses"
	"Kanbanboard/app/helper"
	"Kanbanboard/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	usecase domain.CategoryUsecase
}

func NewCategoryHandler(r *gin.Engine, cat domain.CategoryUsecase) {
	catHandler := CategoryHandler{
		usecase: cat,
	}
	catRoute := r.Group("/categories")
	catRoute.POST("/", middleware.Authorization([]string{"admin"}), catHandler.CreateCategory)
	catRoute.GET("/", catHandler.FindAllCategories)
	catRoute.DELETE("/:id", middleware.Authorization([]string{"admin"}), catHandler.DeleteCategoryByID)
	catRoute.PATCH("/:id", catHandler.UpdateCategoryByID)
}

// @Summary Create Category
// @Description Create Category by Data Provided
// @Tags Categories
// @Accept json
// @Produce json
// @Param data body params.CategoryCreate true "Create Category"
// @Success 200 {object} responses.Response{data=domain.Category}
// @Router /categories [post]
func (cat *CategoryHandler) CreateCategory(c *gin.Context) {
	var req params.CategoryCreate

	err := c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	err = helper.ValidateStruct(req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	respCategory, err := cat.usecase.CreateCategory(req)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusCreated, "category successfully created", respCategory)
}

// @Summary Get All Category
// @Description Get All Category
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response{data=[]domain.Category}
// @Router /categories [get]
func (cat *CategoryHandler) FindAllCategories(c *gin.Context) {
	categories, err := cat.usecase.FindAllCategories()
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "successfully get all data", categories)
}

// @Summary Delete Category
// @Description Delete Category by Data Provided
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Delete Category"
// @Success 200 {object} responses.Response{data=string}
// @Router /categories/{id} [delete]
func (cat *CategoryHandler) DeleteCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	err = cat.usecase.DeleteCategoryByID(id)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "Category has been successfully deleted", nil)
}

// @Summary Patch Category
// @Description Patch Category by Data Provided
// @Tags Categories
// @Accept json
// @Produce json
// @Param data body params.CategoryUpdate true "Patch Task Category"
// @Param id path int true "Category ID"
// @Success 200 {object} responses.Response{data=domain.CategoryUpdateResponse}
// @Router /categories/{id} [patch]
func (cat *CategoryHandler) UpdateCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	var req params.CategoryUpdate
	err = c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	err = helper.ValidateStruct(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	resp, err := cat.usecase.UpdateCategoryByID(id, req)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "category successfully updated", resp)
}
