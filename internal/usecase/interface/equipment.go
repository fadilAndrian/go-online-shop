package interfaces

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"github.com/fadilAndrian/go-online-shop/internal/usecase/dto"
)

type EquipmentUsecase interface {
	ListEquipment() ([]domain.Equipment, error)
	GetEquipment(id int64) (*domain.Equipment, error)
	CreateEquipment(req dto.CreateEquipmentRequest) (*domain.Equipment, error)
	UpdateEquipment(id int64, req dto.UpdateEquipmentRequest) (*domain.Equipment, error)
	DeleteEquipment(id int64) error
}
