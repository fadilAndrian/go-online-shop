package dto

type IngredientCreateRequest struct {
	Name         string `validate:"required,min=3"`
	Unit         string `validate:"required"`
	InitialStock int    `validate:"required,gte=0" json:"initial_stock"`
	CurrentStock int    `validate:"required,gte=0" json:"current_stock"`
}

type IngredientUpdateRequest struct {
	Name         string `validate:"required,min=3"`
	Unit         string `validate:"required"`
	CurrentStock int    `validate:"required,gte=0" json:"current_stock"`
}
