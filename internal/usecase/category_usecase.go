package usecase

import (
	"context"

	"github.com/kodinggo/product-service-gb1/internal/model"
	"github.com/kodinggo/product-service-gb1/internal/utils"
	"github.com/sirupsen/logrus"
)

type categoryUsecase struct {
	categoryRepo model.CategoryRepository
}

func NewCategoryUsecase(categoryRepo model.CategoryRepository) model.CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
	}
}

// FindAll is a usecase function to find all categories
func (cu *categoryUsecase) FindAll(ctx context.Context, query model.CategoryQuery) ([]model.Category, error) {
	logger := logrus.WithField("query", utils.Dump(query))

	categories, err := cu.categoryRepo.FindAll(ctx, query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return categories, nil
}

// FindByID is a usecase function to find a category by its ID
func (cu *categoryUsecase) FindByID(ctx context.Context, id int) (model.Category, error) {
	logger := logrus.WithField("id", id)

	category, err := cu.categoryRepo.FindByID(ctx, id)
	if err != nil {
		logger.Error(err)
		return model.Category{}, err
	}

	return category, nil
}

// Create is a usecase function to create a new category
func (cu *categoryUsecase) Create(ctx context.Context, category model.Category) (model.Category, error) {
	logger := logrus.WithField("category", utils.Dump(category))

	newCategory, err := cu.categoryRepo.Create(ctx, category)
	if err != nil {
		logger.Error(err)
		return model.Category{}, err
	}

	return newCategory, nil
}

// Update is a usecase function to update a category
func (cu *categoryUsecase) Update(ctx context.Context, id int, category model.Category) (model.Category, error) {
	logger := logrus.WithFields(logrus.Fields{
		"id":       id,
		"category": utils.Dump(category),
	})

	updatedCategory, err := cu.categoryRepo.Update(ctx, id, category)
	if err != nil {
		logger.Error(err)
		return model.Category{}, err
	}

	return updatedCategory, nil
}

// Delete is a usecase function to delete a category
func (cu *categoryUsecase) Delete(ctx context.Context, id int) error {
	logger := logrus.WithField("id", id)

	if err := cu.categoryRepo.Delete(ctx, id); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
