package auth

import (
	"fmt"
	authDto "rest-api/internal/dtos/auth"
	"rest-api/internal/infra/repositories"
	"rest-api/internal/infra/services"
)

type LoginUseCase struct {
	UserRepository *repositories.UserRepositoryDb
	JwtService     *services.JwtService
}

func NewLoginUseCase(userRepository *repositories.UserRepositoryDb, jwtService *services.JwtService) *LoginUseCase {
	return &LoginUseCase{
		UserRepository: userRepository,
		JwtService:     jwtService,
	}
}

func (lo *LoginUseCase) Execute(params authDto.AuthUserRequestDto) (authDto.AuthResponseDto, error) {
	user, err := lo.UserRepository.GetByEmail(params.Email)

	if err != nil || user == nil {
		return authDto.AuthResponseDto{}, fmt.Errorf("Error getting user with email: %v", params.Email)
	}

	err = user.Auth(params.Password)

	if err != nil {
		return authDto.AuthResponseDto{}, fmt.Errorf("Error on user %v auth", params.Email)
	}

	jwt, err := lo.JwtService.GenerateCode(*user)

	if err != nil {
		return authDto.AuthResponseDto{}, fmt.Errorf("Error generating jwt for user with email: %v", params.Email)
	}

	return authDto.AuthResponseDto{Token: jwt}, nil
}
