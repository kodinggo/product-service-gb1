package model

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"->"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c Category) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

type CategoryQuery struct {
	Name string `query:"name"`
	Sort string `query:"sort"`
	Size int    `query:"size"`
	Page int    `query:"page"`
}

type CategoryUsecase interface {
	FindAll(ctx context.Context, query CategoryQuery) ([]Category, error)
	FindByID(ctx context.Context, id int) (Category, error)
	Create(ctx context.Context, category Category) (Category, error)
	Update(ctx context.Context, id int, category Category) (Category, error)
	Delete(ctx context.Context, id int) error
}

type CategoryRepository interface {
	FindAll(ctx context.Context, query CategoryQuery) ([]Category, error)
	FindByID(ctx context.Context, id int) (Category, error)
	Create(ctx context.Context, category Category) (Category, error)
	Update(ctx context.Context, id int, category Category) (Category, error)
	Delete(ctx context.Context, id int) error
}
