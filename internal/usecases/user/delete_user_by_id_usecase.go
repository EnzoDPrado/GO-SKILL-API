package user

import (
	"fmt"
	"rest-api/internal/domain"
	"rest-api/internal/infra/repositories"

	"github.com/google/uuid"
)

type DeleteUserByIdUseCase struct {
	UserRepository     repositories.UserRepository
	GetUserByIdUseCase *GetUserByIdUseCase
}

func NewDeleteUserByIdUseCase(repository repositories.UserRepository, getUserByIdUseCase *GetUserByIdUseCase) *DeleteUserByIdUseCase {
	return &DeleteUserByIdUseCase{
		UserRepository:     repository,
		GetUserByIdUseCase: getUserByIdUseCase,
	}
}

func (d *DeleteUserByIdUseCase) Execute(userId uuid.UUID) error {
	user, err := d.GetUserByIdUseCase.Execute(userId)

	if err != nil {
		return err
	}

	if user.Role == domain.ADMIN {
		return fmt.Errorf("Cannot delete a admin user")
	}

	if err := d.UserRepository.DeleteById(userId); err != nil {
		return err
	}

	return nil
}
