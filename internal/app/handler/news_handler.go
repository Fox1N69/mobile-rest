package handler

import (
	"mobile/internal/app/models"
	"mobile/internal/app/parser"
	"mobile/internal/pkg/database"
	"strconv"

	"github.com/gofiber/fiber/v3"
)


func (h *Handler) GetAllNews(c fiber.Ctx) error {
	var news []models.NewsData

	if err := database.DB.Find(&news).Error; err != nil {
		return err
	}

	return c.JSON(news)
}

func (h *Handler) GetFullNews(c fiber.Ctx) error {
	newsIDStr := c.Params("id")
	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid ID"})
	}

	var fullNews models.FullNewsData
	if err := database.DB.Where("news_data_id = ?", newsID).First(&fullNews).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "news not found"})
	}

	return c.JSON(fullNews)
}

func (h *Handler) TriggerParseFullNews(c fiber.Ctx) error {
	err := parser.ParseFullNews()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to parse full news"})
	}
	return c.SendStatus(200)
}

func (h *Handler) TriggerParseNews(c fiber.Ctx) error {
	_, err := parser.ParseNews()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to parse full news"})
	}

	return c.SendStatus(200)
}
