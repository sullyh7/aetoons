package db

import (
	"github.com/sullyh7/aetoons/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("aetoons.db"), &gorm.Config{})
}
