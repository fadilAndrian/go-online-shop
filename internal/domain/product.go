package domain

type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique"`
	Price int
}
