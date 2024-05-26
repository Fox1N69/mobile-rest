package parser

import (
	"os"

	"github.com/go-rod/rod"
	"github.com/sirupsen/logrus"
)

func FetchAbiturePage(url string) {
	// Соединяемся с браузером
	browser := rod.New().ControlURL("").MustConnect()
	defer browser.MustClose()

	// Создаем новую страницу
	page := browser.MustPage(url)
	defer page.MustClose()

	// Отключаем выполнение JavaScript для извлечения HTML и CSS
	page.MustStopLoading()

	// Извлекаем весь HTML код
	html, err := page.HTML()
	if err != nil {
		logrus.Fatalf("Ошибка при извлечении HTML: %v", err)
	}

	// Извлекаем весь CSS код
	css, err := page.Eval("document.querySelector('style').innerText")
	if err != nil {
		logrus.Fatalf("Ошибка при извлечении CSS: %v", err)
	}

	// Записываем HTML в файл
	err = writeToFile("output.html", html)
	if err != nil {
		logrus.Fatalf("Ошибка при записи HTML в файл: %v", err)
	} else {
		logrus.Println("HTML код записан в файл output.html")
	}

	// Записываем CSS в файл
	err = writeToFile("styles.css", css.Value.Str())
	if err != nil {
		logrus.Fatalf("Ошибка при записи CSS в файл: %v", err)
	} else {
		logrus.Println("CSS код записан в файл styles.css")
	}
}

func writeToFile(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
