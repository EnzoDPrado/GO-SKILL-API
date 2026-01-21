package user

import (
	"rest-api/internal/domain"
	"rest-api/internal/dtos/user"
	"rest-api/internal/infra/repositories"
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

	userEntity, err := domain.NewUser(
		createUserRequest.Name,
		createUserRequest.Email,
		createUserRequest.Password,
	)

	if err != nil {
		return nil, err
	}

	user, err := self.UserRepository.Add(userEntity)

	if err != nil {
		return nil, err
	}

	return user, nil
}
