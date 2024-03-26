package controller

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
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

func GenerateJWTToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// install claim the token
	claim := token.Claims.(jwt.MapClaims)
	claim["userID"] = userID
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("secret key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// password hashing function
func hashPassword(password []byte) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return hashed, err
}

func checkPassword(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}

func (ac *AuthController) Login(c fiber.Ctx) error {
	//reading data from the body
	return nil
}

func (ac *AuthController) Register(c fiber.Ctx) error {
	return nil
}

func (ac *AuthController) Logout(c fiber.Ctx) error {
	return nil
}
