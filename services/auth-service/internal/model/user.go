package model

import (
	"database/sql"
	"time"

	_ "github.com/go-playground/validator/v10"
)

type User struct {
	ID        string       `json:"id" db:"id"`
	Name      string       `json:"name" db:"name" validate:"required,min=3,max=50"`
	Email     string       `json:"email" db:"email" validate:"required,email"`
	Password  string       `json:"password" db:"password" validate:"required,min=6,max=50"`
	RoleID    int          `json:"role_id" db:"role_id" validate:"required,gt=0"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
