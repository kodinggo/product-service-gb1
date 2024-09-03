package repository

import (
	"context"

	"github.com/kodinggo/product-service-gb1/internal/model"
	"github.com/kodinggo/product-service-gb1/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository is a function to create a new product repository
func NewProductRepository(db *gorm.DB) model.ProductRepository {
	return &productRepository{
		db: db,
	}
}

// FindAll is a repository function to find all products
func (pr *productRepository) FindAll(ctx context.Context, query model.ProductQuery) ([]model.Product, error) {
	logger := logrus.WithField("query", utils.Dump(query))

	var products []model.Product

	qb := pr.db.WithContext(ctx)

	if query.Name != "" {
		qb = qb.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if query.Sort != "" {
		sortQuery := utils.BuildSortQuery(query.Sort)
		qb = qb.Order(sortQuery)
	}

	if query.Size != 0 {
		qb = qb.Limit(query.Size)
	}

	if query.Page != 0 {
		offset := (query.Page - 1) * query.Size
		qb = qb.Offset(offset)
	}

	if err := qb.Preload("Photos").Find(&products).Error; err != nil {
		logger.Error(err)
		return nil, err
	}

	return products, nil
}

// FindByID is a repository function to find a product by its ID
func (pr *productRepository) FindByID(ctx context.Context, id int) (model.Product, error) {
	logger := logrus.WithField("id", id)

	var product model.Product

	if err := pr.db.WithContext(ctx).Preload("Photos").First(&product, id).Error; err != nil {
		logger.Error(err)
		return model.Product{}, err
	}

	return product, nil
}

// Create is a repository function to create a new product
func (pr *productRepository) Create(ctx context.Context, product model.Product) (model.Product, error) {
	logger := logrus.WithField("product", utils.Dump(product))

	if err := pr.db.WithContext(ctx).Create(&product).Error; err != nil {
		logger.Error(err)
		return model.Product{}, err
	}

	return product, nil
}

// Update is a repository function to update a product
func (pr *productRepository) Update(ctx context.Context, product model.Product) (model.Product, error) {
	logger := logrus.WithField("product", utils.Dump(product))

	if err := pr.db.WithContext(ctx).Save(&product).Error; err != nil {
		logger.Error(err)
		return model.Product{}, err
	}

	return product, nil
}

// Delete is a repository function to delete a product
func (pr *productRepository) Delete(ctx context.Context, id int) error {
	logger := logrus.WithField("id", id)

	if err := pr.db.WithContext(ctx).Delete(&model.Product{}, id).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// FindByIDs is a repository function to find products by their IDs
func (pr *productRepository) FindByIDs(ctx context.Context, ids []int) ([]model.Product, error) {
	logger := logrus.WithField("ids", ids)

	var products []model.Product

	if err := pr.db.WithContext(ctx).Preload("Photos").Find(&products, ids).Error; err != nil {
		logger.Error(err)
		return nil, err
	}

	return products, nil
}
