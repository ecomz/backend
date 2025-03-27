package model

import (
	"database/sql"
	"time"

	_ "github.com/go-playground/validator/v10"
)

type Role struct {
	ID        int          `json:"id" db:"id"`
	Name      string       `json:"name" db:"name" validate:"required"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeleteAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
