package repositories

import (
	"rest-api/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]*domain.User, error)
	Add(newUser *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func New() *UserRepositoryDb {
	return &UserRepositoryDb{}
}

func (ur *UserRepositoryDb) GetAll() ([]*domain.User, error) {
	var users []*domain.User

	err := ur.Db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepositoryDb) Add(newUser *domain.User) (*domain.User, error) {
	err := ur.Db.Create(&newUser).Error

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
