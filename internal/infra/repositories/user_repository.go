package repositories

import (
	"rest-api/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	GetById(id [16]byte) (*domain.User, error)
	DeleteById(id [16]byte) error
	Add(newUser *domain.User) (*domain.User, error)
	ExistsByEmail(email string) (bool, error)
	UpdateUserRole(userId [16]byte, role string) error
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

func (ur *UserRepositoryDb) DeleteById(userId [16]byte) error {
	result := ur.Db.Model(&domain.User{}).
		Where("id = ?", userId).
		Where("status = ?", true).
		Update("status", false)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepositoryDb) UpdateUserRole(userId [16]byte, role string) error {
	result := ur.Db.Model(&domain.User{}).
		Where("id = ?", userId).
		Where("status = ?", true).
		Update("role", role)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepositoryDb) GetById(id [16]byte) (*domain.User, error) {
	var user *domain.User

	result := ur.Db.
		Where("id = ?", id).
		Where("status = ?", true).
		Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepositoryDb) GetByEmail(email string) (*domain.User, error) {
	var user *domain.User

	result := ur.Db.
		Where("email = ?", email).
		Where("status = ?", true).
		Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepositoryDb) ExistsByEmail(email string) (bool, error) {
	var userQuantity int64

	result := ur.Db.Model(&domain.User{}).
		Where("email = ?", email).
		Where("status = ?", true).
		Count(&userQuantity)

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
