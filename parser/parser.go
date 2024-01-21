package parser

type Parser struct {
	minutes []int
	hours   []int
	days    []int
	months  []int
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) string {
	return ""
}
