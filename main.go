package main

import (
	"fmt"
	"os"

	_handler "Kanbanboard/app/delivery"
	_repository "Kanbanboard/app/repository"
	_usecase "Kanbanboard/app/usecase"

	"Kanbanboard/config"

	docs "Kanbanboard/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title KanBanBoard-API
// @version 1.0
// @description This is a API webservice to manage KanBanBoard API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email hacktiv@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	router := gin.Default()
	config.StartDB()
	db := config.GetDBConnection()

	base_url := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if base_url != "" {
		docs.SwaggerInfo.Host = base_url
	} else {
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", port)
	}

	userRepository := _repository.NewUserRepository(db)
	userUsecase := _usecase.NewUserUsecase(userRepository)

	catRepository := _repository.NewCategoryRepository(db)
	catUseCase := _usecase.NewCategoryUsecase(catRepository)

	taskRepository := _repository.NewTaskRepository(db)
	taskUseCase := _usecase.NewTaskService(taskRepository)

	_handler.NewUserHandler(router, userUsecase)
	_handler.NewCategoryHandler(router, catUseCase)
	_handler.NewTaskController(router, taskUseCase)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":" + port)
}
