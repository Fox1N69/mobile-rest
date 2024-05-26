package parser

import (
	"mobile/internal/app/models"
	"mobile/internal/pkg/database"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

func FetchAboutContent(url string) {
	res, err := http.Get(url)
	if err != nil {
		logrus.Fatalf("Ошибка при запросе к URL: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		logrus.Fatal("Статус ответа сервера:%d%s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.Fatal("Ошибка при создании goquery документа")
	}

	title := doc.Find("h1").Text()
	var contents []string

	doc.Find(".entry").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		// Проверяем, является ли элемент тегом 'h2'
		if s.Is("h2") {
			content = "<strong>" + content + "</strong>"
		}
		contents = append(contents, content)
	})

	// Соединяем все элементы content в одну строку
	fullContent := strings.Join(contents, " ")

	// Создаем одну запись с title и content
	data := models.AboutOrganization{Title: title, Content: fullContent}
	result := database.DB.Create(&data)
	if result.Error != nil {
		logrus.Fatal("Ошибка при сохранении в базу данных")
	}
}
