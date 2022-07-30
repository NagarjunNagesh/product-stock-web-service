package mysql

import (
	"context"
	"database/sql"
	"product-stock-web-service/domain"

	"github.com/google/uuid"
)

type productRepository struct {
	Conn *sql.DB
}

// NewProductRepository will create an object that represent the product.Repository interface
func NewProductRepository(Conn *sql.DB) domain.ProductRepository {
	return &productRepository{Conn}
}

func (a *productRepository) Fetch(ctx context.Context) (res []*domain.Product, err error) {
	return nil, nil
}

func (a *productRepository) Update(ctx context.Context, product *domain.Product) error {
	return nil
}

func (a *productRepository) Store(ctx context.Context, product *domain.Product) error {
	return nil
}

func (a *productRepository) Delete(ctx context.Context, id *uuid.UUID) error {
	return nil
}
