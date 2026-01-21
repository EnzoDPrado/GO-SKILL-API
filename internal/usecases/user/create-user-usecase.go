package user

import (
	"rest-api/internal/domain"
	"rest-api/internal/dtos/user"
	"rest-api/internal/infra/repositories"

	"github.com/google/uuid"
)

type CreateUserUseCase struct {
	UserRepository repositories.UserRepository
}

func NewCreateUserUseCase(userRepository repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (self *CreateUserUseCase) Execute(createUserRequest user.CreateUserRequest) (*domain.User, error) {

	userEntity := domain.User{
		ID:   uuid.New(),
		Name: createUserRequest.Name,
	}

	user, err := self.UserRepository.Add(userEntity)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
