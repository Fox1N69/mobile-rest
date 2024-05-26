package routers

import (
	"mobile/internal/app/handler"

	"github.com/gofiber/fiber/v3"
)

type Router struct {
	handler *handler.Handler
}

func NewRouter(h *handler.Handler) *Router {
	return &Router{handler: h}
}

func (r *Router) InitRouter(app *fiber.App) {
	api := app.Group("/api")
	{
		api.Post("/emailTraning", r.handler.SendTraningForm)
		api.Post("/emailArmy", r.handler.SendArmyForm)
		api.Post("/emailPayment", r.handler.SendScholarshipForm)
		api.Get("/about", r.handler.GetAboutInfo)

		news := api.Group("/news")
		{
			news.Get("/", r.handler.GetAllNews)
			news.Get("/:id/full", r.handler.GetFullNews)

		}
	}

	admin := app.Group("/admin")
	{
		admin.Get("/GetChangeData", r.handler.GetChangeData)
	}

	test := app.Group("/test")
	{
		test.Get("/parsAbout", r.handler.ParserAboutO)
		test.Get("/pars", r.handler.TriggerParseNews)
		test.Get("/fullpars", r.handler.TriggerParseFullNews)
		test.Get("/abiture", r.handler.AbiturePage)
	}

	auth := app.Group("/auth")
	{
		auth.Post("/login", r.handler.Login)
		auth.Post("/register", r.handler.Register)
		auth.Post("/logout", r.handler.Logout)
	}
}
