package repository

import "gorm.io/gorm"

type ApiRepository struct {
	DB *gorm.DB
}

func NewApiRepository(db *gorm.DB) *ApiRepository {
	return &ApiRepository{DB: db}
}
