package interfaces

import "github.com/fadilAndrian/go-online-shop/internal/domain"

type ProductRepositoryInterface interface {
	FindAll() ([]domain.Product, error)
	FindById(id int64) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(product *domain.Product) error
}
