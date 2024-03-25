package main

import (
	"mobile/internal/app/handler"
	"mobile/internal/pkg/database"
	"mobile/internal/pkg/routers"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Mobile-REST",
	})

	db := database.InitDB()

	handler := handler.NewHandler(db)

	router := routers.NewRouter(*handler)
	router.InitRouter(app)

	app.Listen(":8000")
}
