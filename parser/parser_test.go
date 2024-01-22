package parser

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
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

func TestShouldReturnErrorWhenInputHasWrongRange(t *testing.T) {
	testScenarios := []struct {
		input   string
		comment string
	}{
		{"61 1 1 1 1 cmd", "Should return error when minutes is not in range"},
		{"1 61 1 1 1 cmd", "Should return error when hours is not in range"},
		{"1 1 32 1 1 cmd", "Should return error when day of month is not in range"},
		{"1 1 1 13 1 cmd", "Should return error when month is not in range"},
		{"1 1 1 1 7 cmd", "Should return error when day of week is not in range range"},
	}

	for _, scenario := range testScenarios {
		parser := NewParser()
		err := parser.Parse(scenario.input)
		if err == nil {
			t.Fatalf("%s", scenario.comment)
		}
	}
}

func TestShouldParseSimpleOneDigitValuesProperly(t *testing.T) {
	parser := NewParser()
	err := parser.Parse("1 1 1 1 1 cmd")

	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, []int{1}) {
		t.Fatalf("Should parse minutes properly, expected: %v, actual: %v", []int{1}, parser.minutes)
	}

	if !reflect.DeepEqual(parser.hours, []int{1}) {
		t.Fatalf("Should parse hours properly, expected: %v, actual: %v", []int{1}, parser.hours)
	}

	if !reflect.DeepEqual(parser.daysOfMonth, []int{1}) {
		t.Fatalf("Should parse days of month properly, expected: %v, actual: %v", []int{1}, parser.daysOfMonth)
	}

	if !reflect.DeepEqual(parser.months, []int{1}) {
		t.Fatalf("Should parse months properly, expected: %v, actual: %v", []int{1}, parser.months)
	}

	if !reflect.DeepEqual(parser.daysOfWeek, []int{1}) {
		t.Fatalf("Should parse days of week properly, expected: %v, actual: %v", []int{1}, parser.daysOfWeek)
	}
}

func TestShouldParseAllValuesAsAsteriksAndReturnAllPossibleValues(t *testing.T) {
	allMinutes := []int{}
	for i := 0; i < 60; i++ {
		allMinutes = append(allMinutes, i)
	}

	allHours := []int{}
	for i := 0; i < 24; i++ {
		allHours = append(allHours, i)
	}

	allDays := []int{}
	for i := 1; i <= 31; i++ {
		allDays = append(allDays, i)
	}

	allMonths := []int{}
	for i := 1; i <= 12; i++ {
		allMonths = append(allMonths, i)
	}

	allDaysOfWeek := []int{}
	for i := 0; i <= 6; i++ {
		allDaysOfWeek = append(allDaysOfWeek, i)
	}

	parser := NewParser()
	err := parser.Parse("* * * * * cmd")
	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, allMinutes) {
		t.Fatalf("Should parse asteriks properly for minutes, expected: %v, actual: %v", allMinutes, parser.minutes)
	}

	if !reflect.DeepEqual(parser.hours, allHours) {
		t.Fatalf("Should parse asteriks properly for hours, expected: %v, actual: %v", allHours, parser.hours)
	}

	if !reflect.DeepEqual(parser.daysOfMonth, allDays) {
		t.Fatalf("Should parse asteriks properly for days of month, expected: %v, actual: %v", allDays, parser.daysOfMonth)
	}

	if !reflect.DeepEqual(parser.months, allMonths) {
		t.Fatalf("Should parse asteriks properly for months, expected: %v, actual: %v", allMonths, parser.months)
	}

	if !reflect.DeepEqual(parser.daysOfWeek, allDaysOfWeek) {
		t.Fatalf("Should parse asteriks properly for days of week, expected: %v, actual: %v", allDaysOfWeek, parser.daysOfWeek)
	}
}

