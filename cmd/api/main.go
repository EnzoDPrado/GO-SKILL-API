package main

import (
	"log"
	"os"
	"rest-api/internal/handlers"
	"rest-api/internal/infra/database"
	"rest-api/internal/infra/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var dbC *database.Connection
var jwtSvc *services.JwtService

func init() {
	_ = godotenv.Load()
	connectDatabase()
	loadJwtService()
}

func main() {
	server := gin.Default()

	handlers.RegisterRoutes(server, dbC.Db, jwtSvc)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}

func loadJwtService() {
	jwtSvc = services.NewJwtService(os.Getenv("JWT_SECRET_KEY"))
}

func connectDatabase() {
	dp, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)

	if err != nil {
		log.Fatalf("Error converting DB_PORT environment variable")
	}

	dbC = database.NewConnection(
		os.Getenv("DB_HOST"),
		dp,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	_, err = dbC.Connect()
	if err != nil {
		log.Fatalf("Error connecting on database")
	}
}
