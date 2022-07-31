package utils

import (
	"context"
	"product-stock-web-service/domain"
)

// ProductCtxKey is a key used for the Product object in the context
type ProductCtxKey struct{}

// Get product from context
func GetProductFromCtx(ctx context.Context) (*domain.Product, error) {
	product, ok := ctx.Value(ProductCtxKey{}).(*domain.Product)
	if !ok {
		return nil, domain.ErrBadParamInput
	}

	return product, nil
}
