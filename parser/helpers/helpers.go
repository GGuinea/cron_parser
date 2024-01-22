package helpers

import "strconv"

func AreAllInRange(values []int, min int, max int) bool {
	for _, value := range values {
		if value < min || value > max {
			return false
		}
	}
	return true
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
