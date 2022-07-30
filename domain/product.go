package domain

import (
	"context"

	"github.com/google/uuid"
)

type Product struct {
	ID    *uuid.UUID `json:"id" validate:"required"`
	Stock *int64     `json:"stock" validate:"required"`
	Name  *string    `json:"name" validate:"required"`
}

type ProductUsecase interface {
	Store(context.Context, *Product) error
	Fetch(ctx context.Context) ([]*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id *uuid.UUID) error
}

// ProductRepository represent the product's repository contract
type ProductRepository interface {
	Fetch(ctx context.Context) (res []*Product, err error)
	Update(ctx context.Context, product *Product) error
	Store(ctx context.Context, a *Product) error
	Delete(ctx context.Context, id *uuid.UUID) error
}
