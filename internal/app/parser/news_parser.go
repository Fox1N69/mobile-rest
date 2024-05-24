package parser

import (
	"mobile/internal/app/models"
	"mobile/internal/pkg/database"

	"github.com/PuerkitoBio/goquery"
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

func ParseFullNews() error {
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
		database.DB.Where(models.FullNewsData{NewsDataID: newsData.ID}).FirstOrCreate(&fullNewsItem)
	}
	return nil
}
