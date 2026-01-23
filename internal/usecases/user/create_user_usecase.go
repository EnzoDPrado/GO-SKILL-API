package user

import (
	"fmt"
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

	err = self.validate(createUserRequest)

	if err != nil {
		return nil, err
	}

	user, err := self.UserRepository.Add(userEntity)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (self *CreateUserUseCase) validate(createUserRequest user.CreateUserRequest) error {
	err := self.validateEmail(createUserRequest.Email)

	if err != nil {
		return err
	}

	return nil
}

func (self *CreateUserUseCase) validateEmail(email string) error {
	existsUserWithEmail, err := self.UserRepository.ExistsByEmail(email)

	if err != nil {
		return err
	}

	if existsUserWithEmail {
		return fmt.Errorf("user with email %s already exists", email)
	}

	return nil
}
