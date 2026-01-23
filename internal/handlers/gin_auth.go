package handlers

import (
	"fmt"
	"net/http"
	authDto "rest-api/internal/dtos/auth"
	JwtService "rest-api/internal/infra/services"
	authUseCase "rest-api/internal/usecases/auth"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type GinAuth struct {
	LoginUseCase *authUseCase.LoginUseCase
	JwtService   *JwtService.JwtService
}

func NewGinAuth(
	loginUseCase *authUseCase.LoginUseCase,
) *GinAuth {
	return &GinAuth{
		LoginUseCase: loginUseCase,
	}
}

func (g *GinAuth) Login(ctx *gin.Context) {
	var request authDto.AuthUserRequestDto

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := govalidator.ValidateStruct(request)
	if !valid || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := g.LoginUseCase.Execute(request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("Bad Credentials").Error()})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
