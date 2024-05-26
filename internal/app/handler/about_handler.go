package handler

import (
	"mobile/internal/app/models"
	"mobile/internal/app/parser"
	"mobile/internal/pkg/database"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) ParserAboutO(c fiber.Ctx) error {
	parser.FetchAboutContent("https://kcpt72.ru/sveden/common/")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pars is OK"})
}

func (h *Handler) GetAboutInfo(c fiber.Ctx) error {
	var data []models.AboutOrganization

	if err := database.DB.Find(&data).Error; err != nil {
		return err
	}

	return c.JSON(data)
}
