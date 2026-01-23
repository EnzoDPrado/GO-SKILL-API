package domain

import (
	"fmt"
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

	if err := user.prepare(); err != nil {
		return nil, err
	}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (user *User) Auth(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return fmt.Errorf("invalid credentials")
	}

	return nil
}

func (user *User) prepare() error {
	user.ID = uuid.New()
	user.Status = true
	user.Role = GUEST

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return err
	}

	user.Password = string(encryptedPassword)

	return nil
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
