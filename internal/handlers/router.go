package handlers

import (
	"rest-api/internal/domain"
	"rest-api/internal/handlers/middlewares"
	"rest-api/internal/infra/repositories"
	"rest-api/internal/infra/services"
	"rest-api/internal/usecases/auth"
	"rest-api/internal/usecases/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, jwt *services.JwtService) {
	// Repositories
	userRepo := &repositories.UserRepositoryDb{Db: db}

	// Use Cases
	createUserUC := user.NewCreateUserUseCase(userRepo)
	getAllUsersUC := user.NewGetAllUsersUseCase(userRepo)
	getUserUC := user.NewGetUserByIdUseCase(userRepo)
	updateUserRoleUC := user.NewUpdateUserRoleUseCase(userRepo, getUserUC)
	deleteUserUC := user.NewDeleteUserByIdUseCase(userRepo, getUserUC)
	loginUC := auth.NewLoginUseCase(userRepo, jwt)

	// Handlers
	userHandler := NewGinUserHandler(createUserUC, getAllUsersUC, updateUserRoleUC, deleteUserUC)
	authHandler := NewGinAuth(loginUC)

	api := r.Group("/api/v1")
	{
		api.POST("/login", authHandler.Login)

		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetAllUsers)
			users.PATCH("/:id/role",
				middlewares.AuthMiddleware(jwt),
				middlewares.RoleMiddleware(domain.ADMIN),
				userHandler.UpdateUserRole,
			)
			users.DELETE("/:id",
				middlewares.AuthMiddleware(jwt),
				middlewares.RoleMiddleware(domain.ADMIN),
				userHandler.DeleteUserById,
			)
		}
	}
}
