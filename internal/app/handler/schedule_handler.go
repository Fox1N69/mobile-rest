package handler

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

var months = map[int]string{
	1:  "января",
	2:  "февраля",
	3:  "марта",
	4:  "апреля",
	5:  "мая",
	6:  "июня",
	7:  "июля",
	8:  "августа",
	9:  "сентября",
	10: "октября",
	11: "ноября",
	12: "декабря",
}

var daysOfWeek = map[int]string{
	0: "понедельник",
	1: "вторник",
	2: "среда",
	3: "четверг",
	4: "пятница",
	5: "суббота",
	6: "воскресенье",
}

func (h *Handler) GetChangeData(c fiber.Ctx) error {
	return nil
}

func (h *Handler) ApiChangesDay(c fiber.Ctx) error {
	dateParam := c.Query("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid date format",
		})
	}

	change, err := h.ScheduleRepository.GetChangesByDay(date)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"data": change,
	})
}

func (h *Handler) ApiClassDay(c fiber.Ctx) error {
	dateParam := c.Query("date")
	group := c.Query("group")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format"})
	}
	// Здесь нужно вызвать функцию Database.GetScheduleWithChangesDay и вернуть результат в формате JSON
	h.ScheduleRepository.GetScheduleWithChangesDay(date, group)
	return c.JSON(fiber.Map{"data": "Replace this with the actual data from Database.GetScheduleWithChangesDay"})
}

func (h *Handler) ApiClassWeek(c fiber.Ctx) error {
	output := make(map[string]interface{})
	dateParam := c.Query("date")
	group := c.Query("group")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format"})
	}
	for i := 0; i < 7; i++ {
		currentDate := date.AddDate(0, 0, i)
		schedule := h.ScheduleRepository.GetScheduleWithChangesDay(currentDate, group)
		output[currentDate.Format("2006-01-02")] = schedule
	}
	return c.JSON(output)
}

func (h *Handler) ApiTeacherDay(c fiber.Ctx) error {
	data := h.ScheduleRepository.GetPrepodsList()
	return c.JSON(fiber.Map{"data": data})
}
