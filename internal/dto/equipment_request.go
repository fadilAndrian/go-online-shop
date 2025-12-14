package dto

type CreateEquipmentRequest struct {
	Name         string `validate:"required,min=3"`
	CurrentStock int    `validate:"required,gt=0" json:"current_stock"`
}

type UpdateEquipmentRequest struct {
	Name         string `validate:"required,min=3"`
	CurrentStock int    `validate:"required,gt=0" json:"current_stock"`
}
