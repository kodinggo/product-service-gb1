package model

import (
	"context"
	"time"

	pb "github.com/kodinggo/product-service-gb1/pb/product"
	"gorm.io/gorm"
)

type Product struct {
	ID          int            `json:"id"`
	CategoryID  int            `json:"category_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Stock       int            `json:"stock"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Photos      []Photo        `json:"photos" gorm:"foreignKey:ProductID"`
}

func (p Product) ToProto() *pb.Product {
	var photos []*pb.Photo
	if len(p.Photos) > 0 {
		for _, photo := range p.Photos {
			photos = append(photos, photo.ToProto())
		}
	}
	return &pb.Product{
		Id:          int32(p.ID),
		CategoryId:  int32(p.CategoryID),
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       int32(p.Stock),
		CreatedAt:   p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   p.UpdatedAt.Format(time.RFC3339),
		Photos:      photos,
	}
}

type Photo struct {
	ID        int            `json:"id"`
	ProductID int            `json:"product_id"`
	URL       string         `json:"url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (p Photo) ToProto() *pb.Photo {
	return &pb.Photo{
		Id:        int32(p.ID),
		ProductId: int32(p.ProductID),
		Url:       p.URL,
		CreatedAt: p.CreatedAt.Format(time.RFC3339),
		UpdatedAt: p.UpdatedAt.Format(time.RFC3339),
	}
}

type ProductQuery struct {
	Name string `query:"name"`
	Sort string `query:"sort"`
	Size int    `query:"size"`
	Page int    `query:"page"`
}

type ProductRepository interface {
	FindAll(ctx context.Context, query ProductQuery) ([]Product, error)
	FindByID(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Delete(ctx context.Context, id int) error
	FindByIDs(ctx context.Context, ids []int) ([]Product, error)
}

type ProductUsecase interface {
	FindAll(ctx context.Context, query ProductQuery) ([]Product, error)
	FindByID(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Delete(ctx context.Context, id int) error
	FindByIDs(ctx context.Context, ids []int) ([]Product, error)
}
