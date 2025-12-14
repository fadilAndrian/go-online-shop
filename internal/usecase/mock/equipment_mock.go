package mock

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"github.com/fadilAndrian/go-online-shop/internal/dto"
)

type EquipmentUsecaseMock struct {
	ListFn   func() ([]domain.Equipment, error)
	GetFn    func(id int64) (*domain.Equipment, error)
	CreateFn func(req dto.CreateEquipmentRequest) (*domain.Equipment, error)
	UpdateFn func(id int64, req dto.UpdateEquipmentRequest) (*domain.Equipment, error)
	DeleteFn func(id int64) error
}

func (m *EquipmentUsecaseMock) ListEquipment() ([]domain.Equipment, error) {
	return m.ListFn()
}

func (m *EquipmentUsecaseMock) GetEquipment(id int64) (*domain.Equipment, error) {
	return m.GetFn(id)
}

func (m *EquipmentUsecaseMock) CreateEquipment(req dto.CreateEquipmentRequest) (*domain.Equipment, error) {
	return m.CreateFn(req)
}

func (m *EquipmentUsecaseMock) UpdateEquipment(id int64, req dto.UpdateEquipmentRequest) (*domain.Equipment, error) {
	return m.UpdateFn(id, req)
}

func (m *EquipmentUsecaseMock) DeleteEquipment(id int64) error {
	return m.DeleteFn(id)
}
