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

func TestShouldParseDaysOfMonthPartProperly(t *testing.T) {
	allDays := []int{}
	for i := 1; i <= 31; i++ {
		allDays = append(allDays, i)
	}

	testScenarios := []struct {
		input    string
		expected []int
		comment  string
	}{
		{"1 1 1 1 1 cmd", []int{1}, "Should parse single value properly"},
		{"1 1 1-4 1 1 cmd", []int{1, 2, 3, 4}, "Should parse range properly"},
		{"1 1 2/10 1 1 cmd", []int{2, 12, 22}, "Should parse range with step and starting point properly"},
		{"1 1 */10 1 1 cmd", []int{1, 11, 21, 31}, "Should parse range with step and asteriks properly"},
		{"1 1 * 1 1 cmd", allDays, "Should parse asteriks properly"},
	}
	for _, scenario := range testScenarios {
		parser := NewParser()
		err := parser.Parse(scenario.input)
		if err != nil {
			t.Fatalf("Should not return error with proper input; %s", err)
		}

		if !reflect.DeepEqual(parser.daysOfMonth, scenario.expected) {
			t.Fatalf("%s, expected: %v, actual: %v", scenario.comment, scenario.expected, parser.daysOfMonth)
		}
	}
}

func TestShouldParseMonthsPartProperly(t *testing.T) {
	allMonths := []int{}
	for i := 1; i <= 12; i++ {
		allMonths = append(allMonths, i)
	}

	testScenarios := []struct {
		input    string
		expected []int
		comment  string
	}{
		{"1 1 1 1 1 cmd", []int{1}, "Should parse single value properly"},
		{"1 1 1 1-4 1 cmd", []int{1, 2, 3, 4}, "Should parse range properly"},
		{"1 1 1 2/10 1 cmd", []int{2, 12}, "Should parse range with step and starting point properly"},
		{"1 1 1 */10 1 cmd", []int{1, 11}, "Should parse range with step and asteriks properly"},
		{"1 1 1 * 1 cmd", allMonths, "Should parse asteriks properly"},
	}
	for _, scenario := range testScenarios {
		parser := NewParser()
		err := parser.Parse(scenario.input)
		if err != nil {
			t.Fatalf("Should not return error with proper input; %s", err)
		}

		if !reflect.DeepEqual(parser.months, scenario.expected) {
			t.Fatalf("%s, expected: %v, actual: %v", scenario.comment, scenario.expected, parser.months)
		}
	}
}

func TestShouldParseDayOfWeekPartProperly(t *testing.T) {
	allDays := []int{}
	for i := 0; i <= 6; i++ {
		allDays = append(allDays, i)
	}

	testScenarios := []struct {
		input    string
		expected []int
		comment  string
	}{
		{"1 1 1 1 1 cmd", []int{1}, "Should parse single value properly"},
		{"1 1 1 1 1-4 cmd", []int{1, 2, 3, 4}, "Should parse range properly"},
		{"1 1 1 1 2/10 cmd", []int{2}, "Should parse range with step and starting point properly"},
		{"1 1 1 1 */10 cmd", []int{0}, "Should parse range with step and asteriks properly"},
		{"1 1 1 1 * cmd", allDays, "Should parse asteriks properly"},
	}
	for _, scenario := range testScenarios {
		parser := NewParser()
		err := parser.Parse(scenario.input)
		if err != nil {
			t.Fatalf("Should not return error with proper input; %s", err)
		}

		if !reflect.DeepEqual(parser.daysOfWeek, scenario.expected) {
			t.Fatalf("%s, expected: %v, actual: %v", scenario.comment, scenario.expected, parser.daysOfWeek)
		}
	}
}

func TestShouldParseCommandPartProperly(t *testing.T) {
	testScenarios := []struct {
		input    string
		expected string
		comment  string
	}{
		{"1 1 1 1 1 cmd", "cmd", "Should parse command properly"},
		{"1 1 1 1 1 cmd with args", "cmd with args", "Should parse command with args properly"},
		{"1 1 1 1 1 $HOME/bin/daily.job >> $HOME/tmp/out 2>&1", "$HOME/bin/daily.job >> $HOME/tmp/out 2>&1", "Should parse command with args properly"},
	}
	for _, scenario := range testScenarios {
		parser := NewParser()
		err := parser.Parse(scenario.input)
		if err != nil {
			t.Fatalf("Should not return error with proper input; %s", err)
		}

		if parser.command != scenario.expected {
			t.Fatalf("%s, expected: %v, actual: %v", scenario.comment, scenario.expected, parser.command)
		}
	}
}
