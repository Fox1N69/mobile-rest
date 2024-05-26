package service

import (
	"fmt"
	"io/ioutil"
	"mobile/internal/app/models"
	"mobile/internal/pkg/database"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func Download() {
	url := "https://kcpt72.ru/schedule/"
	urldownload := ParsKCPT(url)

	excelfile := make([]string, 0)
	for _, i := range urldownload {
		filename := strings.Split(i, "/")
		if filename[3] == "spreadsheets" {
			excelfile = append(excelfile, fmt.Sprintf("https://drive.google.com/u/0/uc?id=%s&export=download", filename[5]))
		}
	}

	path := "Excel"
	num := 1
	for _, i := range excelfile {
		response, err := http.Get(i)
		if err != nil {
			fmt.Println("Ошибка при загрузке файла:", err)
			return
		}
		defer response.Body.Close()

		fileData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Ошибка при чтении данных файла:", err)
			return
		}

		filePath := fmt.Sprintf("%s/Schedule%d.xlsx", path, num)
		err = ioutil.WriteFile(filePath, fileData, 0644)
		if err != nil {
			fmt.Println("Ошибка при записи файла:", err)
			return
		}

		num++
	}

	twoTable("Excel/Schedule2.xlsx", "Excel/Schedule1.xlsx")
}

func ParsKCPT(url string) []string {
	arrAGoogle := make([]string, 0)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при загрузке веб-страницы:", err)
		return arrAGoogle
	}
	defer response.Body.Close()

	htmlContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении данных веб-страницы:", err)
		return arrAGoogle
	}

	respBodyStr := string(htmlContent)
	splitResp := strings.Split(respBodyStr, "<a")
	for _, s := range splitResp {
		if strings.Contains(s, "docs.google.com") {
			startIndex := strings.Index(s, "href=\"") + 6
			endIndex := strings.Index(s[startIndex:], "\"") + startIndex
			arrAGoogle = append(arrAGoogle, s[startIndex:endIndex])
		}
	}

	return arrAGoogle
}

func twoTable(pathfile1 string, pathfile2 string) {
	// Открываем первый файл
	workbook1, err := excelize.OpenFile(pathfile1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Открываем второй файл
	workbook2, err := excelize.OpenFile(pathfile2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Создаем новую книгу Excel
	mergedWorkbook := excelize.NewFile()
	mergedSheet := "Sheet1" // Имя нового листа
	mergedWorkbook.SetSheetName("Sheet1", mergedSheet)

	// Сливаем данные из первого файла
	rows1, err := workbook1.GetRows(mergedSheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows1 {
		// Преобразуем строки в интерфейсы для передачи в SetSheetRow
		interfaceRow := make([]interface{}, len(row))
		for i, v := range row {
			interfaceRow[i] = v
		}
		_ = mergedWorkbook.SetSheetRow(mergedSheet, row[0], &interfaceRow)
	}

	// Сливаем данные из второго файла
	rows2, err := workbook2.GetRows(mergedSheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows2 {
		// Преобразуем строки в интерфейсы для передачи в SetSheetRow
		interfaceRow := make([]interface{}, len(row))
		for i, v := range row {
			interfaceRow[i] = v
		}
		_ = mergedWorkbook.SetSheetRow(mergedSheet, row[0], &interfaceRow)
	}

	// Сохраняем объединенную книгу в новый файл
	if err := mergedWorkbook.SaveAs("Excel/2table.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func AddToDatabase(Raspisanie map[string]map[string][][]interface{}) error {
	for date, rasp := range Raspisanie {
		for group, urok := range rasp {
			for _, i := range urok {
				for _, lessonData := range i {
					lessonSlice, ok := lessonData.([]interface{})
					if !ok {
						fmt.Println("Error: lessonData is not a slice of interfaces")
						continue
					}
					subjectID := lessonSlice[1].(uint)
					prepodID := lessonSlice[3].(uint)
					parsedDate, err := time.Parse("02.01.2006", date)
					if err != nil {
						fmt.Println("Error parsing date:", err)
						continue
					}
					lesson := models.Urok{
						Number:    lessonSlice[0].(int),
						SubjectID: subjectID,
						PrepodID:  prepodID,
						Date:      parsedDate,
						GroupID:   group,
					}
					if len(lessonSlice) >= 4 {
						lesson.Classroom = lessonSlice[2].(string)
					}
					if err := database.DB.Create(&lesson).Error; err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func ParsTwoTable() (map[string]map[string][][]string, error) {
	workbook, err := excelize.OpenFile("Excel/2table.xlsx")
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file: %v", err)
	}
	defer os.Remove("Excel/2table.xlsx") // Clean up after reading

	sheetName := workbook.GetSheetName(1)
	rows, err := workbook.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows from sheet: %v", err)
	}

	var unformatted [][]string
	var raspis [][]string // Changed from []string to [][]string
	output := make(map[string]map[string][][]string)

	for _, row := range rows {
		if len(row) > 0 && row[0] != "" && strings.Contains(row[0], "День") {
			if len(raspis) > 0 {
				unformatted = append(unformatted, raspis...)
			}
			dateParts := strings.Split(row[0], ",")
			if len(dateParts) > 1 {
				raspis = [][]string{{strings.TrimSpace(dateParts[1])}} // Initialize with a slice of slices
			}
		} else if row[0] == "№" {
			for _, cell := range row {
				if cell != "№" && cell != "" {
					raspis = append(raspis, []string{cell}) // Append a new slice for each cell
				}
			}
		} else if _, err := strconv.Atoi(row[0]); err == nil {
			for i := 1; i < len(raspis); i++ {
				raspis[i] = append(raspis[i], row...) // Append the entire row to each group slice
			}
		} else if row[0] == "" && row[2] != "Ауд." {
			for i := 1; i < len(raspis); i++ {
				raspis[i] = append(raspis[i], row[1:]...) // Append the modified row to each group slice
			}
		}
	}

	if len(raspis) > 0 {
		unformatted = append(unformatted, raspis...)
	}

	for _, entry := range unformatted {
		date := strings.TrimSpace(entry[0]) // Correctly access the first string of the first slice
		output[date] = make(map[string][][]string)
		for _, data := range entry[1:] {
			group := string(data[0])
			output[date][group] = append(output[date][group], []string{data[1:]})
		}
	}

	return output, nil
}
