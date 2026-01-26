package handlers

import (
	"net/http"
	"rest-api/internal/dtos/user"
	userDto "rest-api/internal/dtos/user"
	userUseCase "rest-api/internal/usecases/user"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GinUserHandler struct {
	CreateUserUseCase     *userUseCase.CreateUserUseCase
	GetAllUsersUseCase    *userUseCase.GetAllUsersUseCase
	UpdateUserRoleUseCase *userUseCase.UpdateUserRoleUseCase
	DeleteUserByIdUseCase *userUseCase.DeleteUserByIdUseCase
}

func NewGinUserHandler(
	createUserUseCase *userUseCase.CreateUserUseCase,
	getAllUsersUseCase *userUseCase.GetAllUsersUseCase,
	updateUserRoleUseCase *userUseCase.UpdateUserRoleUseCase,
	deleteUserByIdUseCase *userUseCase.DeleteUserByIdUseCase,
) *GinUserHandler {
	return &GinUserHandler{
		CreateUserUseCase:     createUserUseCase,
		GetAllUsersUseCase:    getAllUsersUseCase,
		UpdateUserRoleUseCase: updateUserRoleUseCase,
		DeleteUserByIdUseCase: deleteUserByIdUseCase,
	}
}

func (h *GinUserHandler) UpdateUserRole(ctx *gin.Context) {
	var request user.UpdateUserRoleRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := govalidator.ValidateStruct(request)
	if !valid || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := ctx.Param("id")
	userId, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
		return
	}

	err = h.UpdateUserRoleUseCase.Execute(request, userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *GinUserHandler) DeleteUserById(ctx *gin.Context) {

	idStr := ctx.Param("id")
	userId, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
		return
	}

	err = h.DeleteUserByIdUseCase.Execute(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *GinUserHandler) CreateUser(ctx *gin.Context) {
	var request user.CreateUserRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := govalidator.ValidateStruct(request)
	if !valid || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.CreateUserUseCase.Execute(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, userDto.UserMinimalResponse{ID: createdUser.ID})
}

func (h *GinUserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.GetAllUsersUseCase.Execute()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
