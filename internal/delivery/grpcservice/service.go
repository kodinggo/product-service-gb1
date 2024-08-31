package grpcservice

import (
	"context"

	"github.com/kodinggo/product-service-gb1/internal/model"
	"github.com/kodinggo/product-service-gb1/internal/utils"
	pb "github.com/kodinggo/product-service-gb1/pb/product"
	"github.com/sirupsen/logrus"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	productUsecase model.ProductUsecase
}

func NewProductService(pu model.ProductUsecase) *ProductService {
	return &ProductService{
		productUsecase: pu,
	}
}

func (ps *ProductService) FindProductByID(ctx context.Context, req *pb.ProductRequest) (*pb.Product, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.Dump(ctx),
		"req": utils.Dump(req),
	})

	product, err := ps.productUsecase.FindByID(ctx, int(req.Id))
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &pb.Product{
		Id:          int32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}
