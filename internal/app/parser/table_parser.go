package parser

type TableParser struct {
}

type TableParserI interface {
}

func NewTableParser() *TableParser {
	return &TableParser{}
}

