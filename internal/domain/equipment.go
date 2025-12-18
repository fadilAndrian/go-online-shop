package domain

import "time"

type Equipment struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	CurrentStock int `json:"current_stock"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Equipment) TableName() string {
	return "equipments"
}
