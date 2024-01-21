package parser

import (
	"reflect"
	"testing"
)

func TestShouldReturnNewEmptyParserClient(t *testing.T) {
	parser := NewParser()
	if parser == nil {
		t.Fatalf("Should return new parser")
	}
}

func TestShouldReturnErrorWhenInputStringIsEmpty(t *testing.T) {
	parser := NewParser()
	err := parser.Parse("")
	if err == nil {
		t.Fatalf("Should return error when input is an empty string")
	}
}

func TestShouldReturnErrorWhenInputDoesNotHaveSixParts(t *testing.T) {
	parser := NewParser()
	err := parser.Parse("1 1 1 111")
	if err == nil {
		t.Fatalf("Should return error when input has wrong format")
	}
}

func TestShouldParseProperlySimpleDayInput(t *testing.T) {
	parser := NewParser()
	err := parser.Parse("1 1 1 1 1 1")

	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, []int{1}) {
		t.Fatalf("Should parse minutes properly")
	}
}

func TestShouldParseProperlySimpleDayInputAsAsteriks(t *testing.T) {
	parser := NewParser()
	err := parser.Parse("* 1 1 1 1 1")

	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	expectedRes := []int{}

	for i := 0; i < 60; i++ {
		expectedRes = append(expectedRes, i)
	}

	if !reflect.DeepEqual(parser.minutes, expectedRes) {
		t.Fatalf("Should parse minutes properly")
	}
}

func TestShouldParseProperlyMultipleMinuteValuesSeparatedByCommas(t *testing.T) {
	parser := NewParser()
	err := parser.Parse("1,4,59 1 1 1 1 1")

	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, []int{1, 4, 59}) {
		t.Fatalf("Should parse minutes properly")
	}
}

func TestShouldParseProperlyMultipleMinuteValuesWithRangeOperatorInclusively(t *testing.T) {
	parser := NewParser()
	err := parser.Parse("1-4 1 1 1 1 1")

	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, []int{1, 2, 3, 4}) {
		t.Fatalf("Should parse minutes properly")
	}
}

func TestShouldParseProperlyMultipleMinuteValuesWithStepOperator(t *testing.T) {
	parser := NewParser()
	err := parser.Parse("1/10 1 1 1 1 1")

	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, []int{1, 11, 21, 31, 41, 51}) {
		t.Fatalf("Should parse minutes properly, actual: %v", parser.minutes)
	}
}

func TestShouldParseProperlyMultipleMinuteValuesWithStepAndAsteriskSymbol(t *testing.T) {
	parser := NewParser()
	err := parser.Parse("*/10 1 1 1 1 1")

	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, []int{0, 10, 20, 30, 40, 50}) {
		t.Fatalf("Should parse minutes properly, actual: %v", parser.minutes)
	}
}
