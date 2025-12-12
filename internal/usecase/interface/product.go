package interfaces

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"github.com/fadilAndrian/go-online-shop/internal/usecase/dto"
)

type ProductUsecase interface {
	ListProduct() ([]domain.Product, error)
	GetProduct(id int64) (*domain.Product, error)
	CreateProduct(req dto.CreateProductRequest) (*domain.Product, error)
	UpdateProduct(id int64, req dto.UpdateProductRequest) (*domain.Product, error)
	DeleteProduct(id int64) error
}
