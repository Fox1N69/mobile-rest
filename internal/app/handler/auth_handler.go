package handler

import "github.com/gofiber/fiber/v3"

func (h *Handler) Login(c fiber.Ctx) error {
	return h.AuthController.Login(c)
}

func (h *Handler) Register(c fiber.Ctx) error {
	return h.AuthController.Register(c)
}

func (h *Handler) Logout(c fiber.Ctx) error {
	return h.AuthController.Logout(c)
}
