package usecase

import (
	"context"

	"github.com/kodinggo/product-service-gb1/internal/model"
	"github.com/kodinggo/product-service-gb1/internal/utils"
	"github.com/sirupsen/logrus"
)

type productUsecase struct {
	productRepo model.ProductRepository
}

// NewProductUsecase is a function to create a new product usecase
func NewProductUsecase(productRepo model.ProductRepository) model.ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

// FindAll is a usecase function to find all products
func (pu *productUsecase) FindAll(ctx context.Context, query model.ProductQuery) ([]model.Product, error) {
	logger := logrus.WithField("query", utils.Dump(query))

	products, err := pu.productRepo.FindAll(ctx, query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return products, nil
}

// FindByID is a usecase function to find a product by its ID
func (pu *productUsecase) FindByID(ctx context.Context, id int) (model.Product, error) {
	logger := logrus.WithField("id", id)

	product, err := pu.productRepo.FindByID(ctx, id)
	if err != nil {
		logger.Error(err)
		return model.Product{}, err
	}

	return product, nil
}

// Create is a usecase function to create a new product
func (pu *productUsecase) Create(ctx context.Context, product model.Product) (model.Product, error) {
	logger := logrus.WithField("product", utils.Dump(product))

	product, err := pu.productRepo.Create(ctx, product)
	if err != nil {
		logger.Error(err)
		return model.Product{}, err
	}

	return product, nil
}

// Update is a usecase function to update a product
func (pu *productUsecase) Update(ctx context.Context, product model.Product) (model.Product, error) {
	logger := logrus.WithField("product", utils.Dump(product))

	product, err := pu.productRepo.Update(ctx, product)
	if err != nil {
		logger.Error(err)
		return model.Product{}, err
	}

	return product, nil
}

// Delete is a usecase function to delete a product
func (pu *productUsecase) Delete(ctx context.Context, id int) error {
	logger := logrus.WithField("id", id)

	err := pu.productRepo.Delete(ctx, id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// FindByIDs is a usecase function to find products by their IDs
func (pu *productUsecase) FindByIDs(ctx context.Context, ids []int) ([]model.Product, error) {
	logger := logrus.WithField("ids", ids)

	products, err := pu.productRepo.FindByIDs(ctx, ids)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return products, nil
}

// ReserveProducts is a usecase function to reserve products
func (pu *productUsecase) ReserveProducts(ctx context.Context, reserve []model.ReserveRequest) error {
	logger := logrus.WithField("reserve", utils.Dump(reserve))

	err := pu.productRepo.ReserveProducts(ctx, reserve)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
