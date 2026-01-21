package main

import (
	"rest-api/internal/handlers"
	"rest-api/internal/infra/repositories"
	"rest-api/internal/usecases/user"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	generateGinUserHandler().RegisterRoutes(server)

	server.Run(":8080")
}

func generateGinUserHandler() *handlers.GinUserHandler {
	userRepository := &repositories.UserRepositoryDb{}

	createUserUseCase := user.NewCreateUserUseCase(userRepository)
	getAllUsersUseCase := user.NewGetAllUsersUseCase(userRepository)

	return handlers.NewGinUserHandler(createUserUseCase, getAllUsersUseCase)
}
