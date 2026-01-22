package domain

import (
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type UserRole string

const (
	GUEST UserRole = "guest"
	ADMIN UserRole = "admin"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Name     string    `valid:"required,alpha,length(2|50)" gorm:"not null"`
	Email    string    `valid:"required,email" gorm:"uniqueIndex;not null"`
	Password string    `valid:"required,length(6|125)" gorm:"not null"`
	Status   bool      `valid:"required" gorm:"not null"`
	Role     UserRole  `valid:"required,length(0|55)" gorm:"not null"`
}

func NewUser(name string, email string, password string) (*User, error) {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	user.prepare()

	if error := user.validate(); error != nil {
		return nil, error
	}

	return &user, nil
}

func (user *User) prepare() {
	user.ID = uuid.New()
	user.Status = true
	user.Role = GUEST
}

func (user *User) validate() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		log.Printf("userId %v", user.ID)
		log.Printf(err.Error())
		return err
	}

	return nil
}
