package handler

import (
	"mobile/internal/app/models"
	"mobile/internal/pkg/database"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v3"
)

func ParseNews() ([]models.NewsData, error) {
	doc, err := goquery.NewDocument("https://kcpt72.ru/category/%D0%BD%D0%BE%D0%B2%D0%BE%D1%81%D1%82%D1%8C/")
	if err != nil {
		return nil, err
	}

	var news []models.NewsData
	doc.Find(".new").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".card-title").Text()
		content := s.Find(".entry").Text()
		link, _ := s.Find("a.more-link").Attr("href")

		var count int64
		database.DB.Model(&models.NewsData{}).Where("link = ?", link).Count(&count)
		if count > 0 {
			return
		}

		newsItem := models.NewsData{
			Title:   title,
			Content: content,
			Link:    link,
		}
		news = append(news, newsItem)

		database.DB.Create(&newsItem)
	})

	return news, nil
}

func ParseFullNews() (error) {
	var newsLinks []models.NewsData
	database.DB.Find(&newsLinks)

	for _, newsData := range newsLinks {
		doc, err := goquery.NewDocument(newsData.Link)
		if err != nil {
			continue
		}

		fullContent := doc.Find(".entry").Text()
		fullNewsItem := models.FullNewsData{
			NewsDataID: newsData.ID,
			Title:      newsData.Title,
			Content:    fullContent,
			Link:       newsData.Link,
		}
		database.DB.Model(&models.FullNewsData{}).Where("news_data_id = ?", newsData.ID).Updates(fullNewsItem)
	}
	return nil
}

func (h *Handler) GetAllNews(c fiber.Ctx) error {
	var news []models.NewsData

	if err := database.DB.Find(&news).Error; err != nil {
		return err
	}

	return c.JSON(news)
}

func (h *Handler) GetFullNews(c fiber.Ctx) error {
	newsIDStr := c.Params("id")
	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}

	var fullNews models.FullNewsData
	if err := database.DB.Where("news_data_id = ?", newsID).First(&fullNews).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "news not found"})
	}

	return c.JSON(fullNews)
}
