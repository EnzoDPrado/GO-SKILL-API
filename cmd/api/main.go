package main

import (
	"log"
	"os"
	"rest-api/internal/handlers"
	"rest-api/internal/infra/database"
	"rest-api/internal/infra/repositories"
	"rest-api/internal/usecases/user"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var dbConnection *database.Connection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dp, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)

	dbConnection = database.NewConnection(
		os.Getenv("DB_HOST"),
		dp,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	_, err = dbConnection.Connect()
	if err != nil {
		log.Fatalf("Error connecting on database")
	}

}

func main() {
	server := gin.Default()
	generateGinUserHandler().RegisterRoutes(server)

	server.Run(":8080")
}

func generateGinUserHandler() *handlers.GinUserHandler {
	userRepository := &repositories.UserRepositoryDb{Db: dbConnection.Db}

	createUserUseCase := user.NewCreateUserUseCase(userRepository)
	getAllUsersUseCase := user.NewGetAllUsersUseCase(userRepository)

	return handlers.NewGinUserHandler(createUserUseCase, getAllUsersUseCase)
}
