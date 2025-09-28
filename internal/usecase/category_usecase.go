package usecase

import (
	"context"
	"pakyus_commerce/internal/entity"
	"pakyus_commerce/internal/model"
	"pakyus_commerce/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryUsecase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	CategoryRepository *repository.CategoryRepository
}

func NewCategoryUsecase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	categoryRepository *repository.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		CategoryRepository: categoryRepository,
	}
}

func (c *CategoryUsecase) FindAll(ctx context.Context) ([]model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var categories []entity.Category
	if err := c.CategoryRepository.FindAll(tx, &categories); err != nil {
		c.Log.WithError(err).Error("error fetching categories")
		return nil, err
	}

	err := tx.Commit().Error
	if err != nil {
		c.Log.WithError(err).Error("error fetching categories")
		return nil, err
	}

	// Convert to response
	responses := make([]model.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = model.CategoryResponse{
			ID:       category.ID,
			Name:     category.Name,
			ParentId: category.ParentId,
		}
	}

	return responses, nil
}
