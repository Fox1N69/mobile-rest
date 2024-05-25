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
		return c.Status(400).JSON(fiber.Map{"error": "Request parsing failed"})
	}

	emailBody := fmt.Sprintf("Имя: %s\n Фамилия: %s\n Отчество: %s\n Направление: %s\n Группа: %s\n Количество: %s\n Текст: %s",
		data.FirstName, data.LastName, data.Patronymic, data.Specialty, data.Group, data.Quantity, data.Message)

	err := emailService.SendMail("maksimow-pasha1707@mail.ru", "Новая заявка", emailBody)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to send email"})
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
		return c.Status(400).JSON(fiber.Map{"error": "Request parsing failed"})
	}

	emailBody := fmt.Sprintf("ФИО Студента: %s\n Направление: %s\n Группа: %s\n Название военкомата: %s\n Текст: %s",
		data.Fio, data.Specialty, data.Group, data.ArmyName, data.Message)

	if err := emailService.SendMail("maksimow-pasha1707@mail.ru", "Новая заявка", emailBody); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to send email"})
	}

	return c.JSON(fiber.Map{
		"message": "Email sent successfully",
		"data":    data,
	})
}

func (h *Handler) SendScholarshipForm(c fiber.Ctx) error {
	emailService := service.NewEmailService()

	var data models.ScholarshipForm
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request parsing failed"})
	}

	emailBody := fmt.Sprintf(" ФИО Студента:    %s\n Направление:    %s\n Группа:    %s\n Период выплат:     %s\n Количество:    %s",
		data.Fio, data.Specialty, data.Group, data.PaymentPeriod, data.Quantity)

	if data.SendByEmail {
		emailBody += "\n\n Способ получения:    Отправить на почту \n"
	} else if data.PickupInOffice {
		emailBody += "\n\n Способ получения:    Студент заберет справку самостоятельно\n"
	}

	emailBody += fmt.Sprintf("\n\n Почта куда отправить спраку:    %s\n Номер телефона:   %s\n Сообщение:   %s",
		data.Email, data.PhoneNumber, data.Message)

	if err := emailService.SendMail("maksimow-pasha1707@mail.ru", "Новая заявка", emailBody); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to send email",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Email sent successfully",
		"data":    data,
	})
}
