package usecase

import (
	"context"
	"errors"
	"product-stock-web-service/domain"
	"product-stock-web-service/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDelete(t *testing.T) {
	name := "name"
	stock := int64(10)
	id := int64(12)
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{
		Name:  &name,
		Stock: &stock,
		ID:    &id,
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Delete", mock.Anything, mock.AnythingOfType("*int64")).Return(nil).Once()

		productUsecase := NewProductUsecase(mockProductRepo, time.Second*2)
		err := productUsecase.Delete(context.TODO(), mockProduct.ID)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("Delete", mock.Anything, mock.AnythingOfType("*int64")).Return(errors.New("unable to delete item")).Once()

		productUsecase := NewProductUsecase(mockProductRepo, time.Second*2)
		err := productUsecase.Delete(context.TODO(), mockProduct.ID)

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	name := "name"
	stock := int64(10)
	id := int64(12)
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{
		Name:  &name,
		Stock: &stock,
		ID:    &id,
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Product")).Return(nil).Once()

		productUsecase := NewProductUsecase(mockProductRepo, time.Second*2)
		err := productUsecase.Update(context.TODO(), &mockProduct)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Product")).Return(errors.New("Error updating the DB")).Once()

		productUsecase := NewProductUsecase(mockProductRepo, time.Second*2)
		err := productUsecase.Update(context.TODO(), &mockProduct)

		assert.Error(t, err)
	})

}

func TestStore(t *testing.T) {
	name := "name"
	stock := int64(10)
	mockProductRepo := new(mocks.ProductRepository)
	mockProduct := domain.Product{
		Name:  &name,
		Stock: &stock,
	}	

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Product")).Return(nil).Once()

		productUsecase := NewProductUsecase(mockProductRepo, time.Second*2)
		err := productUsecase.Store(context.TODO(), &mockProduct)

		assert.NoError(t, err)
	})

	
	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Product")).Return(errors.New("unexpected Error while storing")).Once()

		productUsecase := NewProductUsecase(mockProductRepo, time.Second*2)
		err := productUsecase.Store(context.TODO(), &mockProduct)

		assert.Error(t, err)
	})
}

// TODO Fetch Test
