package service

import (
	"github.com/ecomz/backend/libs/logger"
	"github.com/ecomz/backend/product-service/internal/dto"
	"github.com/ecomz/backend/product-service/internal/model"
	"github.com/ecomz/backend/product-service/internal/repository"
	"go.uber.org/zap"
)

type ProductService interface {
	CreateProduct(data *dto.CreateProductRequest) (*model.Product, error)
	GetAllProducts() ([]*model.Product, error)
	GetProductByID(id int) (*model.Product, error)
	UpdateProduct(id int, data *dto.UpdateProductRequest) error
	DeleteProduct(id int) error
}

type productService struct {
	logger logger.Logger
	repo   repository.ProductRepository
}

func NewProductService(logger logger.Logger, productRepository repository.ProductRepository) ProductService {
	return &productService{
		logger: logger,
		repo:   productRepository,
	}
}

func (c *productService) CreateProduct(data *dto.CreateProductRequest) (*model.Product, error) {
	product, err := c.repo.CreateProduct(data)
	if err != nil {
		c.logger.Error("failed to create product", zap.Error(err))
		return nil, err
	}
	c.logger.Info("successfuly create product", zap.Int("productID", product.ID))
	return product, err
}

func (c *productService) GetAllProducts() ([]*model.Product, error) {
	products, err := c.repo.GetAllProducts()
	if err != nil {
		c.logger.Error("failed to get all products", zap.Error(err))
		return nil, err
	}
	c.logger.Info("successfuly to get all products")
	return products, err
}

func (c *productService) GetProductByID(id int) (*model.Product, error) {
	product, err := c.repo.GetProductByID(id)
	if err != nil {
		c.logger.Error("failed to get product by id", zap.Error(err), zap.Int("id", id))
		return nil, err
	}
	return product, err
}

func (c *productService) UpdateProduct(id int, data *dto.UpdateProductRequest) error {
	if err := c.repo.UpdateProduct(id, data); err != nil {
		c.logger.Error("failed to update product", zap.Error(err), zap.Int("id", id))
		return err
	}
	c.logger.Info("successfuly to update product", zap.Any("product", data))
	return nil
}

func (c *productService) DeleteProduct(id int) error {
	if err := c.repo.DeleteProduct(id); err != nil {
		c.logger.Error("failed to update product", zap.Error(err), zap.Int("id", id))
		return err
	}
	c.logger.Info("successfuly to update product", zap.Int("id", id))
	return nil
}
