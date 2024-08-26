package repository

import (
	"context"

	"github.com/kodinggo/product-service-gb1/internal/model"
	"github.com/kodinggo/product-service-gb1/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository is a function to create a new category repository
func NewCategoryRepository(db *gorm.DB) model.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

// FindAll is a repository function to find all categories
func (cr *categoryRepository) FindAll(ctx context.Context, query model.CategoryQuery) ([]model.Category, error) {
	logger := logrus.WithField("query", utils.Dump(query))

	var categories []model.Category

	qb := cr.db.WithContext(ctx)

	if query.Name != "" {
		qb.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if query.Sort != "" {
		utils.BuildSortQuery(qb, query.Sort)
	}

	if query.Size != 0 {
		qb.Limit(query.Size)
	}

	if query.Page != 0 {
		offset := (query.Page - 1) * query.Size
		qb.Offset(offset)
	}

	if err := cr.db.Find(&categories).Error; err != nil {
		logger.Error(err)
		return nil, err
	}

	return categories, nil
}

// FindByID is a repository function to find a category by its ID
func (cr *categoryRepository) FindByID(ctx context.Context, id int) (model.Category, error) {
	logger := logrus.WithField("id", id)

	var category model.Category

	if err := cr.db.First(&category, id).Error; err != nil {
		logger.Error(err)
		return model.Category{}, err
	}

	return category, nil
}

// Create is a repository function to create a new category
func (cr *categoryRepository) Create(ctx context.Context, category model.Category) (model.Category, error) {
	logger := logrus.WithField("category", utils.Dump(category))

	if err := cr.db.Create(&category).Error; err != nil {
		logger.Error(err)
		return model.Category{}, err
	}

	return category, nil
}

// Update is a repository function to update a category
func (cr *categoryRepository) Update(ctx context.Context, id int, category model.Category) (model.Category, error) {
	logger := logrus.WithFields(logrus.Fields{
		"id":       id,
		"category": utils.Dump(category),
	})

	if err := cr.db.Model(&model.Category{}).Where("id = ?", id).Updates(category).Error; err != nil {
		logger.Error(err)
		return model.Category{}, err
	}

	return category, nil
}

// Delete is a repository function to delete a category
func (cr *categoryRepository) Delete(ctx context.Context, id int) error {
	logger := logrus.WithField("id", id)

	if err := cr.db.Delete(&model.Category{}, id).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
