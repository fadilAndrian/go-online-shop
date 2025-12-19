package repository

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"gorm.io/gorm"
)

type IngredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) *IngredientRepository {
	return &IngredientRepository{db}
}

func (r *IngredientRepository) FindAll() ([]domain.Ingredient, error) {
	var ingredients []domain.Ingredient

	err := r.db.Find(&ingredients).Error
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

func (r *IngredientRepository) FindById(id int64) (*domain.Ingredient, error) {
	var ingredient *domain.Ingredient

	err := r.db.First(&ingredient, id).Error
	if err != nil {
		return nil, err
	}

	return ingredient, nil
}

func (r *IngredientRepository) Create(i *domain.Ingredient) error {
	return r.db.Create(&i).Error
}

func (r *IngredientRepository) Update(i *domain.Ingredient) error {
	return r.db.Save(&i).Error
}

func (r *IngredientRepository) Delete(i *domain.Ingredient) error {
	return r.db.Delete(&i).Error
}
