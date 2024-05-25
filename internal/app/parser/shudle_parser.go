package parser

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PuerkitoBio/goquery"
)

func download() {
	url := "https://kcpt72.ru/schedule/"
	urldownload := parsKCPT(url)

	var excelfiles []string
	for _, link := range urldownload {
		filename := strings.Split(link, "/")
		if filename[3] == "spreadsheets" {
			excelfiles = append(excelfiles, "https://drive.google.com/u/0/uc?id="+filename[5]+"&export=download")
		}
	}

	path := "Excel"
	os.MkdirAll(path, os.ModePerm)

	for i, fileURL := range excelfiles {
		resp, err := http.Get(fileURL)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		filePath := filepath.Join(path, "Schedule"+strconv.Itoa(i+1)+".xlsx")
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			panic(err)
		}
	}

	twoTable("Excel/Schedule2.xlsx", "Excel/Schedule1.xlsx")
}

func twoTable(pathfile1, pathfile2 string) error {
	workbook1, err := excelize.OpenFile(pathfile1)
	if err != nil {
		return err
	}

	workbook2, err := excelize.OpenFile(pathfile2)
	if err != nil {
		return err
	}

	mergedWorkbook := excelize.NewFile()
	mergedWorkbook.NewSheet("Sheet1")
	currentRow := 1

	for _, sheetName := range workbook1.GetSheetMap() {
		rows := workbook1.GetRows(sheetName)
		for _, row := range rows {
			cell := fmt.Sprintf("A%d", currentRow)
			mergedWorkbook.SetSheetRow("Sheet1", cell, &row)
			currentRow++
		}
	}

	for _, sheetName := range workbook2.GetSheetMap() {
		rows := workbook2.GetRows(sheetName)
		for _, row := range rows {
			cell := fmt.Sprintf("A%d", currentRow)
			mergedWorkbook.SetSheetRow("Sheet1", cell, &row)
			currentRow++
		}
	}

	if err := mergedWorkbook.SaveAs("Excel/2table.xlsx"); err != nil {
		return err
	}

	return nil
}

func parsKCPT(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, "docs.google.com") {
			links = append(links, href)
		}
	})

	return links
}
