package repositories

import (
	"rest-api/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Add(newUser *domain.User) (*domain.User, error)
	ExistsByEmail(email string) (bool, error)
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

func (ur *UserRepositoryDb) GetByEmail(email string) (*domain.User, error) {
	var user *domain.User

	result := ur.Db.Where("email = ?", email).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepositoryDb) ExistsByEmail(email string) (bool, error) {
	var userQuantity int64

	result := ur.Db.Table("users").Where("email = ?", email).Count(&userQuantity)

	if result.Error != nil {
		return false, result.Error
	}

	if userQuantity == 0 {
		return false, nil
	}

	return true, nil
}

func (ur *UserRepositoryDb) Add(newUser *domain.User) (*domain.User, error) {
	err := ur.Db.Create(&newUser).Error

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
