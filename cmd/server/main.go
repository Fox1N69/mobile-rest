package main

import (
	"mobile/internal/app/handler"
	"mobile/internal/pkg/routers"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Mobile-REST",
	})

	handler := handler.NewHandler()

	router := routers.NewRouter(*handler)
	router.InitRouter(app)

	app.Listen(":8000")
}
