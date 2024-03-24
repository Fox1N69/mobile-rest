package routers

import (
	"mobile/internal/app/handler"

	"github.com/gofiber/fiber/v3"
)

type Router struct {
	handler handler.Handler
}

func NewRouter(h handler.Handler) *Router {
	return &Router{handler: h}
}

func (r *Router) InitRouter(app *fiber.App) {
	api := app.Group("/api")
	{
		api.Get("/getAllData", r.handler.GetAllData)
	}

	auth := app.Group("/auth")
	{
		auth.Post("/login", r.handler.Login)
		auth.Post("/register", r.handler.Register)
		auth.Post("/logout", r.handler.Logout)
	}
}
