package usecase

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"github.com/fadilAndrian/go-online-shop/internal/dto"
	"github.com/fadilAndrian/go-online-shop/internal/repository"
)

type EquipmentUsecase struct {
	r repository.EquipmentRepository
}

func NewEquipmentUsecase(r repository.EquipmentRepository) *EquipmentUsecase {
	return &EquipmentUsecase{r}
}

func (u *EquipmentUsecase) ListEquipment() ([]domain.Equipment, error) {
	return u.r.FindAll()
}

func (u *EquipmentUsecase) GetEquipment(id int64) (*domain.Equipment, error) {
	return u.r.FindById(id)
}

func (u *EquipmentUsecase) CreateEquipment(req dto.CreateEquipmentRequest) (*domain.Equipment, error) {
	equipment := domain.Equipment{
		Name:         req.Name,
		CurrentStock: req.CurrentStock,
	}

	if err := u.r.Create(&equipment); err != nil {
		return nil, err
	}

	return &equipment, nil
}

func (u *EquipmentUsecase) UpdateEquipment(id int64, req dto.UpdateEquipmentRequest) (*domain.Equipment, error) {
	equipment, err := u.r.FindById(id)
	if err != nil {
		return nil, err
	}

	equipment.Name = req.Name
	equipment.CurrentStock = req.CurrentStock

	if err := u.r.Update(equipment); err != nil {
		return nil, err
	}

	return equipment, nil
}

func (u *EquipmentUsecase) DeleteEquipment(id int64) error {
	equipment, err := u.r.FindById(id)
	if err != nil {
		return err
	}

	return u.r.Delete(equipment)
}
