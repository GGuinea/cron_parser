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

func TestShouldParseMinutesPartProperly(t *testing.T) {
	allMinutes := []int{}
	for i := 0; i < 60; i++ {
		allMinutes = append(allMinutes, i)
	}

	testScenarios := []struct {
		input    string
		expected []int
		comment  string
	}{
		{"1 1 1 1 1 cmd", []int{1}, "Should parse single value properly"},
		{"1-4 1 1 1 1 cmd", []int{1, 2, 3, 4}, "Should parse range properly"},
		{"1/10 1 1 1 1 cmd", []int{1, 11, 21, 31, 41, 51}, "Should parse range with step and starting point properly"},
		{"*/10 1 1 1 1 cmd", []int{0, 10, 20, 30, 40, 50}, "Should parse range with step and asteriks properly"},
		{"* 1 1 1 1 cmd", allMinutes, "Should parse asteriks properly"},
	}

	for _, scenario := range testScenarios {
		parser := NewParser()
		err := parser.Parse(scenario.input)
		if err != nil {
			t.Fatalf("Should not return error with proper input; %s", err)
		}

		if !reflect.DeepEqual(parser.minutes, scenario.expected) {
			t.Fatalf("%s, expected: %v, actual: %v", scenario.comment, scenario.expected, parser.minutes)
		}
	}
}

func TestShouldParseHoursPartProperly(t *testing.T) {
	allHours := []int{}
	for i := 0; i < 24; i++ {
		allHours = append(allHours, i)
	}

	testScenarios := []struct {
		input    string
		expected []int
		comment  string
	}{
		{"1 3 1 1 1 cmd", []int{3}, "Should parse single value properly"},
		{"1 1-4 1 1 1 cmd", []int{1, 2, 3, 4}, "Should parse range properly"},
		{"1 1/10 1 1 1 cmd", []int{1, 11, 21}, "Should parse range with step and starting point properly"},
		{"1 */10 1 1 1 cmd", []int{0, 10, 20}, "Should parse range with step and asteriks properly"},
		{"1 * 1 1 1 cmd", allHours, "Should parse asteriks properly"},
	}
	for _, scenario := range testScenarios {
		parser := NewParser()
		err := parser.Parse(scenario.input)
		if err != nil {
			t.Fatalf("Should not return error with proper input; %s", err)
		}

		if !reflect.DeepEqual(parser.hours, scenario.expected) {
			t.Fatalf("%s, expected: %v, actual: %v", scenario.comment, scenario.expected, parser.hours)
		}
	}
}
