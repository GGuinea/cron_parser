package helpers

import (
	"cron_expression_parser/parser/consts"
	"errors"
	"fmt"
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

func GenerateValuesForRangeWithStep(start string, step string, maxValue, minValue int) ([]int, error) {
	if start == consts.ASTERIKS {
		return GenerateValuesForRangeWithStep(fmt.Sprint(minValue), step, maxValue, minValue)
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
