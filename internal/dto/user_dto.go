package dto

import "github.com/google/uuid"

type OutputUser struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type InputCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InputUpdateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
