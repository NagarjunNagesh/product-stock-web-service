package middleware

import (
	"net/http"
	"product-stock-web-service/domain"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ProductHandler  represent the httphandler for product
type ProductHandler struct {
	ProductUseCase domain.ProductUsecase
}

// NewProductHandler will initialize the products/ resources endpoint
func NewProductHandler(e *echo.Echo, us domain.ProductUsecase) {
	handler := &ProductHandler{
		ProductUseCase: us,
	}
	e.GET("/products", handler.Fetch)
	e.POST("/products", handler.Store)
	e.GET("/products/:id", handler.Update)
	e.DELETE("/products/:id", handler.Delete)
}

// Fetch will fetch the product based on given params
func (a *ProductHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	product, err := a.ProductUseCase.Fetch(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, product)
}

// Store will store the product by given request body
func (a *ProductHandler) Store(c echo.Context) (err error) {
	var product domain.Product
	err = c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&product); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.ProductUseCase.Store(ctx, &product)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

// Update will store the product by given request body
func (a *ProductHandler) Update(c echo.Context) (err error) {
	var product domain.Product
	err = c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&product); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.ProductUseCase.Update(ctx, &product)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

// Delete will delete product by given param
func (a *ProductHandler) Delete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ctx := c.Request().Context()

	err = a.ProductUseCase.Delete(ctx, &id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
