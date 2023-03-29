package repository

import (
	"context"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/services/logger"

	"gorm.io/gorm"
)

type repositoryFunc func(new bool) *gorm.DB

type Repository struct {
	Exchange repositoryFunc
	Ticker   repositoryFunc
	Coin     repositoryFunc
}

func GetRepository(db *gorm.DB, log logger.Logger) *Repository {
	return &Repository{
		Exchange: func(new bool) *gorm.DB {
			return createModel(db, &entities.Exchange{}, new)
		},
		Ticker: func(new bool) *gorm.DB {
			return createModel(db, &entities.Ticker{}, new)
		},
		Coin: func(new bool) *gorm.DB {
			return createModel(db, &entities.Coin{}, new)
		},
	}
}

func createModel(db *gorm.DB, model any, new bool) *gorm.DB {
	ctx := db.Statement.Context
	if new {
		ctx = context.Background()
	}
	return db.Model(&model).Session(&gorm.Session{NewDB: new}).WithContext(ctx)
}
