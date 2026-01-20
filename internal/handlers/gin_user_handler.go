package handlers

import (
	"net/http"
	"rest-api/internal/dtos/user"
	userDto "rest-api/internal/dtos/user"
	userUseCase "rest-api/internal/usecases/user"

	"github.com/gin-gonic/gin"
)

type GinUserHandler struct {
	CreateUserUseCase  *userUseCase.CreateUserUseCase
	GetAllUsersUseCase *userUseCase.GetAllUsersUseCase
}

func NewGinUserHandler(
	createUserUseCase *userUseCase.CreateUserUseCase,
	getAllUsersUseCase *userUseCase.GetAllUsersUseCase,
) *GinUserHandler {
	return &GinUserHandler{
		CreateUserUseCase:  createUserUseCase,
		GetAllUsersUseCase: getAllUsersUseCase,
	}
}

func (h *GinUserHandler) CreateUser(ctx *gin.Context) {
	var request user.CreateUserRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	createdUser, err := h.CreateUserUseCase.Execute(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, userDto.UserMinimalResponse{ID: createdUser.ID})
}

func (h *GinUserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.GetAllUsersUseCase.Execute()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *GinUserHandler) RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/users")
	{
		v1.POST("/", h.CreateUser)
		v1.GET("/", h.GetAllUsers)
	}
}
