package interfaces

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"github.com/fadilAndrian/go-online-shop/internal/dto"
)

type IngredientUsecaseInterface interface {
	ListIngredient() ([]domain.Ingredient, error)
	GetIngredient(id int64) (*domain.Ingredient, error)
	CreateIngredient(req dto.IngredientCreateRequest) (*domain.Ingredient, error)
	UpdateIngredient(id int64, req dto.IngredientUpdateRequest) (*domain.Ingredient, error)
	DeleteIngredient(id int64) error
}
