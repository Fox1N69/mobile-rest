package handler

import (
	"mobile/internal/app/models"
	"mobile/internal/pkg/database"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v3"
)


func (h *Handler) ParseNews(c fiber.Ctx) error {
	doc, err := goquery.NewDocument("https://kcpt72.ru/category/%D0%BD%D0%BE%D0%B2%D0%BE%D1%81%D1%82%D1%8C/")
	if err != nil {
		return err
	}

	var news []models.NewsData

	doc.Find(".new").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".card-title").Text()
		content := s.Find(".entry").Text()
		link, _ := s.Find("a.more-link").Attr("href")

		newsItem := models.NewsData{
			Title:   title,
			Content: content,
			Link:    link,
		}
		news = append(news, newsItem)




		result := database.DB.Create(&newsItem) // Предполагается, что у вас есть глобальный объект DB
		if result.Error != nil {
			c.JSON("error")
		}
	})

	return c.JSON(news)
}


func (h *Handler) GetAllNews(c fiber.Ctx) error {
	var news []models.NewsData

	if err := database.DB.Find(&news).Error; err != nil {
		return err
	}

	return c.JSON(news)
}


