package domain

import (
	"rest-api/internal/dtos/user"

	"github.com/google/uuid"
)

type CreateUserProvider interface {
	Execute(createUserRequest user.CreateUserRequest) (*User, error)
}

type DeleteUserByIdProvider interface {
	Execute(userId uuid.UUID) error
}

type GetAllUsersProvider interface {
	Execute() ([]user.UserSimplifiedResponse, error)
}

type GetUserByIdProvider interface {
	Execute(id uuid.UUID) (*User, error)
}

type UpdateUserRoleProvider interface {
	Execute(input user.UpdateUserRoleRequest, userId uuid.UUID) error
}
