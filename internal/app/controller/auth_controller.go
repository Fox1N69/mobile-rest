package controller

import (
	"mobile/internal/app/models"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthControllerI interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
	Logout(c fiber.Ctx) error
	RefreshToken(c fiber.Ctx) error
	Restricted(c fiber.Ctx) error
}

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

func hashPassword(user *models.AuthUser) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	return string(hashed), err
}

func (ac *AuthController) Login(c fiber.Ctx) error {
	return nil
}

func (ac *AuthController) Register(c fiber.Ctx) error {
	return nil
}

func (ac *AuthController) Logout(c fiber.Ctx) error {
	return nil
}
