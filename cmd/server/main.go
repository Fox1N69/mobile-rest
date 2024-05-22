package main

import (
	"mobile/internal/app/handler"
	"mobile/internal/pkg/database"
	"mobile/internal/pkg/routers"

	"github.com/gofiber/fiber/v3"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Mobile-REST",
	})

	db := database.GetDB()

	handlerInstance := handler.NewHandler(db) // Изменено имя переменной

	c := cron.New()
	c.AddFunc("0 0 * * *", func() {
		_, err := handler.ParseNews()
		if err != nil {
			logrus.Infoln("Error pars news cron: ", err)
		}
		logrus.Info("ParserNews seccus")
	})
	c.AddFunc("1 0 * * *", func() {
		err := handler.ParseFullNews()
		if err != nil {
			logrus.Infoln("Error pars full news cron: ", err)
		}
		logrus.Info("ParserFullNews seccus")

	})
	c.Start()

	router := routers.NewRouter(handlerInstance) // Удалено разыменование
	router.InitRouter(app)

	app.Listen(":8000")

	select {}
}
