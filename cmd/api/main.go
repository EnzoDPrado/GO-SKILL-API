package main

import (
	"log"
	"os"
	"rest-api/internal/handlers"
	"rest-api/internal/infra/database"
	"rest-api/internal/infra/repositories"
	"rest-api/internal/infra/services"
	"rest-api/internal/usecases/auth"
	"rest-api/internal/usecases/user"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var dbConnection *database.Connection
var jwtService *services.JwtService

func init() {
	loadEnv()
	connectDatabase()
	loadJwtService()
}

func main() {
	server := gin.Default()
	generateGinUserHandler().RegisterUserRoutes(server)
	generateGinAuthHandler().RegisterAuthRoutes(server)

	server.Run(":8080")
}

func loadJwtService() {
	jwtService = services.NewJwtService(os.Getenv("JWT_SECRET_KEY"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}

func connectDatabase() {
	dp, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)

	if err != nil {
		log.Fatalf("Error converting DB_PORT environment variable")
	}

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

func generateGinUserHandler() *handlers.GinUserHandler {
	userRepository := &repositories.UserRepositoryDb{Db: dbConnection.Db}

	createUserUseCase := user.NewCreateUserUseCase(userRepository)
	getAllUsersUseCase := user.NewGetAllUsersUseCase(userRepository)

	return handlers.NewGinUserHandler(createUserUseCase, getAllUsersUseCase)
}

func generateGinAuthHandler() *handlers.GinAuth {
	userRepository := repositories.UserRepositoryDb{Db: dbConnection.Db}

	loginUseCase := auth.NewLoginUseCase(userRepository, jwtService)

	return handlers.NewGinAuth(*loginUseCase)
}
