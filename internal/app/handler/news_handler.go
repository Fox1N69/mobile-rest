package handler

import (
	"mobile/internal/app/models"
	"mobile/internal/pkg/database"
	"strconv"

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

		result := database.DB.Create(&newsItem)
		if result.Error != nil {
			c.JSON("error")
		}

		// After saving NewsData, save FullNewsData
		fullNewsItem := models.FullNewsData{
			NewsDataID: newsItem.ID, // ID from the just saved NewsData
		}

		result = database.DB.Create(&fullNewsItem)
		if result.Error != nil {
			c.JSON("error")
		}
	})

	return c.JSON(news)
}

func (h *Handler) ParseFullNews(c fiber.Ctx) error {
	var newsLinks []models.NewsData
	// Fetch all news data including links and IDs
	if err := database.DB.Find(&newsLinks).Error; err != nil {
		return err
	}

	for _, newsData := range newsLinks {
		doc, err := goquery.NewDocument(newsData.Link)
		if err != nil {
			continue // Handle error appropriately
		}

		// Parsing data from the page
		fullContent := doc.Find(".entry").Text() // Use the correct selector for your case

		fullNewsItem := models.FullNewsData{
			NewsDataID: newsData.ID, // Set the NewsDataID
			Title: newsData.Title,
			Content:    fullContent,
			Link:       newsData.Link,
		}

		// Update the full news item in the database
		result := database.DB.Model(&models.FullNewsData{}).Where("news_data_id = ?", newsData.ID).Updates(fullNewsItem)
		if result.Error != nil {
			continue // Handle error appropriately
		}
	}

	return c.JSON(fiber.Map{"message": "success"})
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