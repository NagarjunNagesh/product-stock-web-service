package usecase

import (
	"context"
	"product-stock-web-service/domain"
	"product-stock-web-service/utils"
	"time"
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

func (a *productUsecase) Store(ctx context.Context, product *domain.Product) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	// TODO sql injection attack security chck
	// TODO check if product already exists

	if err := utils.ValidateStruct(ctx, product); err != nil {
		return domain.ErrBadParamInput
	}

	err := a.productRepo.Store(ctx, product)
	return err
}

func (a *productUsecase) Fetch(ctx context.Context) ([]*domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	// TODO introduce Pagination (limit and offset) to enable query

	products, err := a.productRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (a *productUsecase) Update(ctx context.Context, product *domain.Product) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	// TODO sql injection attack security chck

	return a.productRepo.Update(ctx, product)
}

func (a *productUsecase) Delete(ctx context.Context, id *int64) error {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	// Check if product exists
	// check if authorized

	return a.productRepo.Delete(ctx, id)
}
