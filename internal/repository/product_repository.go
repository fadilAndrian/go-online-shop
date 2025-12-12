package repository

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (p *ProductRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product

	err := p.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) FindById(id int64) (*domain.Product, error) {
	var product domain.Product

	err := p.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepository) Create(product *domain.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductRepository) Update(product *domain.Product) error {
	return p.DB.Save(product).Error
}

func (p *ProductRepository) Delete(product *domain.Product) error {
	return p.DB.Delete(product).Error
}
