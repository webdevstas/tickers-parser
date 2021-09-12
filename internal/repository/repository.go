package repository

import (
	"gorm.io/gorm"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/services/logger"
)

type Repositories struct {
	Exchange *gorm.DB
}

func GetRepositories(db *gorm.DB, log logger.Logger) *Repositories {
	return &Repositories{
		Exchange: db.Model(&entities.Exchange{}),
	}
}
