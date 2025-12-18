package domain

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
