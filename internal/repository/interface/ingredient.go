package interfaces

import "github.com/fadilAndrian/go-online-shop/internal/domain"

type IngredientRepositoryInteface interface {
	FindAll() ([]domain.Ingredient, error)
	FindById(id int64) (*domain.Ingredient, error)
	Create(i *domain.Ingredient) error
	Update(i *domain.Ingredient) error
	Delete(i *domain.Ingredient) error
}
