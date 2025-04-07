package repository

import (
	"database/sql"

	"github.com/ecomz/backend/product-service/internal/dto"
	"github.com/ecomz/backend/product-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	CreateCategory(data *dto.CreateCategoryRequest) (*model.Category, error)
	GetAllCategories() ([]*model.Category, error)
	GetCategoryByID(id int) (*model.Category, error)
	UpdateCategory(id int, data *dto.UpdateCategoryRequest) error
	DeleteCategory(id int) error
}

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(data *dto.CreateCategoryRequest) (*model.Category, error) {
	category := &model.Category{
		Name: data.Name,
	}

	query := `INSERT INTO categories (name, created_at, updated_at) VALUES ($1, NOW(), NOW()) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, category.Name).Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) GetAllCategories() ([]*model.Category, error) {
	var categories []*model.Category
	err := r.db.Select(&categories, "SELECT * FROM categories")
	return categories, err
}

func (r *categoryRepository) GetCategoryByID(id int) (*model.Category, error) {
	var category model.Category

	err := r.db.Get(&category, "SELECT * FROM categories WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &category, err
}

func (r *categoryRepository) UpdateCategory(id int, data *dto.UpdateCategoryRequest) error {
	_, err := r.db.Exec("UPDATE categories SET name=$1, updated_at=NOW() WHERE id=$2", data.Name, id)
	return err
}

func (r *categoryRepository) DeleteCategory(id int) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id=$1", id)
	return err
}
