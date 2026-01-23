package auth

import (
	"fmt"
	authDto "rest-api/internal/dtos/auth"
	"rest-api/internal/infra/repositories"
	"rest-api/internal/infra/services"
)

type LoginUseCase struct {
	UserRepository repositories.UserRepositoryDb
	JwtService     services.JwtService
}

func (lo *LoginUseCase) Execute(params authDto.AuthUserRequestDto) (authDto.AuthResponseDto, error) {
	user, err := lo.UserRepository.GetByEmail(params.Email)

	if err != nil {
		return authDto.AuthResponseDto{}, fmt.Errorf("Unauthorized")
	}

	jwt, err := lo.JwtService.GenerateCode(*user)

	if err != nil {
		return authDto.AuthResponseDto{}, fmt.Errorf("Unauthorized")
	}

	return authDto.AuthResponseDto{Token: jwt}, nil
}
