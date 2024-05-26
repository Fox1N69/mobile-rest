package main

import (
	"mobile/internal/app/handler"
	"mobile/internal/app/parser"
	"mobile/internal/pkg/database"
	"mobile/internal/pkg/routers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		ServerHeader: "Mobile-REST",
		Views:        engine,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	db := database.GetDB()

	handlerInstance := handler.NewHandler(db) // Изменено имя переменной

	c := cron.New()
	c.AddFunc("0 0 * * *", func() {
		_, err := parser.ParseNews()
		if err != nil {
			logrus.Infoln("Error pars news cron: ", err)
		}
		logrus.Info("ParserNews seccus")
	})
	c.AddFunc("1 0 * * *", func() {
		err := parser.ParseFullNews()
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
