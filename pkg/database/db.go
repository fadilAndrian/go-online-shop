package database

import (
	"github.com/fadilAndrian/go-online-shop/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate()
	return db, nil
}