func TestShouldParseListingValuesProperly(t *testing.T) {
	parser := NewParser()
	expected := []int{1, 2, 3}
	err := parser.Parse("1,2,3 1,2,3 1,2,3 1,2,3 1,2,3 cmd")
	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, expected) {
		t.Fatalf("Should parse listing values properly for minutes, expected: %v, actual: %v", expected, parser.minutes)
	}

	if !reflect.DeepEqual(parser.hours, expected) {
		t.Fatalf("Should parse listing values properly for hours, expected: %v, actual: %v", expected, parser.hours)
	}

	if !reflect.DeepEqual(parser.daysOfMonth, expected) {
		t.Fatalf("Should parse listing values properly for days of month, expected: %v, actual: %v", expected, parser.daysOfMonth)
	}

	if !reflect.DeepEqual(parser.months, expected) {
		t.Fatalf("Should parse listing values properly for months, expected: %v, actual: %v", expected, parser.months)
	}

	if !reflect.DeepEqual(parser.daysOfWeek, expected) {
		t.Fatalf("Should parse listing values properly for days of week, expected: %v, actual: %v", expected, parser.daysOfWeek)
	}
}

func TestShouldParseRangeValuesProperly(t *testing.T) {
	parser := NewParser()
	expected := []int{1, 2, 3}
	err := parser.Parse("1-3 1-3 1-3 1-3 1-3 cmd")
	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, expected) {
		t.Fatalf("Should parse listing values properly for minutes, expected: %v, actual: %v", expected, parser.minutes)
	}

	if !reflect.DeepEqual(parser.hours, expected) {
		t.Fatalf("Should parse listing values properly for hours, expected: %v, actual: %v", expected, parser.hours)
	}

	if !reflect.DeepEqual(parser.daysOfMonth, expected) {
		t.Fatalf("Should parse listing values properly for days of month, expected: %v, actual: %v", expected, parser.daysOfMonth)
	}

	if !reflect.DeepEqual(parser.months, expected) {
		t.Fatalf("Should parse listing values properly for months, expected: %v, actual: %v", expected, parser.months)
	}

	if !reflect.DeepEqual(parser.daysOfWeek, expected) {
		t.Fatalf("Should parse listing values properly for days of week, expected: %v, actual: %v", expected, parser.daysOfWeek)
	}
}

func TestShouldParseListingAndRangeValuesProperly(t *testing.T) {
	parser := NewParser()
	expected := []int{1, 2, 3, 5, 6, 9, 10, 11}
	expectedDayOfWeek := []int{1, 2, 3, 5, 6}
	err := parser.Parse("1-3,5,6,9-11 1-3,5,6,9-11 1-3,5,6,9-11 1-3,5,6,9-11 1-3,5,6 cmd")
	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, expected) {
		t.Fatalf("Should parse listing values properly for minutes, expected: %v, actual: %v", expected, parser.minutes)
	}

	if !reflect.DeepEqual(parser.hours, expected) {
		t.Fatalf("Should parse listing values properly for hours, expected: %v, actual: %v", expected, parser.hours)
	}

	if !reflect.DeepEqual(parser.daysOfMonth, expected) {
		t.Fatalf("Should parse listing values properly for days of month, expected: %v, actual: %v", expected, parser.daysOfMonth)
	}

	if !reflect.DeepEqual(parser.months, expected) {
		t.Fatalf("Should parse listing values properly for months, expected: %v, actual: %v", expected, parser.months)
	}

	if !reflect.DeepEqual(parser.daysOfWeek, expectedDayOfWeek) {
		t.Fatalf("Should parse listing values properly for days of week, expected: %v, actual: %v", expectedDayOfWeek, parser.daysOfWeek)
	}
}

