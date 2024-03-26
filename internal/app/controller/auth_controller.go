package controller

import (
	"encoding/json"
	"mobile/internal/app/models"
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

var secretKey = "secret"

func generateJWTToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte(secretKey))
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
}

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
	loginData := new(models.LoginData)
	if err := json.Unmarshal(c.Body(), &loginData); err != nil {
		return err
	}

	//user search by username
	var user models.User
	result := ac.DB.Where("username = ?", &user).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "User not found"})
	}

	if !checkPassword(user.Password, loginData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Uncorecct password"})
	}

	token, err := generateJWTToken(&user)
	if err != nil {
		return err
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"token": token, "message": "Autorization was successful", "user": user})
}

func (ac *AuthController) Register(c fiber.Ctx) error {
	user := new(models.User)
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return err
	}

	// check for a username
	var existingUser models.User
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

	//generate token
	token, err := generateJWTToken(user.ID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User successfully create", "token": token})
}

func (ac *AuthController) Logout(c fiber.Ctx) error {
	return nil
}
