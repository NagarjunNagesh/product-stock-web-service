package domain

import (
	"context"
	"time"
)

type Product struct {
	ID        *int64    `json:"id"`
	Stock     *int64    `json:"stock" validate:"required"`
	Name      *string   `json:"name" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductUsecase interface {
	Store(ctx context.Context, product *Product) error
	Fetch(ctx context.Context) ([]*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id *int64) error
}

// ProductRepository represent the product's repository contract
type ProductRepository interface {
	Fetch(ctx context.Context) (res []*Product, err error)
	Update(ctx context.Context, product *Product) error
	Store(ctx context.Context, a *Product) error
	Delete(ctx context.Context, id *int64) error
}
