package mock

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"github.com/fadilAndrian/go-online-shop/internal/dto"
)

type ProductUsecaseMock struct {
	ListFn   func() ([]domain.Product, error)
	GetFn    func(id int64) (*domain.Product, error)
	CreateFn func(req dto.CreateProductRequest) (*domain.Product, error)
	UpdateFn func(id int64, req dto.UpdateProductRequest) (*domain.Product, error)
	DeleteFn func(id int64) error
}

func (m *ProductUsecaseMock) ListProduct() ([]domain.Product, error) {
	return m.ListFn()
}

func (m *ProductUsecaseMock) GetProduct(id int64) (*domain.Product, error) {
	return m.GetFn(id)
}

func (m *ProductUsecaseMock) CreateProduct(req dto.CreateProductRequest) (*domain.Product, error) {
	return m.CreateFn(req)
}

func (m *ProductUsecaseMock) UpdateProduct(id int64, req dto.UpdateProductRequest) (*domain.Product, error) {
	return m.UpdateFn(id, req)
}

func (m *ProductUsecaseMock) DeleteProduct(id int64) error {
	return m.DeleteFn(id)
}
