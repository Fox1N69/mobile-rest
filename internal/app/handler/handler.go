package handler

import (
	"mobile/internal/app/controller"

	"gorm.io/gorm"
)

type Handler struct {
	AuthController *controller.AuthController
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{AuthController: controller.NewAuthController(db)}
}
