package main

import (
	"os"

	_handler "Kanbanboard/app/delivery"
	_repository "Kanbanboard/app/repository"
	_usecase "Kanbanboard/app/usecase"

	"Kanbanboard/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.StartDB()
	db := config.GetDBConnection()

	userRepository := _repository.NewUserRepository(db)
	userUsecase := _usecase.NewUserUsecase(userRepository)

	api := router.Group("/")
	_handler.NewUserHandler(api, userUsecase)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
