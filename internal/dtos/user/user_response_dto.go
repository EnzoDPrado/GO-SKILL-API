package user

import "github.com/google/uuid"

type UserMinimalResponse struct {
	ID uuid.UUID `json:"id"`
}

type UserSimplifiedResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
