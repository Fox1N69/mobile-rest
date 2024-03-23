package parser

type ParserI interface {
	GetData() error
}

func (p *Parser) GetData() error {
	return nil
}
