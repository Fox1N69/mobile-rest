package repository

import (
	"mobile/internal/app/models"

	"gorm.io/gorm"
)

type AuthRepositoryI interface {
	CreateUser(user *models.User) error
	GetAllUsers(user *models.User) error
	GetUserByID(id uint) (string, error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepostiroy(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (ar *AuthRepository) CreateUser(user *models.User) error {
	return ar.DB.Save(user).Error
}

func (ar *AuthRepository) GetAllUsers(user []*models.User) ([]*models.User, error) {
	result := ar.DB.Find(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ar *AuthRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := ar.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ar *AuthRepository) UpdateUser(user *models.User) error {
	return nil
}
func (ar *AuthRepository) DeleteUser(user *models.User) error {
	return ar.DB.Delete(user).Error
}
