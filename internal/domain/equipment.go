package domain

type Equipment struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	CurrentStock int `json:"current_stock"`
}

func (Equipment) TableName() string {
	return "equipments"
}
