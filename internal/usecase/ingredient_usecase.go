package usecase

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"github.com/fadilAndrian/go-online-shop/internal/dto"
	repointerface "github.com/fadilAndrian/go-online-shop/internal/repository/interface"
)

type IngredientUsecase struct {
	r repointerface.IngredientRepositoryInteface
}

func NewIngredientUsecase(r repointerface.IngredientRepositoryInteface) *IngredientUsecase {
	return &IngredientUsecase{r}
}

func (uc *IngredientUsecase) ListIngredient() ([]domain.Ingredient, error) {
	return uc.r.FindAll()
}

func (uc *IngredientUsecase) GetIngredient(id int64) (*domain.Ingredient, error) {
	return uc.r.FindById(id)
}

func (uc *IngredientUsecase) CreateIngredient(req dto.IngredientCreateRequest) (*domain.Ingredient, error) {
	ingredient := &domain.Ingredient{
		Name:         req.Name,
		Unit:         req.Unit,
		InitialStock: req.InitialStock,
		CurrentStock: req.CurrentStock,
	}

	if err := uc.r.Create(ingredient); err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (uc *IngredientUsecase) UpdateIngredient(id int64, req dto.IngredientUpdateRequest) (*domain.Ingredient, error) {
	ingredient, err := uc.r.FindById(id)
	if err != nil {
		return nil, err
	}

	ingredient.Name = req.Name
	ingredient.Unit = req.Unit
	ingredient.CurrentStock = req.CurrentStock

	if err := uc.r.Update(ingredient); err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (uc *IngredientUsecase) DeleteIngredient(id int64) error {
	ingredient, err := uc.r.FindById(id)
	if err != nil {
		return err
	}

	return uc.r.Delete(ingredient)
}
