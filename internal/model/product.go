package model

import (
	"context"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	CategoryID  int       `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	Photos      []Photo   `json:"photos"`
}

type Photo struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type ProductQuery struct {
	Name string
	Sort string
	Size int
	Page int
}

type ProductRepository interface {
	FindAll(ctx context.Context, query ProductQuery) ([]Product, error)
	FindByID(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Delete(ctx context.Context, id int) error
}

type ProductUsecase interface {
	FindAll(ctx context.Context, query ProductQuery) ([]Product, error)
	FindByID(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, product Product) (Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Delete(ctx context.Context, id int) error
}
