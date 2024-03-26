package parser

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type TableParser struct {
	Service *sheets.Service
}

type TableParserI interface {
	ParseSpreadsheet(spreadsheetID, readRange string) (string, error)
}

func NewTableParser(ctx context.Context) (*TableParser, error) {
	creds, err := google.FindDefaultCredentials(ctx, sheets.SpreadsheetsScope)
	if err != nil {
		return nil, err
	}

	service, err := sheets.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		return nil, err
	}

	return &TableParser{Service: service}, nil
}

func (tp *TableParser) ParseSpreadsheet(spreadsheetID, readRange string) (string, error) {
	resp, err := tp.Service.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
