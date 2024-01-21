package parser

import (
	"cron_expression_parser/parser/consts"
	"cron_expression_parser/parser/helpers"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	minutes []int
	hours   []int
	days    []int
	months  []int
	command string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) error {
	inputTab, err := getSplitInput(input)
	if err != nil {
		return err
	}
	err = p.performParse(inputTab)
	fmt.Println(consts.GetMinMinutesValue())
	if err != nil {
		return err
	}

	return nil
}

func getSplitInput(input string) ([]string, error) {
	if len(input) == 0 {
		return []string{}, errors.New("Input length should not be 0")
	}

	split := strings.Split(input, " ")

	if len(split) != 6 {
		return []string{}, errors.New("Input is in wrong format, should be: '* * * * * Command_to_execute'")
	}
	return split, nil
}

func (p *Parser) performParse(slicedInput []string) error {
	res, err := parsePart(slicedInput[0])
	p.minutes = res
	if err != nil {
		return err
	}

	return nil
}

func parsePart(minutes string) ([]int, error) {
	if strings.Contains(minutes, consts.STEP_OPERATOR) {
		split := strings.Split(minutes, consts.STEP_OPERATOR)
		multipleValues, err := helpers.GenerateValuesForRangeWithStep(split[0], split[1], 59)
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
		multipleValues, err := helpers.GenerateValuesForRange("0", "59")
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
