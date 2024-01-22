package parser

import (
	"cron_expression_parser/parser/consts"
	"cron_expression_parser/parser/helpers"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Parser struct {
	minutes     []int
	hours       []int
	daysOfMonth []int
	daysOfWeek  []int
	months      []int
	command     string
}

var valueToParseMap = map[int]consts.Value{}

func NewParser() *Parser {
	valueToParseMap[0] = &consts.Minutes{Name: "minutes"}
	valueToParseMap[1] = &consts.Hours{Name: "hours"}
	valueToParseMap[2] = &consts.DayOfMonth{Name: "day of month"}
	valueToParseMap[3] = &consts.Month{Name: "month"}
	valueToParseMap[4] = &consts.DayOfWeek{Name: "day of week"}

	return &Parser{}
}

func (p *Parser) Parse(input string) error {
	inputTab, err := getSplitInput(input)
	if err != nil {
		return err
	}
	err = p.performParse(inputTab)
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) PrintCurrentCronExpression() {
	printField("minute", p.minutes)
	printField("hour", p.hours)
	printField("day of month", p.daysOfMonth)
	printField("month", p.months)
	printField("day of week", p.daysOfWeek)
	printStringField("command", p.command)
}

func printField(fieldName string, field []int) {
	fmt.Printf("%-14s", fieldName)
	for i := 0; i < len(field); i++ {
		fmt.Printf("%d ", field[i])
	}
	fmt.Printf("\n")
}

func printStringField(fieldName string, field string) {
	fmt.Printf("%-14s%s", fieldName, field)
}

func getSplitInput(input string) ([]string, error) {
	if len(input) == 0 {
		return []string{}, errors.New("Input length should not be 0")
	}

	split := strings.Split(input, " ")

	if len(split) < 6 {
		return []string{}, errors.New("Input is in wrong format, should be: '* * * * * Command_to_execute'")
	}
	return split, nil
}

func (p *Parser) performParse(slicedInput []string) error {

	for key, val := range valueToParseMap {
		res, err := parsePart(slicedInput[key], val)
		if err != nil {
			return err
		}

		if !isOutputValid(res, val) {
			return errors.New(fmt.Sprintf("Wrong range for %s, range: %d-%d; should be with range: %d-%d",
				val.GetName(), slices.Min(res), slices.Max(res), val.GetMinValue(), val.GetMaxValue()))
		}

		switch key {
		case 0:
			p.minutes = res
		case 1:
			p.hours = res
		case 2:
			p.daysOfMonth = res
		case 3:
			p.months = res
		case 4:
			p.daysOfWeek = res
		}
	}

	command := slicedInput[5:]
	p.command = strings.Join(command, " ")
	return nil
}

func isOutputValid(output []int, partType consts.Value) bool {
	if len(output) == 0 {
		return false
	}

	if !helpers.AreAllInRange(output, partType.GetMinValue(), partType.GetMaxValue()) {
		return false
	}
	return true
}

func parsePart(inputPart string, partType consts.Value) ([]int, error) {
	if strings.Contains(inputPart, consts.LISTING_OPRATOR) {
		return parsePartWithListing(inputPart, partType)
	}

	if strings.Contains(inputPart, consts.RANGE_OPERATOR) {
		return parsePartWithRange(inputPart, partType)
	}

	if strings.Contains(inputPart, consts.STEP_OPERATOR) {
		return parsePartWithStep(inputPart, partType)
	}

	if inputPart == consts.ASTERIKS {
		return parsePartWithAsteriks(inputPart, partType)
	}

	res, err := strconv.Atoi(inputPart)
	if err != nil {
		return []int{}, err
	}
	return []int{res}, nil
}

func parsePartWithStep(partWithStep string, partType consts.Value) ([]int, error) {
	split := strings.Split(partWithStep, consts.STEP_OPERATOR)
	multipleValues, err := helpers.GenerateValuesForRangeWithStep(split[0], split[1], partType.GetMaxValue(), partType.GetMinValue())
	if err != nil {
		return []int{}, err
	}
	return multipleValues, nil
}

func parsePartWithRange(partWithRange string, partType consts.Value) ([]int, error) {
	split := strings.Split(partWithRange, consts.RANGE_OPERATOR)
	multipleValues, err := helpers.GenerateValuesForRange(split[0], split[1])
	if err != nil {
		return []int{}, err
	}
	return multipleValues, nil
}

func parsePartWithListing(partWithListing string, partType consts.Value) ([]int, error) {
	split := strings.Split(partWithListing, consts.LISTING_OPRATOR)
	res := []int{}
	for _, elem := range split {
		if strings.Contains(elem, consts.RANGE_OPERATOR) {
			multipleValues, err := parsePartWithRange(elem, partType)
			if err != nil {
				return []int{}, err
			}
			res = append(res, multipleValues...)
		} else {
			parsed, err := strconv.Atoi(elem)
			if err != nil {
				return []int{}, err
			}
			res = append(res, parsed)
		}
	}
	return res, nil
}

func parsePartWithAsteriks(partWithAsteriks string, partType consts.Value) ([]int, error) {
	multipleValues, err := helpers.GenerateValuesForRange(fmt.Sprint(partType.GetMinValue()), fmt.Sprint(partType.GetMaxValue()))
	if err != nil {
		return []int{}, err
	}
	return multipleValues, nil
}
