package user

import (
	"fmt"
	"rest-api/internal/domain"
	userDto "rest-api/internal/dtos/user"
	"rest-api/internal/infra/repositories"

	"github.com/google/uuid"
)

type UpdateUserRoleUseCase struct {
	UserRepository     repositories.UserRepository
	GetUserByIdUseCase domain.GetUserByIdProvider
}

func NewUpdateUserRoleUseCase(repository repositories.UserRepository, getUserByIdUseCase domain.GetUserByIdProvider) *UpdateUserRoleUseCase {
	return &UpdateUserRoleUseCase{
		UserRepository:     repository,
		GetUserByIdUseCase: getUserByIdUseCase,
	}
}

func (u *UpdateUserRoleUseCase) Execute(input userDto.UpdateUserRoleRequest, userId uuid.UUID) error {
	if _, err := u.GetUserByIdUseCase.Execute(userId); err != nil {
		return err
	}

	if err := u.validateInput(input); err != nil {
		return err
	}

	if err := u.UserRepository.UpdateUserRole(userId, input.Role); err != nil {
		return err
	}

	return nil
}

func (u *UpdateUserRoleUseCase) validateInput(input userDto.UpdateUserRoleRequest) error {
	inputRole := input.Role

	if _, err := domain.CastUserRole(inputRole); err != nil {
		return fmt.Errorf("Role %v is a invalid role", inputRole)
	}

	return nil
}
