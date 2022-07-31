package mysql

import (
	"context"
	"product-stock-web-service/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	name := "name"
	stock := int64(10)
	mockProduct := domain.Product{
		Name:  &name,
		Stock: &stock,
	}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "stock", "updated_at", "created_at"}).
			AddRow(mockProduct.ID, mockProduct.Name, mockProduct.Stock,
				mockProduct.UpdatedAt, mockProduct.CreatedAt)

		mock.ExpectQuery(fetchProductQuery).WillReturnRows(rows)

		productUsecase := NewProductRepository(db)
		products, err := productUsecase.Fetch(context.TODO())

		assert.NoError(t, err)
		assert.NotNil(t, products)
		assert.Equal(t, name, *products[len(products)-1].Name)
	})
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	t.Run("success", func(t *testing.T) {
		id := int64(12)
		prep := mock.ExpectPrepare(deleteProduct)
		prep.ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(12, 1))

		productUsecase := NewProductRepository(db)
		err := productUsecase.Delete(context.TODO(), &id)

		assert.NoError(t, err)
	})
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	name := "name"
	stock := int64(10)
	productId := int64(12)
	mockProduct := domain.Product{
		Name:  &name,
		Stock: &stock,
		ID:    &productId,
	}

	t.Run("success", func(t *testing.T) {
		updateProductQuery := "UPDATE product SET name=\\?, stock=\\? WHERE id=\\?"
		prep := mock.ExpectPrepare(updateProductQuery)
		prep.ExpectExec().WithArgs(name, stock, productId).WillReturnResult(sqlmock.NewResult(12, 1))

		productUsecase := NewProductRepository(db)
		err := productUsecase.Update(context.TODO(), &mockProduct)

		assert.NoError(t, err)
	})
}

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	name := "name"
	stock := int64(10)
	mockProduct := domain.Product{
		Name:  &name,
		Stock: &stock,
	}

	t.Run("success", func(t *testing.T) {
		insertProductQuery := "INSERT product SET name=\\?, stock=\\?"
		prep := mock.ExpectPrepare(insertProductQuery)
		prep.ExpectExec().WithArgs(name, stock).WillReturnResult(sqlmock.NewResult(12, 1))

		productUsecase := NewProductRepository(db)
		err := productUsecase.Store(context.TODO(), &mockProduct)

		assert.NoError(t, err)
	})
}
