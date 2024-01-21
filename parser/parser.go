package parser

import (
	"errors"
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

// * * * * * Command_to_execute

func NewParser() *Parser {
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
	if strings.Contains(minutes, "/") {
		split := strings.Split(minutes, "/")
		multipleValues, err := generateValuesForRangeWithStep(split[0], split[1], 59)
		if err != nil {
			return []int{}, err
		}
		return multipleValues, nil
	}

	if strings.Contains(minutes, "-") {
		split := strings.Split(minutes, "-")
		multipleValues, err := generateValuesForRange(split[0], split[1])
		if err != nil {
			return []int{}, err
		}
		return multipleValues, nil
	}

	if strings.Contains(minutes, ",") {
		split := strings.Split(minutes, ",")
		multipleValues, err := parseMultipleIntsFromStringsSlice(split)
		if err != nil {
			return []int{}, err
		}
		return multipleValues, nil
	}

	if minutes == "*" {
		multipleValues, err := generateValuesForRange("0", "59")
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

func parseMultipleIntsFromStringsSlice(inputToParse []string) ([]int, error) {
	res := []int{}

	for _, elem := range inputToParse {
		parsed, err := strconv.Atoi(elem)
		if err != nil {
			return []int{}, err
		}
		res = append(res, parsed)
	}

	return res, nil
}

func generateValuesForRange(start string, stop string) ([]int, error) {
	startParsed, err := strconv.Atoi(start)
	if err != nil {
		return []int{}, err
	}

	stopParsed, err := strconv.Atoi(stop)
	if err != nil {
		return []int{}, err
	}

	if startParsed > stopParsed {
		return []int{}, errors.New("Wrong range values")
	}

	res := []int{}
	for i := startParsed; i <= stopParsed; i++ {
		res = append(res, i)
	}

	return res, nil
}

func generateValuesForRangeWithStep(start string, step string, maxValue int) ([]int, error) {
	if start == "*" {
		return generateValuesForRangeWithStep("0", step, maxValue)
	}

	startParsed, err := strconv.Atoi(start)
	if err != nil {
		return []int{}, err
	}

	stepParsed, err := strconv.Atoi(step)
	if err != nil {
		return []int{}, err
	}

	res := []int{}
	for i := startParsed; i <= maxValue; i += stepParsed {
		res = append(res, i)
	}

	return res, nil
}
