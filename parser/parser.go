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
	fmt.Println()
}

func printField(fieldName string, field []int) {
	fmt.Printf("%-14s", fieldName)
	for i := 0; i < len(field); i++ {
		fmt.Printf("%d ", field[i])
	}
	fmt.Printf("\n")
}

func printStringField(fieldName string, field string) {
	fmt.Printf("%-14s%s\n", fieldName, field)
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

func parsePart(minutes string, partType consts.Value) ([]int, error) {
	if strings.Contains(minutes, consts.STEP_OPERATOR) {
		split := strings.Split(minutes, consts.STEP_OPERATOR)
		multipleValues, err := helpers.GenerateValuesForRangeWithStep(split[0], split[1], partType.GetMaxValue(), partType.GetMinValue())
		if err != nil {
			return []int{}, err
		}
		return multipleValues, nil
	}

	if strings.Contains(minutes, consts.RANGE_OPERATOR) {
		split := strings.Split(minutes, consts.RANGE_OPERATOR)
		multipleValues, err := helpers.GenerateValuesForRange(split[0], split[1])
		if err != nil {
			return []int{}, err
		}
		return multipleValues, nil
	}

	if strings.Contains(minutes, consts.LISTING_OPRATOR) {
		split := strings.Split(minutes, consts.LISTING_OPRATOR)
		multipleValues, err := helpers.GetMultipleIntsFromStringsSlice(split)
		if err != nil {
			return []int{}, err
		}
		return multipleValues, nil
	}

	if minutes == consts.ASTERIKS {
		multipleValues, err := helpers.GenerateValuesForRange(fmt.Sprint(partType.GetMinValue()), fmt.Sprint(partType.GetMaxValue()))
		if err != nil {
			return []int{}, err
		}
		return multipleValues, nil
	}

	res, err := strconv.Atoi(minutes)
	if err != nil {
		return []int{}, err
	}
	return []int{res}, nil
}
