package repository

import (
	"gorm.io/gorm"
)

type AuthRepositoryI interface {
}

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepostiroy(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

