package domain

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Name string    `gorm:"column:name;type:notnull"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}
