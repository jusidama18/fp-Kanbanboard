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
	catRoute.POST("/", middleware.AuthorizeAdmin(), catHandler.CreateCategory)
	catRoute.GET("/", catHandler.FindAllCategories)
	catRoute.DELETE("/:id", middleware.AuthorizeAdmin(), catHandler.DeleteCategoryByID)
	catRoute.PATCH("/:id", catHandler.UpdateCategoryByID)
}

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

func (cat *CategoryHandler) FindAllCategories(c *gin.Context) {
	categories, err := cat.usecase.FindAllCategories()
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "successfully get all data", categories)
}

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

	responses.Success(c, http.StatusOK, "Category has been successfully deleted")
}

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
