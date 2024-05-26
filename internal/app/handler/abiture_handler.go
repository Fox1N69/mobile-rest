package handler

import (
	"mobile/internal/app/parser"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) AbiturePage(c fiber.Ctx) error {
	parser.FetchAbiturePage("https://kcpt72.ru/abitur/")
	return c.JSON("OK")
}
