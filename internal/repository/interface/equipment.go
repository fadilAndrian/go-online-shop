package interfaces

import "github.com/fadilAndrian/go-online-shop/internal/domain"

type EquipmentRepositoryInterface interface {
	FindAll() ([]domain.Equipment, error)
	FindById(id int64) (*domain.Equipment, error)
	Create(e *domain.Equipment) error
	Update(e *domain.Equipment) error
	Delete(e *domain.Equipment) error
}
