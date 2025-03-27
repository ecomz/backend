package dto

type RoleRequest struct {
	Name string `json:"name" validate:"required"`
}
