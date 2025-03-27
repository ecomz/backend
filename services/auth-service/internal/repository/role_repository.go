package repository

import (
	"github.com/ecomz/backend/auth-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type RoleRepository interface {
	GetRoleByID(roleID int) (*model.Role, error)
	GetAllRoles() ([]*model.Role, error)
	CreateRole(name string) error
	DeleteRole(roleID int) error
}

type roleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetRoleByID(roleID int) (*model.Role, error) {
	var role model.Role
	err := r.db.Get(&role, "SELECT * FROM roles WHERE id = $1", roleID)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetAllRoles() ([]*model.Role, error) {
	var roles []*model.Role
	if err := r.db.Select(&roles, "SELECT * FROM roles ORDER BY name ASC"); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) CreateRole(name string) error {
	_, err := r.db.Exec("INSERT INTO roles (name) VALUES ($1)", name)
	return err
}

func (r *roleRepository) DeleteRole(roleID int) error {
	_, err := r.db.Exec("DELETE FROM roles WHERE role_id = ?", roleID)
	return err
}
