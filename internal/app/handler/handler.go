package handler

import (
	"mobile/internal/app/controller"
	"mobile/internal/app/repository"

	"gorm.io/gorm"
)

type Handler struct {
	AuthController     *controller.AuthController
	ScheduleRepository *repository.ScheduleRespository
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{AuthController: controller.NewAuthController(db), ScheduleRepository: repository.NewScheduleRespository(db)}
}
