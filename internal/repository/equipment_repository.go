package repository

import (
	"github.com/fadilAndrian/go-online-shop/internal/domain"
	"gorm.io/gorm"
)

type EquipmentRepository struct {
	db *gorm.DB
}

func NewEquipmentRepository(db *gorm.DB) *EquipmentRepository {
	return &EquipmentRepository{db}
}

func (r *EquipmentRepository) FindAll() ([]domain.Equipment, error) {
	var equipments []domain.Equipment

	err := r.db.Find(&equipments, domain.Equipment{}).Error
	if err != nil {
		return nil, err
	}

	return equipments, nil
}

func (r *EquipmentRepository) FindById(id int64) (*domain.Equipment, error) {
	var equipment domain.Equipment

	err := r.db.First(&equipment, id).Error
	if err != nil {
		return nil, err
	}

	return &equipment, nil
}

func (r *EquipmentRepository) Create(e *domain.Equipment) error {
	return r.db.Create(&e).Error
}

func (r *EquipmentRepository) Update(e *domain.Equipment) error {
	return r.db.Save(&e).Error
}

func (r *EquipmentRepository) Delete(e *domain.Equipment) error {
	return r.db.Delete(&e).Error
}
