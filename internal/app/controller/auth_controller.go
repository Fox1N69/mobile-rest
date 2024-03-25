package controller

import (
	"encoding/json"
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

func hashPassword(password []byte) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return hashed, err
}

func (ac *AuthController) Login(c fiber.Ctx) error {
	return nil
}

func (ac *AuthController) Register(c fiber.Ctx) error {
	user := new(models.AuthUser)
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return err
	}

	// check for a username
	var existingUser models.AuthUser
	if result := ac.DB.Where("username = ?", user.Username).First(&existingUser); result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "A user with such an username alredy existing"})
	}

	//hashed password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = []byte(hashedPassword)

	//create user
	if result := ac.DB.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Error when creating user"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User successfully create"})
}

func (ac *AuthController) Logout(c fiber.Ctx) error {
	return nil
}
