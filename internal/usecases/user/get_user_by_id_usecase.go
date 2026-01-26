package user

import (
	"fmt"
	"rest-api/internal/domain"
	"rest-api/internal/infra/repositories"

	"github.com/google/uuid"
)

type GetUserByIdUseCase struct {
	UserRepository repositories.UserRepository
}

func NewGetUserByIdUseCase(repository repositories.UserRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		UserRepository: repository,
	}
}

func (g *GetUserByIdUseCase) Execute(id uuid.UUID) (*domain.User, error) {
	user, err := g.UserRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	if !user.Status {
		return nil, fmt.Errorf("User not founded")
	}

	return user, nil
}
