package dto

import "time"

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3,max=50"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"omitempty,min=3,max=50"`
}

type CreateCategoryResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
