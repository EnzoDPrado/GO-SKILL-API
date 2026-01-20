package repositories

import "rest-api/internal/domain"

type UserRepository interface {
	GetAll() ([]domain.User, error)
	Add(newUser domain.User) (domain.User, error)
}

type UserRepositoryDb struct {
	users []domain.User
}

func New() *UserRepositoryDb {
	return &UserRepositoryDb{}
}

func (ur *UserRepositoryDb) GetAll() ([]domain.User, error) {
	return ur.users, nil
}

func (ur *UserRepositoryDb) Add(newUser domain.User) (domain.User, error) {
	ur.users = append(ur.users, newUser)

	return newUser, nil
}
