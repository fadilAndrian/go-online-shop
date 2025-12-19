package mock

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"github.com/fadilAndrian/go-online-shop/internal/dto"
)

type IngredientUsecaseMock struct {
	ListFn   func() ([]domain.Ingredient, error)
	GetFn    func(id int64) (*domain.Ingredient, error)
	CreateFn func(req dto.IngredientCreateRequest) (*domain.Ingredient, error)
	UpdateFn func(id int64, req dto.IngredientUpdateRequest) (*domain.Ingredient, error)
	DeleteFn func(id int64) error
}

func (m *IngredientUsecaseMock) ListIngredient() ([]domain.Ingredient, error) {
	return m.ListFn()
}

func (m *IngredientUsecaseMock) GetIngredient(id int64) (*domain.Ingredient, error) {
	return m.GetFn(id)
}

func (m *IngredientUsecaseMock) CreateIngredient(req dto.IngredientCreateRequest) (*domain.Ingredient, error) {
	return m.CreateFn(req)
}

func (m *IngredientUsecaseMock) UpdateIngredient(id int64, req dto.IngredientUpdateRequest) (*domain.Ingredient, error) {
	return m.UpdateFn(id, req)
}

func (m *IngredientUsecaseMock) DeleteIngredient(id int64) error {
	return m.DeleteFn(id)
}
