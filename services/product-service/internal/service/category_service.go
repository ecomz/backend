package service

import (
	"github.com/ecomz/backend/libs/logger"
	"github.com/ecomz/backend/product-service/internal/dto"
	"github.com/ecomz/backend/product-service/internal/model"
	"github.com/ecomz/backend/product-service/internal/repository"
	"go.uber.org/zap"
)

type CategoryService interface {
	CreateCategory(data *dto.CreateCategoryRequest) (*model.Category, error)
	GetAllCategories() ([]*model.Category, error)
	GetCategoryByID(id int) (*model.Category, error)
	UpdateCategory(id int, data *dto.UpdateCategoryRequest) error
	DeleteCategory(id int) error
}

type categoryService struct {
	logger logger.Logger
	repo   repository.CategoryRepository
}

func NewCategoryService(logger logger.Logger, categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		logger: logger,
		repo:   categoryRepository,
	}
}

func (c *categoryService) CreateCategory(data *dto.CreateCategoryRequest) (*model.Category, error) {
	category, err := c.repo.CreateCategory(data)
	if err != nil {
		c.logger.Error("failed to create category", zap.Error(err))
		return nil, err
	}
	c.logger.Info("successfuly create category", zap.Int("categoryID", category.ID))
	return category, err
}

func (c *categoryService) GetAllCategories() ([]*model.Category, error) {
	categories, err := c.repo.GetAllCategories()
	if err != nil {
		c.logger.Error("failed to get all categories", zap.Error(err))
		return nil, err
	}
	c.logger.Info("successfuly to get all categories")
	return categories, err
}

func (c *categoryService) GetCategoryByID(id int) (*model.Category, error) {
	category, err := c.repo.GetCategoryByID(id)
	if err != nil {
		c.logger.Error("failed to get category by id", zap.Error(err), zap.Int("id", id))
		return nil, err
	}
	return category, err
}

func (c *categoryService) UpdateCategory(id int, data *dto.UpdateCategoryRequest) error {
	if err := c.repo.UpdateCategory(id, data); err != nil {
		c.logger.Error("failed to update category", zap.Error(err), zap.Int("id", id))
		return err
	}
	c.logger.Info("successfuly to update category", zap.Any("category", data))
	return nil
}

func (c *categoryService) DeleteCategory(id int) error {
	if err := c.repo.DeleteCategory(id); err != nil {
		c.logger.Error("failed to update category", zap.Error(err), zap.Int("id", id))
		return err
	}
	c.logger.Info("successfuly to update category", zap.Int("id", id))
	return nil
}
