package dto

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=50"`
	Description string  `json:"description" validate:"required,max=255"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	CategoryID  int     `json:"category_id" validate:"required,gt=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"omitempty,min=3,max=50"`
	Description string  `json:"description" validate:"omitempty,max=255"`
	Price       float64 `json:"price" validate:"omitempty,gt=0"`
	CategoryID  int     `json:"category_id" validate:"omitempty,gt=0"`
}
