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

	return product.ToProto(), nil
}

func (ps *ProductService) FindProductByIDs(ctx context.Context, req *pb.ProductRequest) (*pb.Products, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.Dump(ctx),
		"req": utils.Dump(req),
	})

	ids := []int{}
	for _, id := range req.Ids {
		ids = append(ids, int(id))
	}

	products, err := ps.productUsecase.FindByIDs(ctx, ids)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var pbProducts []*pb.Product
	for _, product := range products {
		pbProducts = append(pbProducts, product.ToProto())
	}

	return &pb.Products{
		Products: pbProducts,
	}, nil
}

func (ps *ProductService) ReserveProducts(ctx context.Context, req *pb.ReserveProductRequest) (*pb.ReserveProductResponse, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.Dump(ctx),
		"req": utils.Dump(req),
	})

	reserves := []model.ReserveRequest{}
	for _, product := range req.Products {
		reserves = append(reserves, model.ReserveRequest{
			ID:  int(product.Id),
			Qty: int(product.Qty),
		})
	}

	var res pb.ReserveProductResponse

	err := ps.productUsecase.ReserveProducts(ctx, reserves)
	if err != nil {
		logger.Error(err)
		res.Error = err.Error()
		return &res, nil
	}

	return &pb.ReserveProductResponse{
		Error: "",
	}, nil
}
