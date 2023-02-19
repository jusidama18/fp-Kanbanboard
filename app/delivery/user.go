package delivery

import (
	"net/http"

	"Kanbanboard/app/delivery/middleware"
	"Kanbanboard/app/delivery/params"
	"Kanbanboard/app/delivery/responses"
	"Kanbanboard/app/helper"
	"Kanbanboard/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(r *gin.Engine, userUsecase domain.UserUsecase) {
	handler := &UserHandler{userUsecase}
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
	userRoute := r.Group("/users")
	userRoute.POST("/register", handler.Register)
	userRoute.POST("/register-admin", handler.RegisterAdmin) // For testing purpose
	userRoute.POST("/login", handler.Login)
	userRoute.Use(middleware.Authentication())
	userRoute.PUT("/update-account", handler.UpdateAccount)
	userRoute.DELETE("/delete-account", handler.DeleteAccount)
}

// @Summary Register New User
// @Description Register New User by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param mygram body params.UserRegister true "Register User"
// @Success 200 {object} responses.Response{data=domain.User}
// @Router /users/register [post]
func (u *UserHandler) Register(ctx *gin.Context) {
	var userRegister params.UserRegister
	err := ctx.ShouldBindJSON(&userRegister)
	if err != nil {
		responses.BadRequestError(ctx, err.Error())
		return
	}
	err = helper.ValidateStruct(userRegister)
	if err != nil {
		responses.BadRequestError(ctx, err.Error())
		return
	}
	var user domain.User
	copier.Copy(&user, &userRegister)
	userData, err := u.userUsecase.Register(ctx.Request.Context(), &user)
	if err != nil {
		responses.Success(ctx, getStatusCode(err), err.Error(), nil)
		return
	}
	responses.Success(ctx, http.StatusCreated, "User Registered.", userData)
}

// @Summary Register New Admin
// @Description Register New Admin by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param mygram body params.UserRegister true "Register Admin"
// @Success 200 {object} responses.Response{data=domain.User}
// @Router /users/register-admin [post]
func (u *UserHandler) RegisterAdmin(ctx *gin.Context) {
	var userRegister params.UserRegister
	err := ctx.ShouldBindJSON(&userRegister)
	if err != nil {
		responses.BadRequestError(ctx, err.Error())
		return
	}
	err = helper.ValidateStruct(userRegister)
	if err != nil {
		responses.BadRequestError(ctx, err.Error())
		return
	}
	var user domain.User
	copier.Copy(&user, &userRegister)
	userData, err := u.userUsecase.RegisterAdmin(ctx.Request.Context(), &user)
	if err != nil {
		responses.Success(ctx, getStatusCode(err), err.Error(), nil)
		return
	}
	responses.Success(ctx, http.StatusCreated, "Admin Created.", userData)
}

// @Summary Login Account
// @Description Login Account by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param mygram body params.UserLogin true "Login Account"
// @Success 200 {object} responses.Response{data=string}
// @Router /users/login [post]
func (u *UserHandler) Login(ctx *gin.Context) {
	var userLogin params.UserLogin
	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		responses.BadRequestError(ctx, err.Error())
		return
	}
	err = helper.ValidateStruct(userLogin)
	if err != nil {
		responses.BadRequestError(ctx, err.Error())
		return
	}
	var user domain.User
	copier.Copy(&user, &userLogin)
	token, err := u.userUsecase.Login(ctx.Request.Context(), &user)
	if err != nil {
		responses.Success(ctx, getStatusCode(err), err.Error(), nil)
		return
	}
	responses.Success(ctx, http.StatusOK, "Login Success.", token)
}

// @Summary Update Account
// @Description Update User by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param mygram body params.UserUpdate true "Update User"
// @Success 200 {object} responses.Response{data=domain.User}
// @Router /users/update-account [put]
func (u *UserHandler) UpdateAccount(ctx *gin.Context) {
	var userUpdate params.UserUpdate
	err := ctx.ShouldBindJSON(&userUpdate)
	if err != nil {
		responses.BadRequestError(ctx, err.Error())
		return
	}
	err = helper.ValidateStruct(userUpdate)
	if err != nil {
		responses.BadRequestError(ctx, err.Error())
		return
	}
	var user domain.User
	copier.Copy(&user, &userUpdate)
	userAuth := ctx.MustGet("user").(jwt.MapClaims)
	userID := int64(userAuth["id"].(float64))
	user.ID = userID
	userData, err := u.userUsecase.UpdateUser(ctx.Request.Context(), &user)
	if err != nil {
		responses.Success(ctx, getStatusCode(err), err.Error(), nil)
		return
	}
	responses.Success(ctx, http.StatusOK, "Account Updated.", userData)
}

// @Summary Delete User
// @Description Delete User through the authentication process must be done with the help of JsonWebToken.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response{data=string}
// @Router /users/delete-account [delete]
func (u *UserHandler) DeleteAccount(ctx *gin.Context) {
	userAuth := ctx.MustGet("user").(jwt.MapClaims)
	userID := int64(userAuth["id"].(float64))
	err := u.userUsecase.DeleteUser(ctx, userID)
	if err != nil {
		responses.Success(ctx, getStatusCode(err), err.Error(), nil)
		return
	}
	responses.Success(ctx, http.StatusOK, "Your account has been successfully deleted.", nil)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
