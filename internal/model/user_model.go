package model

import "github.com/google/uuid"

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Name      string `json:"name,omitempty"`
	Token     string `json:"token,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Name     string `json:"name" validate:"required,max=100"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"-" validate:"required"`
	Password string    `json:"password,omitempty" validate:"max=100"`
	Name     string    `json:"name,omitempty" validate:"max=100"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	ID uuid.UUID `json:"-" validate:"required"`
}

type GetUserRequest struct {
	ID uuid.UUID `json:"-" validate:"required"`
}
