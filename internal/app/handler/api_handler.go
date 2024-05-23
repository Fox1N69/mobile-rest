package handler

import (
	"encoding/json"
	"fmt"
	"mobile/internal/app/models"
	"mobile/internal/app/service"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) SendTraningForm(c fiber.Ctx) error {
	emailService := service.NewEmailService()

	var data models.FormDataAboutTraning
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request parsing failed"})
	}

	emailBody := fmt.Sprintf("Имя: %s\nФамилия: %s\nОтчество: %s\nНаправление: %s\nГруппа: %s\nКоличество: %d\nТекст: %s",
		data.FirstName, data.LastName, data.Patronymic, data.Direction, data.Group, data.Quantity, data.Message)

	err := emailService.SendMail("maksimow-pasha1707@mail.ru", "Новая заявка", emailBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email"})
	}

	return c.JSON(fiber.Map{
		"message": "Email sent successfully",
		"data":    data,
	})
}

func (h *Handler) SendArmyForm(c fiber.Ctx) error {
	emailService := service.NewEmailService()

	var data models.FormArmy
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request parsing failed"})
	}

	emailBody := fmt.Sprintf("ФИО: %s\n Направление: %s\n Группа: %s\n Название военкомата: %s\n Текст: %s",
		data.Fio, data.Direction, data.Group, data.ArmyName, data.Message)

	if err := emailService.SendMail("maksimow-pasha1707@mail.ru", "Новая заявка", emailBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email"})
	}

	return c.JSON(fiber.Map{
		"message": "Email sent successfully",
		"data":    data,
	})
}
