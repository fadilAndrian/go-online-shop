package domain

import "time"

type Ingredient struct {
	ID           int64  `gorm:"primaryKey"`
	Name         string `gorm:"unique"`
	Unit         string
	InitialStock int `json:"initial_stock"`
	CurrentStock int `json:"current_stock"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
