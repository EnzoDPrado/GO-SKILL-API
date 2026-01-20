package user

import (
	"rest-api/internal/dtos/user"
	"rest-api/internal/repositories"
)

type GetAllUsersUseCase struct {
	UserRepository repositories.UserRepository
}

func NewGetAllUsersUseCase(userRepository repositories.UserRepository) *GetAllUsersUseCase {
	return &GetAllUsersUseCase{
		UserRepository: userRepository,
	}
}

func (self *GetAllUsersUseCase) Execute() ([]user.UserSimplifiedResponse, error) {
	users, err := self.UserRepository.GetAll()

	if err != nil {
		return nil, err
	}

	var usersResponse []user.UserSimplifiedResponse

	for _, userMap := range users {
		usersResponse = append(usersResponse, user.UserSimplifiedResponse{
			ID:   userMap.ID,
			Name: userMap.Name,
		})

	}

	return usersResponse, nil
}