func TestShouldParseStepValuesProperly(t *testing.T) {
	parser := NewParser()

	expectedMinutes := []int{1, 21, 41}
	expectedHours := []int{1, 11, 21}
	expectedDaysOfMonth := []int{1, 11, 21, 31}
	expectedMonth := []int{1, 11}
	expectedDayOfWeek := []int{1, 3, 5}

	err := parser.Parse("1/20 1/10 1/10 1/10 1/2 cmd")
	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, expectedMinutes) {
		t.Fatalf("Should parse listing values properly for minutes, expected: %v, actual: %v", expectedMinutes, parser.minutes)
	}

	if !reflect.DeepEqual(parser.hours, expectedHours) {
		t.Fatalf("Should parse listing values properly for hours, expected: %v, actual: %v", expectedHours, parser.hours)
	}

	if !reflect.DeepEqual(parser.daysOfMonth, expectedDaysOfMonth) {
		t.Fatalf("Should parse listing values properly for days of month, expected: %v, actual: %v", expectedDaysOfMonth, parser.daysOfMonth)
	}

	if !reflect.DeepEqual(parser.months, expectedMonth) {
		t.Fatalf("Should parse listing values properly for months, expected: %v, actual: %v", expectedMonth, parser.months)
	}

	if !reflect.DeepEqual(parser.daysOfWeek, expectedDayOfWeek) {
		t.Fatalf("Should parse listing values properly for days of week, expected: %v, actual: %v", expectedDayOfWeek, parser.daysOfWeek)
	}
}

func TestShouldParseRangeWithStepValuesProperly(t *testing.T) {
	parser := NewParser()
	expected := []int{1, 3, 5, 7, 9}
	expectedDayOfWeek := []int{1, 3, 5}
	err := parser.Parse("1-10/2 1-10/2 1-10/2 1-10/2 1-6/2 cmd")
	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	if !reflect.DeepEqual(parser.minutes, expected) {
		t.Fatalf("Should parse listing values properly for minutes, expected: %v, actual: %v", expected, parser.minutes)
	}

	if !reflect.DeepEqual(parser.hours, expected) {
		t.Fatalf("Should parse listing values properly for hours, expected: %v, actual: %v", expected, parser.hours)
	}

	if !reflect.DeepEqual(parser.daysOfMonth, expected) {
		t.Fatalf("Should parse listing values properly for days of month, expected: %v, actual: %v", expected, parser.daysOfMonth)
	}

	if !reflect.DeepEqual(parser.months, expected) {
		t.Fatalf("Should parse listing values properly for months, expected: %v, actual: %v", expected, parser.months)
	}

	if !reflect.DeepEqual(parser.daysOfWeek, expectedDayOfWeek) {
		t.Fatalf("Should parse listing values properly for days of week, expected: %v, actual: %v", expectedDayOfWeek, parser.daysOfWeek)
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

func TestShouldPrintAllValuesWithProperFormat(t *testing.T) {
	expectedOutput :=
		`minute        1 
hour          1 
day of month  1 
month         1 
day of week   1 
command       cmd`

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	parser := NewParser()
	err := parser.Parse("1 1 1 1 1 cmd")
	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}

	parser.PrintCurrentCronExpression()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout
	fmt.Println(strings.Contains(string(out), expectedOutput))

	evaluated := strings.Trim(string(out), "")
	evaluated = strings.Trim(string(out), "\n")

	if reflect.DeepEqual(evaluated, expectedOutput) != true {
		t.Fatalf("Should print all values with proper format")
	}

}

func TestShouldParseProperlyExample(t *testing.T) {
	input := "*/15 0 1,15 * 1-5 /usr/bin/find"
	parser := NewParser()
	err := parser.Parse(input)
	if err != nil {
		t.Fatalf("Should not return error with proper input; %s", err)
	}
	if reflect.DeepEqual(parser.minutes, []int{0, 15, 30, 45}) != true {
		t.Fatalf("Should parse minutes properly")
	}
	if reflect.DeepEqual(parser.hours, []int{0}) != true {
		t.Fatalf("Should parse hours properly")
	}

	if reflect.DeepEqual(parser.daysOfMonth, []int{1, 15}) != true {
		t.Fatalf("Should parse days of month properly")
	}

	if reflect.DeepEqual(parser.months, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}) != true {
		t.Fatalf("Should parse months properly")
	}

	if reflect.DeepEqual(parser.daysOfWeek, []int{1, 2, 3, 4, 5}) != true {
		t.Fatalf("Should parse days of week properly")
	}

	if parser.command != "/usr/bin/find" {
		t.Fatalf("Should parse command properly")
	}
}
