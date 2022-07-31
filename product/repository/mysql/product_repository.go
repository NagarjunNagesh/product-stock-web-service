package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"product-stock-web-service/domain"

	"github.com/sirupsen/logrus"
)

type productRepository struct {
	Conn *sql.DB
}

// NewProductRepository will create an object that represent the product.Repository interface
func NewProductRepository(Conn *sql.DB) domain.ProductRepository {
	return &productRepository{Conn}
}

func (p *productRepository) Fetch(ctx context.Context) (res []*domain.Product, err error) {
	rows, err := p.Conn.QueryContext(ctx, fetchProductQuery)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result := make([]*domain.Product, 0)
	for rows.Next() {
		product := domain.Product{}
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Stock,
			&product.UpdatedAt,
			&product.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, &product)
	}

	return result, nil
}

func (p *productRepository) Update(ctx context.Context, product *domain.Product) error {
	stmt, err := p.Conn.PrepareContext(ctx, updateProduct)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, product.Name, product.Stock, product.ID)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return err
	}

	// TODO Return Updated date and created date along with the response (Depending on the functional requirements)

	return nil
}

func (p *productRepository) Store(ctx context.Context, product *domain.Product) error {
	stmt, err := p.Conn.PrepareContext(ctx, insertProduct)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, product.Name, product.Stock)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	product.ID = &lastID

	// TODO set the updated time and creation date or remove them depending on the functional requirements
	return err
}

func (p *productRepository) Delete(ctx context.Context, id *int64) error {
	stmt, err := p.Conn.PrepareContext(ctx, deleteProduct)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAfected)
		return err
	}

	return nil
}
