package usecase

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	repointerface "github.com/fadilAndrian/go-online-shop/internal/repository/interface"
	"github.com/fadilAndrian/go-online-shop/internal/usecase/dto"
)

type ProductUsecase struct {
	repo repointerface.ProductRepository
}

func NewProductUsecase(repo repointerface.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo}
}

func (r ProductUsecase) ListProduct() ([]domain.Product, error) {
	return r.repo.FindAll()
}

func (r ProductUsecase) GetProduct(id int64) (*domain.Product, error) {
	return r.repo.FindById(id)
}

func (r ProductUsecase) CreateProduct(req dto.CreateProductRequest) (*domain.Product, error) {
	product := domain.Product{
		Name:  req.Name,
		Price: req.Price,
	}

	err := r.repo.Create(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r ProductUsecase) UpdateProduct(id int64, req dto.UpdateProductRequest) (*domain.Product, error) {
	product, err := r.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	product.Name = req.Name
	product.Price = req.Price

	if err := r.repo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (r ProductUsecase) DeleteProduct(id int64) error {
	product, err := r.repo.FindById(id)
	if err != nil {
		return err
	}

	return r.repo.Delete(product)
}
