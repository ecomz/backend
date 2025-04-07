package repository

import (
	"database/sql"

	"github.com/ecomz/backend/product-service/internal/dto"
	"github.com/ecomz/backend/product-service/internal/model"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	CreateProduct(data *dto.CreateProductRequest) (*model.Product, error)
	GetAllProducts() ([]*model.Product, error)
	GetProductByID(id int) (*model.Product, error)
	UpdateProduct(id int, data *dto.UpdateProductRequest) error
	DeleteProduct(id int) error
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) CreateProduct(data *dto.CreateProductRequest) (*model.Product, error) {
	product := &model.Product{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		CategoryID:  data.CategoryID,
	}

	query := `INSERT INTO products (name, description, price, category_id, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.CategoryID).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepository) GetAllProducts() ([]*model.Product, error) {
	var products []*model.Product
	err := r.db.Select(&products, "SELECT * FROM products")
	return products, err
}

func (r *productRepository) GetProductByID(id int) (*model.Product, error) {
	var product model.Product

	err := r.db.Get(&product, "SELECT * FROM products WHERE id = $1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &product, err
}

func (r *productRepository) UpdateProduct(id int, data *dto.UpdateProductRequest) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, category_id = $4, updated_at = NOW() WHERE id = $5`
	_, err := r.db.Exec(
		query,
		data.Name,
		data.Description,
		data.Price,
		data.CategoryID,
		id,
	)

	return err
}

func (r *productRepository) DeleteProduct(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}
