package repository

import (
	"pakyus_commerce/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	Log *logrus.Logger
}

func NewCategoryRepository(log *logrus.Logger) *CategoryRepository {
	return &CategoryRepository{
		Log: log,
	}
}

func (r *CategoryRepository) FindAll(db *gorm.DB, categories *[]entity.Category) error {
	return db.Find(categories).Error
}
