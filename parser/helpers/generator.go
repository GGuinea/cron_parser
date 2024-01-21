package helpers

import (
	"cron_expression_parser/parser/consts"
	"errors"
	"strconv"
)

func GenerateValuesForRange(start string, stop string) ([]int, error) {
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

func GenerateValuesForRangeWithStep(start string, step string, maxValue int) ([]int, error) {
	if start == consts.ASTERIKS {
		return GenerateValuesForRangeWithStep("0", step, maxValue)
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

func GetMultipleIntsFromStringsSlice(inputToParse []string) ([]int, error) {
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
