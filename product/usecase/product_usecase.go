package usecase

import (
	"context"
	"product-stock-web-service/domain"
	"time"

	"github.com/google/uuid"
)

type productUsecase struct {
	productRepo    domain.ProductRepository
	contextTimeout time.Duration
}

// NewProductUsecase will create new an productUsecase object representation of domain.ProductUsecase interface
func NewProductUsecase(a domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *productUsecase) Store(context.Context, *domain.Product) error {
	return nil
}

func (a *productUsecase) Fetch(ctx context.Context) ([]*domain.Product, error) {
	return nil, nil
}

func (a *productUsecase) Update(ctx context.Context, product *domain.Product) error {
	return nil
}

func (a *productUsecase) Delete(ctx context.Context, id *uuid.UUID) error {
	return nil
}
